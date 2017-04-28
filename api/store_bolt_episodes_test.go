package api

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/Castcloud/castcloud-go-server/api/schema"
)

func TestStoreGetEpisode(t *testing.T) {
	episode := store.GetEpisode(1)
	assert.NotNil(t, episode)
	assert.Equal(t, uint64(1), episode.CastID)
	assert.Equal(t, "guid", episode.GUID)
}

func TestStoreSaveEpisode(t *testing.T) {
	episode := &Episode{
		CastID: 10,
		GUID:   "guid",
	}

	// It should not save episodes with existing GUIDs
	err := store.SaveEpisode(episode)
	assert.Nil(t, err)
	assert.Equal(t, uint64(0), episode.ID)

	// It should save episodes with new GUIDs
	episode.GUID = "so_unique"
	err = store.SaveEpisode(episode)
	assert.Nil(t, err)
	assert.NotEqual(t, uint64(0), episode.ID)
	saved := store.GetEpisode(episode.ID)
	assert.NotNil(t, saved)
	assert.Equal(t, "so_unique", saved.GUID)

	episodes := []Episode{}
	episodes = append(episodes, *episode)
	episodes = append(episodes, Episode{
		CastID: 10,
		GUID:   "another_one",
	})

	err = store.SaveEpisodes(episodes)
	assert.Nil(t, err)
}
