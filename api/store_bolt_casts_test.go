package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoreGetCasts(t *testing.T) {
	casts := store.GetCasts()
	assert.NotNil(t, casts)
	assert.True(t, len(casts) > 0)
	assert.Equal(t, uint64(1), casts[0].ID)
	assert.Equal(t, "test.go", casts[0].URL)
}
