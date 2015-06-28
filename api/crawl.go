package api

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
)

type crawler struct {
	fetching chan fetchJob
}

type fetchJob struct {
	url    string
	result chan *Cast
}

func newCrawler() *crawler {
	return &crawler{
		fetching: make(chan fetchJob, 4096),
	}
}

func (c *crawler) start(maxConn int) {
	for i := 0; i < maxConn; i++ {
		go c.fetcher()
	}
}

func (c *crawler) stop() {
	close(c.fetching)
}

func (c *crawler) fetch(url string) chan *Cast {
	resultCh := make(chan *Cast, 1)
	c.fetching <- fetchJob{url: url, result: resultCh}
	return resultCh
}

func (c *crawler) fetcher() {
	for {
		job, ok := <-c.fetching
		if !ok {
			job.result <- nil
			return
		}

		job.result <- c.download(job.url)
	}
}

func (c *crawler) download(url string) *Cast {
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil
	}

	return c.decode(resp.Body, url)
}

func (c *crawler) decode(body io.Reader, url string) *Cast {
	decoder := xml.NewDecoder(body)
	for {
		token, _ := decoder.Token()
		if token == nil {
			return nil
		}

		switch t := token.(type) {
		case xml.StartElement:
			if t.Name.Local == "channel" {
				feed := feedRSS{}
				decoder.DecodeElement(&feed, &t)

				cast := &Cast{URL: url, Name: feed.Title}
				v, _ := json.Marshal(feed)
				cast.Feed = (*json.RawMessage)(&v)

				store.SaveCast(cast)
				return cast
			} else if t.Name.Local == "feed" {
				feed := feedAtom{}
				decoder.DecodeElement(&feed, &t)

				cast := &Cast{URL: url, Name: feed.Title}
				v, _ := json.Marshal(feed)
				cast.Feed = (*json.RawMessage)(&v)

				store.SaveCast(cast)
				return cast
			}
		}
	}
}
