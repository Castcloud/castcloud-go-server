package api

import (
	"encoding/binary"
	"encoding/json"
	"errors"

	"github.com/boltdb/bolt"
	"golang.org/x/crypto/bcrypt"
)

var (
	boltBucketUsers         = []byte("users")
	boltBucketUsernameIndex = []byte("index_username")
	boltBucketTokenIndex    = []byte("index_token")
	boltBucketCasts         = []byte("casts")
	boltBucketCastURLIndex  = []byte("index_cast_url")

	ErrSubscriptionExists = errors.New("Subscription already exists")
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

func (s *BoltStore) GetUser(username string) *User {
	var user *User

	s.db.View(func(tx *bolt.Tx) error {
		index := tx.Bucket(boltBucketUsernameIndex)
		id := index.Get([]byte(username))
		if id == nil {
			return nil
		}

		b := tx.Bucket(boltBucketUsers)
		user = &User{}
		return json.Unmarshal(b.Get(id), user)
	})

	return user
}

func (s *BoltStore) GetUserByToken(token string) *User {
	var user *User

	s.db.View(func(tx *bolt.Tx) error {
		index := tx.Bucket(boltBucketTokenIndex)
		id := index.Get([]byte(token))
		if id == nil {
			return nil
		}

		b := tx.Bucket(boltBucketUsers)
		user = &User{}
		return json.Unmarshal(b.Get(id), user)
	})

	return user
}

func (s *BoltStore) AddUser(user *User) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		username := []byte(user.Username)
		index := tx.Bucket(boltBucketUsernameIndex)
		if index.Get(username) != nil {
			return errors.New("Username already exists")
		}

		b := tx.Bucket(boltBucketUsers)

		var err error
		user.ID, err = b.NextSequence()
		if err != nil {
			return err
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			return err
		}
		user.Password = string(hash)

		v, err := json.Marshal(user)
		if err != nil {
			return err
		}

		id := uint64Bytes(user.ID)
		err = index.Put(username, id)
		if err != nil {
			return err
		}

		return b.Put(id, v)
	})
}

func (s *BoltStore) AddClient(userid uint64, client *Client) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(boltBucketUsers)
		id := uint64Bytes(userid)
		v := b.Get(id)
		if v == nil {
			return errors.New("User does not exist")
		}

		var user User
		json.Unmarshal(v, &user)

		user.Clients = append(user.Clients, client)

		v, err := json.Marshal(user)
		if err != nil {
			return err
		}

		index := tx.Bucket(boltBucketTokenIndex)
		err = index.Put([]byte(client.Token), id)
		if err != nil {
			return err
		}

		return b.Put(id, v)
	})
}

func (s *BoltStore) AddSubscription(userid, castid uint64) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(boltBucketUsers)
		id := uint64Bytes(userid)
		v := b.Get(id)
		if v == nil {
			return errors.New("User does not exist")
		}

		var user User
		json.Unmarshal(v, &user)

		for _, subid := range user.Subscriptions {
			if castid == subid {
				return ErrSubscriptionExists
			}
		}

		user.Subscriptions = append(user.Subscriptions, castid)

		v, err := json.Marshal(user)
		if err != nil {
			return err
		}

		return b.Put(id, v)
	})
}

func (s *BoltStore) GetCasts(ids []uint64) []Cast {
	casts := make([]Cast, len(ids))

	s.db.View(func(tx *bolt.Tx) error {
		var err error
		b := tx.Bucket(boltBucketCasts)

		for i := 0; i < len(ids); i++ {
			err = json.Unmarshal(b.Get(uint64Bytes(ids[i])), &casts[i])
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
		return json.Unmarshal(b.Get(id), cast)
	})

	return cast
}

func (s *BoltStore) SaveCast(cast *Cast) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		url := []byte(cast.URL)
		index := tx.Bucket(boltBucketCastURLIndex)
		if index.Get(url) != nil {
			return errors.New("Cast URL already exists")
		}

		b := tx.Bucket(boltBucketCasts)

		var err error
		cast.ID, err = b.NextSequence()
		if err != nil {
			return err
		}

		v, err := json.Marshal(cast)
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

func uint64Bytes(i uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, i)
	return b
}
