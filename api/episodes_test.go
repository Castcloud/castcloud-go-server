package api

import (
	"testing"

	"github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

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

	// It should return 500 if the ID is invalid
	req.URL.Path = "/library/episodes/datcast"
	assert.Equal(t, 500, req.send().Code)
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

	// It should return 500 if the ID is invalid
	req.URL.Path = "/library/episode/datepisode"
	assert.Equal(t, 500, req.send().Code)
}
