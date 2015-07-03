package api

import (
	"encoding/json"
	"net/url"
	"testing"
	"time"

	"github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

func TestGetNewEpisodes(t *testing.T) {
	r := createRouter()

	req := testRequest(r, "GET", "/library/newepisodes?since=32503679999", nil)
	req.Header.Set("Authorization", "token")
	res := checkNewEpisodes(t, req)
	assert.Len(t, res.Episodes, 2)
	assert.Equal(t, Episode{
		ID:      2,
		CastID:  69,
		GUID:    "since1",
		CrawlTS: 32503680000,
	}, res.Episodes[0])
	assert.Equal(t, Episode{
		ID:      3,
		CastID:  69,
		GUID:    "since2",
		CrawlTS: 32503680001,
	}, res.Episodes[1])

	req.URL, _ = url.Parse("/library/newepisodes?since=32503680000")
	res = checkNewEpisodes(t, req)
	assert.Len(t, res.Episodes, 1)
	assert.Equal(t, Episode{
		ID:      3,
		CastID:  69,
		GUID:    "since2",
		CrawlTS: 32503680001,
	}, res.Episodes[0])

	req.URL, _ = url.Parse("/library/newepisodes?since=32503680001")
	res = checkNewEpisodes(t, req)
	assert.Len(t, res.Episodes, 0)

	req.URL, _ = url.Parse("/library/newepisodes")
	res = checkNewEpisodes(t, req)
	assert.True(t, len(res.Episodes) > 2)

	req.URL, _ = url.Parse("/library/newepisodes?since=what")
	resp := req.send()
	assert.Equal(t, 400, resp.Code)
}

func checkNewEpisodes(t *testing.T, req testReq) newEpisodes {
	now := time.Now().Unix()
	res := req.send()
	assert.Equal(t, 200, res.Code)
	data := newEpisodes{}
	err := json.Unmarshal(res.Body.Bytes(), &data)
	assert.Nil(t, err)
	assert.True(t, data.Timestamp >= now)
	return data
}

func TestGetEpisodes(t *testing.T) {
	r := createRouter()

	expectedJSON := testJSON([]Episode{
		Episode{
			ID:     1,
			CastID: 1,
			GUID:   "guid",
		},
	})

	// It should return 200 and a list of episodes
	req := testRequest(r, "GET", "/library/episodes/1", nil)
	req.Header.Set("Authorization", "token")
	res := req.send()
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, expectedJSON, res.Body.String())

	// It should return 200 and an empty list if the cast is not found
	req.URL.Path = "/library/episodes/1337"
	res = req.send()
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, testJSON([]Episode{}), res.Body.String())

	// It should return 400 if the ID is invalid
	req.URL.Path = "/library/episodes/datcast"
	assert.Equal(t, 400, req.send().Code)
}

func TestGetEpisode(t *testing.T) {
	r := createRouter()

	expectedJSON := testJSON(Episode{
		ID:     1,
		CastID: 1,
		GUID:   "guid",
	})

	req := testRequest(r, "GET", "/library/episode/1", nil)
	req.Header.Set("Authorization", "token")
	res := req.send()
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, expectedJSON, res.Body.String())

	// It should return 400 if the ID is invalid
	req.URL.Path = "/library/episode/datepisode"
	assert.Equal(t, 400, req.send().Code)
}
