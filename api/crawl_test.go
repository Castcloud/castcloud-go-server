package api

import (
	"testing"

	"github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

func TestCrawlFetch(t *testing.T) {
	// It should return the cast
	cast := <-crawl.fetch(testRSS)
	assert.NotNil(t, cast)
	assert.Equal(t, testRSS, cast.URL)
	assert.Equal(t, "BSD Now HD", cast.Name)

	// It should return nil if the url is bad
	assert.Nil(t, <-crawl.fetch("so_bad"))

	// It should return nil if the the status is not 200
	assert.Nil(t, <-crawl.fetch(testServer.URL))

	// It should return nil if it cant parse things
	assert.Nil(t, <-crawl.fetch(testServer.URL+"/notxml"))

	// It should return nil if the xml is not a proper feed
	assert.Nil(t, <-crawl.fetch(testServer.URL+"/notfeed"))
	assert.Nil(t, <-crawl.fetch(testServer.URL+"/badrss"))
	assert.Nil(t, <-crawl.fetch(testServer.URL+"/badatom"))
}

func TestCrawlFetchAtom(t *testing.T) {
	cast := <-crawl.fetch(testAtom)
	assert.NotNil(t, cast)
	assert.Equal(t, testAtom, cast.URL)
	assert.Equal(t, "Example Feed", cast.Name)
}
