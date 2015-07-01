package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/github.com/khlieng/mxj"
)

const (
	feedFormatRSS = iota
	feedFormatAtom
)

type crawler struct {
	fetching chan fetchJob
	saving   chan saveJob
}

type fetchJob struct {
	url    string
	result chan *Cast
}

type saveJob struct {
	url    string
	xml    mxj.Map
	result chan *Cast
}

func newCrawler() *crawler {
	return &crawler{
		fetching: make(chan fetchJob, 4096),
		saving:   make(chan saveJob, 256),
	}
}

func (c *crawler) start(maxConn int) {
	for i := 0; i < maxConn; i++ {
		go c.fetcher()
	}

	for i := 0; i < 64; i++ {
		go c.saver()
	}
}

func (c *crawler) stop() {
	close(c.fetching)
	close(c.saving)
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

		if !c.download(job) {
			job.result <- nil
		}
	}
}

func (c *crawler) saver() {
	for {
		job, ok := <-c.saving
		if !ok {
			job.result <- nil
			return
		}

		if !c.save(job) {
			job.result <- nil
		}
	}
}

func (c *crawler) download(job fetchJob) bool {
	resp, err := http.Get(job.url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return false
	}

	m, err := mxj.NewMapXmlReader(resp.Body)
	if err != nil {
		return false
	}

	c.saving <- saveJob{url: job.url, xml: m, result: job.result}
	return true
}

func (c *crawler) save(job saveJob) bool {
	var err error
	var name string
	var episodes []interface{}
	var format int
	ts := time.Now().Unix()

	if job.xml.Exists("rss.channel") {
		format = feedFormatRSS
		name, err = job.xml.ValueForPathString("rss.channel.title")
		episodes, err = job.xml.ValuesForPath("rss.channel.item")
		job.xml.Remove("rss.channel.item")
	} else if job.xml.Exists("feed") {
		format = feedFormatAtom
		name, err = job.xml.ValueForPathString("feed.title")
		episodes, err = job.xml.ValuesForPath("feed.entry")
		job.xml.Remove("feed.entry")
	} else {
		return false
	}
	if err != nil {
		return false
	}

	feed, err := job.xml.Json()
	if err != nil {
		return false
	}

	cast := &Cast{URL: job.url, Name: name, Feed: (*json.RawMessage)(&feed)}
	err = store.SaveCast(cast)
	if err != nil {
		return false
	}

	job.result <- cast

	eps := make([]Episode, len(episodes))
	for i := range episodes {
		ep := episodes[i].(map[string]interface{})
		eps[i].CastID = cast.ID
		eps[i].GUID = extractGUID(ep, format)
		eps[i].CrawlTS = ts
		feed, _ = json.Marshal(ep)
		eps[i].Feed = (*json.RawMessage)(&feed)
	}

	store.SaveEpisodes(eps)
	return true
}

func extractGUID(episode map[string]interface{}, format int) string {
	switch format {
	case feedFormatRSS:
		m, ok := episode["guid"].(map[string]interface{})
		if ok {
			guid, ok := m["_"].(string)
			if ok {
				return guid
			}
		} else {
			guid, ok := episode["guid"].(string)
			if ok {
				return guid
			}
		}

	case feedFormatAtom:
		guid, ok := episode["id"].(string)
		if ok {
			return guid
		}
	}
	return ""
}
