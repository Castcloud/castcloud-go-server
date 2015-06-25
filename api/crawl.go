package api

import (
	"encoding/xml"
	"io"
	"net/http"
	"runtime"
)

type crawler struct {
	fetching chan fetchJob
	parsing  chan parseJob
}

type fetchJob struct {
	url    string
	result chan *Cast
}

type parseJob struct {
	data   io.ReadCloser
	url    string
	result chan *Cast
}

func newCrawler() *crawler {
	return &crawler{
		fetching: make(chan fetchJob, 4096),
		parsing:  make(chan parseJob, 64),
	}
}

func (c *crawler) start(maxConn int) {
	for i := 0; i < runtime.NumCPU(); i++ {
		go c.parser()
	}

	for i := 0; i < maxConn; i++ {
		go c.fetcher()
	}
}

func (c *crawler) stop() {
	close(c.fetching)
	close(c.parsing)
}

func (c *crawler) fetch(url string) chan *Cast {
	castCh := make(chan *Cast, 1)
	c.fetching <- fetchJob{url: url, result: castCh}
	return castCh
}

func (c *crawler) fetcher() {
	for {
		job, ok := <-c.fetching
		if !ok {
			job.result <- nil
			return
		}

		resp, err := http.Get(job.url)
		if err != nil {
			job.result <- nil
			return
		}

		if resp.StatusCode == 200 {
			c.parsing <- parseJob{data: resp.Body, url: job.url, result: job.result}
		} else {
			job.result <- nil
		}
	}
}

func (c *crawler) parser() {
	for {
		job, ok := <-c.parsing
		defer job.data.Close()
		if !ok {
			job.result <- nil
			return
		}

		f := feed{}

		decoder := xml.NewDecoder(job.data)
		for {
			t, _ := decoder.Token()
			if t == nil {
				job.result <- nil
				return
			}

			switch se := t.(type) {
			case xml.StartElement:
				if se.Name.Local == "channel" {
					decoder.DecodeElement(&f, &se)
					println("RSS: " + f.Title)
					cast := &Cast{URL: job.url, Name: f.Title}
					store.SaveCast(cast)
					job.result <- cast
					return
				} else if se.Name.Local == "feed" {
					decoder.DecodeElement(&f, &se)
					println("Atom: " + f.Title)
					job.result <- &Cast{URL: job.url, Name: f.Title}
					return
				}
			}
		}
	}
}
