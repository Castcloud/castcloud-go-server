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
	cast = <-crawl.fetch("so_bad")
	assert.Nil(t, cast)

	// It should return nil if the the status is not 200
	cast = <-crawl.fetch(testServer.URL)
	assert.Nil(t, cast)

	// It should return nil if it cant parse things
	cast = <-crawl.fetch(testServer.URL + "/notxml")
	assert.Nil(t, cast)
}

func TestCrawlFetchAtom(t *testing.T) {
	cast := <-crawl.fetch(testAtom)
	assert.NotNil(t, cast)
	assert.Equal(t, testAtom, cast.URL)
	assert.Equal(t, "Example Feed", cast.Name)
}
