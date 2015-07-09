package api

import (
	"bytes"
	"encoding/binary"
	"strconv"
	"strings"

	"github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/github.com/boltdb/bolt"

	. "github.com/Castcloud/castcloud-go-server/api/schema"
)

func (s *BoltStore) GetLabel(id uint64) *Label {
	var label *Label

	s.db.View(func(tx *bolt.Tx) error {
		label = s.getLabel(tx, uint64Bytes(id))
		return nil
	})

	return label
}

func (s *BoltStore) GetLabels(userid uint64) []Label {
	labels := []Label{}

	s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(boltBucketLabels)
		c := tx.Bucket(boltBucketLabelUserIDIndex).Cursor()
		prefix := uint64Bytes(userid)

		for key, id := c.Seek(prefix); bytes.HasPrefix(key, prefix); key, id = c.Next() {
			v := b.Get(id)
			label := &Label{}
			label.UnmarshalMsg(v)
			labels = append(labels, *label)
		}

		return nil
	})

	return labels
}

func (s *BoltStore) SaveLabel(label *Label, userid uint64) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		index := tx.Bucket(boltBucketLabelNameIndex)
		idxNameID := []byte(strconv.FormatUint(userid, 10) + ":" + label.Name)
		if index.Get(idxNameID) != nil {
			return ErrLabelExists
		}

		var err error
		b := tx.Bucket(boltBucketLabels)
		if label.ID == 0 {
			label.ID, err = b.NextSequence()
			if err != nil {
				return err
			}
		}

		v, err := label.MarshalMsg(nil)
		if err != nil {
			return err
		}

		id := uint64Bytes(label.ID)
		err = index.Put(idxNameID, id)
		if err != nil {
			return err
		}

		idxID := make([]byte, 16)
		binary.BigEndian.PutUint64(idxID[:8], userid)
		binary.BigEndian.PutUint64(idxID[8:], label.ID)
		index = tx.Bucket(boltBucketLabelUserIDIndex)
		err = index.Put(idxID, id)
		if err != nil {
			return err
		}

		s.addToRootLabel(tx, "label", label.ID, idxID[:8])

		return b.Put(id, v)
	})
}

func (s *BoltStore) RemoveLabel(id, userid uint64) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		idxID := make([]byte, 16)
		binary.BigEndian.PutUint64(idxID[:8], userid)
		binary.BigEndian.PutUint64(idxID[8:], id)

		index := tx.Bucket(boltBucketLabelUserIDIndex)
		err := index.Delete(idxID)
		if err != nil {
			return err
		}

		label := s.getLabel(tx, idxID[8:])
		idxNameID := []byte(strconv.FormatUint(userid, 10) + ":" + label.Name)
		index = tx.Bucket(boltBucketLabelNameIndex)
		err = index.Delete(idxNameID)
		if err != nil {
			return err
		}

		s.removeFromRootLabel(tx, "label", label.ID, idxID[:8])

		b := tx.Bucket(boltBucketLabels)
		return b.Delete(idxID[8:])
	})
}

func (s *BoltStore) getLabel(tx *bolt.Tx, id []byte) *Label {
	b := tx.Bucket(boltBucketLabels)
	v := b.Get(id)
	if v != nil {
		label := &Label{}
		label.UnmarshalMsg(v)
		return label
	}
	return nil
}

func (s *BoltStore) addToRootLabel(tx *bolt.Tx, itemType string, itemID uint64, userid []byte) {
	var err error
	label := itemType + "/" + strconv.FormatUint(itemID, 10)
	b := tx.Bucket(boltBucketLabels)
	index := tx.Bucket(boltBucketLabelRootIndex)
	id := index.Get(userid)
	if id != nil {
		v := b.Get(id)
		root := &Label{}
		root.UnmarshalMsg(v)
		if strings.Contains(root.Content, label) {
			return
		}

		if root.Content != "" {
			root.Content += ","
		}
		root.Content += label

		v, err = root.MarshalMsg(nil)
		if err != nil {
			return
		}

		b.Put(id, v)
	} else {
		root := &Label{
			Name:    "root",
			Content: itemType + "/" + strconv.FormatUint(itemID, 10),
			Root:    true,
		}

		root.ID, err = b.NextSequence()
		if err != nil {
			return
		}

		v, err := root.MarshalMsg(nil)
		if err != nil {
			return
		}

		id = uint64Bytes(root.ID)
		index.Put(userid, id)

		idxID := []byte{}
		idxID = append(idxID, userid...)
		idxID = append(idxID, id...)
		index = tx.Bucket(boltBucketLabelUserIDIndex)
		err = index.Put(idxID, id)
		if err != nil {
			return
		}

		b.Put(id, v)
	}
}

func (s *BoltStore) removeFromRootLabel(tx *bolt.Tx, itemType string, itemID uint64, userid []byte) {
	index := tx.Bucket(boltBucketLabelRootIndex)
	id := index.Get(userid)
	if id != nil {
		label := itemType + "/" + strconv.FormatUint(itemID, 10)
		b := tx.Bucket(boltBucketLabels)
		v := b.Get(id)
		root := &Label{}
		root.UnmarshalMsg(v)

		if strings.Contains(root.Content, ","+label) {
			root.Content = strings.Replace(root.Content, ","+label, "", 1)
		} else if strings.HasPrefix(root.Content, label+",") {
			root.Content = strings.Replace(root.Content, label+",", "", 1)
		} else {
			root.Content = strings.Replace(root.Content, label, "", 1)
		}

		v, err := root.MarshalMsg(nil)
		if err != nil {
			return
		}

		b.Put(id, v)
	}
}
