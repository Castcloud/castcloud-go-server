package schema

import (
	"encoding/json"
)

//go:generate msgp

type User struct {
	ID            uint64
	Username      string
	Password      string
	Clients       []*Client
	Subscriptions []uint64
}

type Client struct {
	Token string
	UUID  string
	Name  string
}

type Cast struct {
	ID       uint64           `json:"id"`
	URL      string           `json:"url"`
	Name     string           `json:"name"`
	Feed     *json.RawMessage `json:"feed" msg:"-" `
	FeedMsgp []byte           `json:"-"`
}

type Episode struct {
	ID        uint64           `json:"id"`
	CastID    uint64           `json:"castid"`
	LastEvent *Event           `json:"lastevent"`
	Feed      *json.RawMessage `json:"feed" msg:"-" `
	FeedMsgp  []byte           `json:"-"`
	GUID      string           `json:"-"`
	CrawlTS   int64            `json:"-"`
}

type Event struct {
	Type              int    `json:"type"`
	EpisodeID         uint64 `json:"episodeid"`
	PositionTS        int    `json:"positionts"`
	ClientTS          uint64 `json:"clientts"`
	ConcurrentOrder   int    `json:"concurrentorder"`
	ClientName        string `json:"clientname"`
	ClientDescription string `json:"clientdescription"`
	ClientUUID        string `json:"-"`
}

type Label struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Content  string `json:"content"`
	Expanded bool   `json:"expanded"`
	Root     bool   `json:"root"`
}

func (u *User) UUID(token string) string {
	for _, client := range u.Clients {
		if client.Token == token {
			return client.UUID
		}
	}
	return ""
}

func (c *Cast) EncodeFeed() {
	if c.Feed != nil {
		c.FeedMsgp = *c.Feed
	}
}

func (c *Cast) DecodeFeed() {
	if c.FeedMsgp != nil {
		c.Feed = (*json.RawMessage)(&c.FeedMsgp)
	}
}

func (e *Episode) EncodeFeed() {
	if e.Feed != nil {
		e.FeedMsgp = *e.Feed
	}
}

func (e *Episode) DecodeFeed() {
	if e.FeedMsgp != nil {
		e.Feed = (*json.RawMessage)(&e.FeedMsgp)
	}
}
