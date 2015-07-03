package api

import (
	"bytes"
	"encoding/binary"
	"encoding/json"

	"github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/github.com/boltdb/bolt"
)

func (s *BoltStore) GetEpisode(id uint64) *Episode {
	var episode *Episode

	s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(boltBucketEpisodes)
		v := b.Get(uint64Bytes(id))
		if v != nil {
			episode = &Episode{}
			json.Unmarshal(v, episode)
		}

		return nil
	})

	return episode
}

func (s *BoltStore) GetEpisodesByCast(castid uint64) []Episode {
	episodes := []Episode{}

	s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(boltBucketEpisodes)
		c := tx.Bucket(boltBucketEpisodeCastIDIndex).Cursor()
		prefix := uint64Bytes(castid)

		for key, id := c.Seek(prefix); bytes.HasPrefix(key, prefix); key, id = c.Next() {
			v := b.Get(id)
			episode := Episode{}
			json.Unmarshal(v, &episode)
			episodes = append(episodes, episode)
		}

		return nil
	})

	return episodes
}

func (s *BoltStore) GetEpisodesSince(ts int64) []Episode {
	episodes := []Episode{}

	s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(boltBucketEpisodes)
		c := tx.Bucket(boltBucketEpisodeCrawlTSIndex).Cursor()
		min := uint64Bytes(uint64(ts) + 1)

		for _, id := c.Seek(min); id != nil; _, id = c.Next() {
			v := b.Get(id)
			episode := Episode{}
			json.Unmarshal(v, &episode)
			episodes = append(episodes, episode)
		}

		return nil
	})

	return episodes
}

func (s *BoltStore) SaveEpisode(episode *Episode) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		return s.saveEpisode(tx, episode)
	})
}

func (s *BoltStore) SaveEpisodes(episodes []Episode) error {
	var err error

	return s.db.Update(func(tx *bolt.Tx) error {
		for _, ep := range episodes {
			err = s.saveEpisode(tx, &ep)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (s *BoltStore) saveEpisode(tx *bolt.Tx, ep *Episode) error {
	index := tx.Bucket(boltBucketEpisodeGUIDIndex)
	guid := []byte(ep.GUID)
	if index.Get(guid) != nil {
		return nil
	}

	var err error
	b := tx.Bucket(boltBucketEpisodes)
	ep.ID, err = b.NextSequence()
	if err != nil {
		return err
	}

	v, err := json.Marshal(ep)
	if err != nil {
		return err
	}

	id := uint64Bytes(ep.ID)
	err = index.Put(guid, id)
	if err != nil {
		return err
	}

	idxID := make([]byte, 16)
	binary.BigEndian.PutUint64(idxID[:8], ep.CastID)
	binary.BigEndian.PutUint64(idxID[8:], ep.ID)
	castIndex := tx.Bucket(boltBucketEpisodeCastIDIndex)
	err = castIndex.Put(idxID, id)
	if err != nil {
		return err
	}

	crawlTSIndex := tx.Bucket(boltBucketEpisodeCrawlTSIndex)
	idxID = make([]byte, 16)
	binary.BigEndian.PutUint64(idxID[:8], uint64(ep.CrawlTS))
	binary.BigEndian.PutUint64(idxID[8:], ep.ID)
	err = crawlTSIndex.Put(idxID, id)
	if err != nil {
		return err
	}

	return b.Put(id, v)
}
