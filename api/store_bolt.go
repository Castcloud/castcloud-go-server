package api

import (
	"encoding/binary"
	"errors"

	"github.com/khlieng/castcloud-go/Godeps/_workspace/src/github.com/boltdb/bolt"
)

var (
	boltBucketUsers         = []byte("users")
	boltBucketUsernameIndex = []byte("index_username")
	boltBucketTokenIndex    = []byte("index_token")
	boltBucketCasts         = []byte("casts")
	boltBucketCastURLIndex  = []byte("index_cast_url")

	ErrUsernameUnavailable  = errors.New("Username already in use")
	ErrUserNotFound         = errors.New("User does not exist")
	ErrSubscriptionExists   = errors.New("Subscription already exists")
	ErrSubsctiptionNotFound = errors.New("Subscription does not exist")
	ErrCastNotFound         = errors.New("Cast does not exist")
)

type BoltStore struct {
	db *bolt.DB
}

func NewBoltStore(path string) (*BoltStore, error) {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(boltBucketUsers)
		_, err = tx.CreateBucketIfNotExists(boltBucketUsernameIndex)
		_, err = tx.CreateBucketIfNotExists(boltBucketTokenIndex)
		_, err = tx.CreateBucketIfNotExists(boltBucketCasts)
		_, err = tx.CreateBucketIfNotExists(boltBucketCastURLIndex)
		return err
	})

	if err != nil {
		db.Close()
		return nil, err
	}

	return &BoltStore{
		db: db,
	}, nil
}

func (s *BoltStore) Close() error {
	return s.db.Close()
}

func uint64Bytes(i uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, i)
	return b
}
