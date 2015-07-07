package api

import (
	"bytes"
	"encoding/binary"

	"github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/github.com/boltdb/bolt"

	. "github.com/Castcloud/castcloud-go-server/api/schema"
)

func (s *BoltStore) GetEvents(userid, since uint64, excludeUUID string) []Event {
	events := []Event{}
	event := &Event{}

	s.db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(boltBucketEvents).Cursor()
		start := make([]byte, 16)
		binary.BigEndian.PutUint64(start[:8], userid)
		binary.BigEndian.PutUint64(start[8:], since+1)
		prefix := start[:8]

		if excludeUUID != "" {
			for k, v := c.Seek(start); bytes.HasPrefix(k, prefix); k, v = c.Next() {
				event.UnmarshalMsg(v)
				if event.ClientUUID != excludeUUID {
					events = append(events, *event)
				}
			}
		} else {
			for k, v := c.Seek(start); bytes.HasPrefix(k, prefix); k, v = c.Next() {
				event.UnmarshalMsg(v)
				events = append(events, *event)
			}
		}
		return nil
	})

	return events
}

func (s *BoltStore) AddEvents(events []Event, userid uint64, clientUUID string) error {
	var err error

	return s.db.Update(func(tx *bolt.Tx) error {
		for _, event := range events {
			event.ClientUUID = clientUUID
			err = s.addEvent(tx, &event, userid)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (s *BoltStore) addEvent(tx *bolt.Tx, event *Event, userid uint64) error {
	b := tx.Bucket(boltBucketEvents)
	eventID, err := b.NextSequence()
	if err != nil {
		return err
	}

	v, err := event.MarshalMsg(nil)
	if err != nil {
		return err
	}

	id := make([]byte, 24)
	binary.BigEndian.PutUint64(id[:8], userid)
	binary.BigEndian.PutUint64(id[8:16], event.ClientTS)
	binary.BigEndian.PutUint64(id[16:], eventID)

	return b.Put(id, v)
}
