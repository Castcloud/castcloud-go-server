package api

import (
	"encoding/binary"
	"errors"

	"github.com/boltdb/bolt"
)

var (
	boltBucketUsers               = []byte("users")
	boltBucketUsernameIndex       = []byte("index_username")
	boltBucketTokenIndex          = []byte("index_token")
	boltBucketSettings            = []byte("settings")
	boltBucketSettingUserIDIndex  = []byte("index_setting_userid")
	boltBucketCasts               = []byte("casts")
	boltBucketCastURLIndex        = []byte("index_cast_url")
	boltBucketEpisodes            = []byte("episodes")
	boltBucketEpisodeCastIDIndex  = []byte("index_episode_castid")
	boltBucketEpisodeGUIDIndex    = []byte("index_episode_guid")
	boltBucketEpisodeCrawlTSIndex = []byte("index_episode_crawlts")
	boltBucketEvents              = []byte("events")
	boltBucketLabels              = []byte("labels")
	boltBucketLabelUserIDIndex    = []byte("index_label_userid")
	boltBucketLabelNameIndex      = []byte("index_label_name")
	boltBucketLabelRootIndex      = []byte("index_label_root")

	ErrUsernameUnavailable  = errors.New("Username already in use")
	ErrUserNotFound         = errors.New("User does not exist")
	ErrSubscriptionExists   = errors.New("Subscription already exists")
	ErrSubsctiptionNotFound = errors.New("Subscription does not exist")
	ErrSettingNotFound      = errors.New("Setting does not exist")
	ErrCastNotFound         = errors.New("Cast does not exist")
	ErrLabelExists          = errors.New("Label already exists")
)

type BoltStore struct {
	db *bolt.DB
}

func NewBoltStore(path string) (*BoltStore, error) {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}

	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists(boltBucketUsers)
		tx.CreateBucketIfNotExists(boltBucketUsernameIndex)
		tx.CreateBucketIfNotExists(boltBucketTokenIndex)
		tx.CreateBucketIfNotExists(boltBucketSettings)
		tx.CreateBucketIfNotExists(boltBucketSettingUserIDIndex)
		tx.CreateBucketIfNotExists(boltBucketCasts)
		tx.CreateBucketIfNotExists(boltBucketCastURLIndex)
		tx.CreateBucketIfNotExists(boltBucketEpisodes)
		tx.CreateBucketIfNotExists(boltBucketEpisodeCastIDIndex)
		tx.CreateBucketIfNotExists(boltBucketEpisodeGUIDIndex)
		tx.CreateBucketIfNotExists(boltBucketEpisodeCrawlTSIndex)
		tx.CreateBucketIfNotExists(boltBucketEvents)
		tx.CreateBucketIfNotExists(boltBucketLabels)
		tx.CreateBucketIfNotExists(boltBucketLabelUserIDIndex)
		tx.CreateBucketIfNotExists(boltBucketLabelNameIndex)
		tx.CreateBucketIfNotExists(boltBucketLabelRootIndex)
		return nil
	})

	return &BoltStore{
		db: db,
	}, nil
}

func (s *BoltStore) Close() error {
	return s.db.Close()
}

func uint64Bytes(i uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)
	return b
}
