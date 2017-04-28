package api

import (
	"encoding/binary"

	"github.com/boltdb/bolt"

	. "github.com/Castcloud/castcloud-go-server/api/schema"
)

func (s *BoltStore) GetCast(id uint64) *Cast {
	var cast *Cast

	s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(boltBucketCasts)
		id := uint64Bytes(id)
		v := b.Get(id)
		if v != nil {
			cast = &Cast{}
			cast.Unmarshal(v)
			cast.DecodeFeed()
		}

		return nil
	})

	return cast
}

func (s *BoltStore) GetCasts() []Cast {
	casts := []Cast{}
	cast := &Cast{}

	s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(boltBucketCasts)

		return b.ForEach(func(k, v []byte) error {
			cast.Unmarshal(v)
			casts = append(casts, *cast)
			return nil
		})
	})

	return casts
}

func (s *BoltStore) GetCastsByID(ids []uint64) []Cast {
	casts := make([]Cast, len(ids))

	s.db.View(func(tx *bolt.Tx) error {
		var err error
		b := tx.Bucket(boltBucketCasts)

		for i := 0; i < len(ids); i++ {
			_, err = (&casts[i]).Unmarshal(b.Get(uint64Bytes(ids[i])))
			casts[i].DecodeFeed()
		}

		return err
	})

	return casts
}

func (s *BoltStore) GetCastByURL(url string) *Cast {
	var cast *Cast

	s.db.View(func(tx *bolt.Tx) error {
		index := tx.Bucket(boltBucketCastURLIndex)
		id := index.Get([]byte(url))
		if id == nil {
			return nil
		}

		b := tx.Bucket(boltBucketCasts)
		cast = &Cast{}
		cast.Unmarshal(b.Get(id))
		cast.DecodeFeed()
		return nil
	})

	return cast
}

func (s *BoltStore) SaveCast(cast *Cast) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		var err error
		b := tx.Bucket(boltBucketCasts)
		index := tx.Bucket(boltBucketCastURLIndex)
		url := []byte(cast.URL)

		idxID := index.Get(url)
		if idxID != nil {
			cast.ID = binary.BigEndian.Uint64(idxID)
		} else if cast.ID == 0 {
			cast.ID, err = b.NextSequence()
			if err != nil {
				return err
			}
		}

		cast.EncodeFeed()
		v, err := cast.Marshal(nil)
		if err != nil {
			return err
		}

		id := uint64Bytes(cast.ID)
		err = index.Put(url, id)
		if err != nil {
			return err
		}

		return b.Put(id, v)
	})
}
