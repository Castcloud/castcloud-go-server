package api

import (
	"testing"

	"github.com/khlieng/castcloud-go/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

func TestCrawlFetch(t *testing.T) {
	cast := <-crawl.fetch(testServer.URL)
	assert.NotNil(t, cast)
	assert.Equal(t, testServer.URL, cast.URL)
	assert.Equal(t, "BSD Now HD", cast.Name)
}
