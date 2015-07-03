package api

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

func TestGetCasts(t *testing.T) {
	r := createRouter()

	expectedJSON := testJSON([]Cast{
		Cast{
			ID:   1,
			URL:  "test.go",
			Name: "test",
		},
	})

	req := testRequest(r, "GET", "/library/casts", nil)
	req.Header.Set("Authorization", "token")
	res := req.send()
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, expectedJSON, res.Body.String())
}

func TestAddCast(t *testing.T) {
	r := createRouter()
	cast := &Cast{}

	// It should return an existing cast
	req := testRequest(r, "POST", "/library/casts", nil)
	req.Header.Set("Authorization", "token")
	req.PostForm = url.Values{}
	req.PostForm.Set("feedurl", "test.go")
	res := req.send()
	assert.Equal(t, 200, res.Code)
	json.Unmarshal(res.Body.Bytes(), cast)
	assert.Equal(t, "test.go", cast.URL)
	assert.Equal(t, "test", cast.Name)

	// There should still be only 1 subscription
	user := store.GetUser("test")
	assert.Len(t, user.Subscriptions, 1)

	// It should return 500 when the crawling fails
	req.PostForm.Set("feedurl", "dat_url")
	res = req.send()
	assert.Equal(t, 500, res.Code)

	// There should still be only 1 subscription
	user = store.GetUser("test")
	assert.Len(t, user.Subscriptions, 1)

	// It should return a new cast
	req.PostForm.Set("feedurl", testRSS)
	res = req.send()
	assert.Equal(t, 200, res.Code)
	json.Unmarshal(res.Body.Bytes(), cast)
	assert.Equal(t, testRSS, cast.URL)
	assert.Equal(t, "BSD Now HD", cast.Name)
	assert.NotNil(t, cast.Feed)

	// There should now be 2 subscriptions
	user = store.GetUser("test")
	assert.Len(t, user.Subscriptions, 2)

	// The new cast should be in the store
	cast = store.GetCastByURL(testRSS)
	assert.NotNil(t, cast)
	assert.Equal(t, "BSD Now HD", cast.Name)
}

func TestRenameCast(t *testing.T) {
	r := createRouter()

	req := testRequest(r, "PUT", "/library/casts/1", nil)
	req.Header.Set("Authorization", "token")
	req.PostForm = url.Values{}
	req.PostForm.Set("name", "new")
	res := req.send()
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "new", store.GetCast(1).Name)

	// It should return 400 if the ID is invalid
	req.URL.Path = "/library/casts/nope"
	assert.Equal(t, 400, req.send().Code)

	// It returns 200 if the cast is not found
	req.URL.Path = "/library/casts/1337"
	assert.Equal(t, 200, req.send().Code)
}

func TestRemoveCast(t *testing.T) {
	r := createRouter()

	req := testRequest(r, "DELETE", "/library/casts/1", nil)
	req.Header.Set("Authorization", "token")
	assert.Equal(t, 200, req.send().Code)
	user := store.GetUser("test")
	assert.NotContains(t, user.Subscriptions, uint64(1))

	// It should return 400 if the ID is invalid
	req.URL.Path = "/library/casts/nope"
	assert.Equal(t, 400, req.send().Code)
}
