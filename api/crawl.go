package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/github.com/khlieng/mxj"

	. "github.com/Castcloud/castcloud-go-server/api/schema"
)

const (
	feedFormatRSS = iota
	feedFormatAtom
)

type crawler struct {
	interval time.Duration
	fetching chan fetchJob
	saving   chan saveJob
	quit     chan struct{}
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

func newCrawler(interval time.Duration) *crawler {
	return &crawler{
		interval: interval,
		fetching: make(chan fetchJob, 4096),
		saving:   make(chan saveJob, 256),
		quit:     make(chan struct{}),
	}
}

func (c *crawler) start(maxConn int) {
	for i := 0; i < maxConn; i++ {
		go c.fetcher()
	}

	for i := 0; i < 64; i++ {
		go c.saver()
	}

	go c.run()
}

func (c *crawler) stop() {
	close(c.quit)
	close(c.fetching)
	close(c.saving)
}

func (c *crawler) fetch(url string) chan *Cast {
	resultCh := make(chan *Cast, 1)
	c.fetching <- fetchJob{url: url, result: resultCh}
	return resultCh
}

func (c *crawler) run() {
	tick := time.Tick(c.interval)

	c.crawlCasts()
	for {
		select {
		case <-c.quit:
			return

		case <-tick:
			c.crawlCasts()
		}
	}
}

func (c *crawler) crawlCasts() {
	for _, cast := range store.GetCasts() {
		c.fetch(cast.URL)
	}
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
	var castFeed interface{}
	var episodes []interface{}
	var format int
	ts := time.Now().Unix()

	if job.xml.Exists("rss.channel") {
		format = feedFormatRSS
		name, err = job.xml.ValueForPathString("rss.channel.title")
		if err != nil {
			return false
		}

		castFeed, _ = job.xml.ValueForPath("rss.channel")
		episodes, _ = job.xml.ValuesForPath("rss.channel.item")
		job.xml.Remove("rss.channel.item")
	} else if job.xml.Exists("feed") {
		format = feedFormatAtom
		name, err = job.xml.ValueForPathString("feed.title")
		if err != nil {
			return false
		}

		castFeed, _ = job.xml.ValueForPath("feed")
		episodes, _ = job.xml.ValuesForPath("feed.entry")
		job.xml.Remove("feed.entry")
	} else {
		return false
	}

	feed, err := json.Marshal(castFeed)
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

		guid := extractGUID(ep, format)
		if guid == "" {
			title, ok := ep["title"].(string)
			if ok {
				guid = md5Hash(cast.URL + title)
			}
		}

		eps[i].GUID = guid
		eps[i].CrawlTS = ts
		feed, _ := json.Marshal(ep)
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
