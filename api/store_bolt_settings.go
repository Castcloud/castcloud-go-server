package api

import (
	"bytes"
	"encoding/binary"
	"strconv"

	"github.com/boltdb/bolt"

	. "github.com/Castcloud/castcloud-go-server/api/schema"
)

func (s *BoltStore) GetSettings(userid uint64, clientUUID string) []Setting {
	m := make(map[string]*Setting)

	s.db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(boltBucketSettings).Cursor()
		id := strconv.FormatUint(userid, 10)
		prefix := []byte(id)

		for k, v := c.Seek(prefix); bytes.HasPrefix(k, prefix); k, v = c.Next() {
			setting := &Setting{}
			setting.Unmarshal(v)
			if !setting.ClientSpecific {
				m[setting.Name] = setting
			}
		}

		prefix = []byte(id + ":" + clientUUID)

		for k, v := c.Seek(prefix); bytes.HasPrefix(k, prefix); k, v = c.Next() {
			setting := &Setting{}
			setting.Unmarshal(v)
			m[setting.Name] = setting
		}

		return nil
	})

	settings := []Setting{}
	for _, setting := range m {
		settings = append(settings, *setting)
	}

	return settings
}

func (s *BoltStore) SaveSettings(settings []Setting, userid uint64, clientUUID string) error {
	var err error

	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(boltBucketSettings)
		index := tx.Bucket(boltBucketSettingUserIDIndex)

		for _, setting := range settings {
			sid := strconv.FormatUint(userid, 10)
			if setting.ClientSpecific {
				sid += ":" + clientUUID
			}
			sid += ":" + setting.Name
			id := []byte(sid)

			v := b.Get(id)
			if v == nil {
				setting.ID, err = b.NextSequence()
				if err != nil {
					return err
				}
			} else {
				s := &Setting{}
				s.Unmarshal(v)
				setting.ID = s.ID
			}

			v, err = setting.Marshal(nil)
			if err != nil {
				return err
			}

			idxID := make([]byte, 16)
			binary.BigEndian.PutUint64(idxID[:8], userid)
			binary.BigEndian.PutUint64(idxID[8:], setting.ID)

			index.Put(idxID, id)
			b.Put(id, v)
		}

		return nil
	})
}

func (s *BoltStore) RemoveSetting(id, userid uint64) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		idxID := make([]byte, 16)
		binary.BigEndian.PutUint64(idxID[:8], userid)
		binary.BigEndian.PutUint64(idxID[8:], id)

		index := tx.Bucket(boltBucketSettingUserIDIndex)
		id := index.Get(idxID)
		if id == nil {
			return ErrSettingNotFound
		}

		err := index.Delete(idxID)
		if err != nil {
			return err
		}

		return tx.Bucket(boltBucketSettings).Delete(id)
	})
}
