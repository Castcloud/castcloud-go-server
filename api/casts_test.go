package api

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
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

	expectedJSON := testJSON(Cast{
		ID:   1,
		URL:  "test.go",
		Name: "test",
	})

	// It should return an existing cast
	req := testRequest(r, "POST", "/library/casts", nil)
	req.Header.Set("Authorization", "token")
	req.PostForm = url.Values{}
	req.PostForm.Set("feedurl", "test.go")
	res := req.send()
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, expectedJSON, res.Body.String())

	// There should still be only 1 subscription
	user := store.GetUser("test")
	assert.Len(t, user.Subscriptions, 1)

	// It should (probably) return a null when the crawling fails
	req.PostForm.Set("feedurl", "dat_url")
	res = req.send()
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, testJSON(nil), res.Body.String())

	// There should still be only 1 subscription
	user = store.GetUser("test")
	assert.Len(t, user.Subscriptions, 1)

	// It should return a new cast
	expectedJSON = testJSON(Cast{
		ID:   2,
		URL:  "http://feeds.feedburner.com/BsdNowHd",
		Name: "BSD Now HD",
	})

	req.PostForm.Set("feedurl", "http://feeds.feedburner.com/BsdNowHd")
	res = req.send()
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, expectedJSON, res.Body.String())

	// There should now be 2 subscriptions
	user = store.GetUser("test")
	assert.Len(t, user.Subscriptions, 2)
}
