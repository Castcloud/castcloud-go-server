package api

import (
	"encoding/json"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	. "github.com/Castcloud/castcloud-go-server/api/schema"
)

func TestGetEvents(t *testing.T) {
	r := createRouter(true)

	// It should return 400 if since is bad
	req := testRequest(r, "GET", "/library/events?since=yesterday", nil)
	req.Header.Set("Authorization", "evtest1")
	assert.Equal(t, 400, req.send().Code)

	// It should return 200 with a timestamp and a list of events
	req.URL, _ = url.Parse("/library/events?since=5")
	res := checkEvents(t, req)
	assert.Len(t, res.Events, 1)

	// It should exclude events sent from this clients UUID when
	// exclude_self is true
	req.URL, _ = url.Parse("/library/events?since=5&exclude_self=true")
	res = checkEvents(t, req)
	assert.Len(t, res.Events, 0)
}

func checkEvents(t *testing.T, req testReq) events {
	now := time.Now().Unix()
	res := req.send()
	assert.Equal(t, 200, res.Code)
	data := events{}
	err := json.Unmarshal(res.Body.Bytes(), &data)
	assert.Nil(t, err)
	assert.True(t, data.Timestamp >= now)
	return data
}

func TestAddEvents(t *testing.T) {
	r := createRouter(true)

	// It should return 400 if json is not set
	req := testRequest(r, "POST", "/library/events", nil)
	req.Header.Set("Authorization", "evtest1")
	assert.Equal(t, 400, req.send().Code)

	// It should return 400 if the json is bad
	req.PostForm = url.Values{}
	req.PostForm.Set("json", "real_bad_json")
	assert.Equal(t, 400, req.send().Code)

	// It should return 200 when proper json is sent
	req.PostForm.Set("json", testJSON([]Event{
		Event{
			Type:              30,
			EpisodeID:         10,
			PositionTS:        481,
			ClientTS:          11,
			ConcurrentOrder:   0,
			ClientName:        "Castcloud",
			ClientDescription: "oink",
		},
	}))
	assert.Equal(t, 200, req.send().Code)

	// There should now be 2 events
	req.Method = "GET"
	res := checkEvents(t, req)
	assert.Len(t, res.Events, 2)
}
