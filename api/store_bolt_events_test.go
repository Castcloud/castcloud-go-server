package api

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/Castcloud/castcloud-go-server/api/schema"
)

func TestEvents(t *testing.T) {
	err := store.AddEvents([]Event{
		Event{
			Type:              10,
			EpisodeID:         1,
			PositionTS:        134,
			ClientTS:          100,
			ConcurrentOrder:   0,
			ClientName:        "Castcloud",
			ClientDescription: "oink",
		},
		Event{
			Type:              10,
			EpisodeID:         1,
			PositionTS:        112,
			ClientTS:          50,
			ConcurrentOrder:   0,
			ClientName:        "Castcloud",
			ClientDescription: "oink",
		},
	}, 2, "evuuid1")
	assert.Nil(t, err)

	events := store.GetEvents(2, 100, "")
	assert.Len(t, events, 0)

	events = store.GetEvents(2, 99, "")
	assert.Len(t, events, 1)
	assert.Equal(t, int32(10), events[0].Type)
	assert.Equal(t, int32(134), events[0].PositionTS)

	events = store.GetEvents(2, 99, "evuuid1")
	assert.Len(t, events, 0)

	events = store.GetEvents(2, 99, "nope")
	assert.Len(t, events, 1)

	events = store.GetEvents(2, 25, "")
	assert.Len(t, events, 2)
	assert.Equal(t, int32(112), events[0].PositionTS)
}
