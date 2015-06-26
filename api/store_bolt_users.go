package api

import (
	"encoding/json"

	"github.com/khlieng/castcloud-go/Godeps/_workspace/src/github.com/boltdb/bolt"
	"github.com/khlieng/castcloud-go/Godeps/_workspace/src/golang.org/x/crypto/bcrypt"
)

func (s *BoltStore) GetUser(username string) *User {
	var user *User

	s.db.View(func(tx *bolt.Tx) error {
		index := tx.Bucket(boltBucketUsernameIndex)
		id := index.Get([]byte(username))
		if id == nil {
			return nil
		}

		b := tx.Bucket(boltBucketUsers)
		v := b.Get(id)
		if v != nil {
			user = &User{}
			json.Unmarshal(v, user)
		}

		return nil
	})

	return user
}

func (s *BoltStore) GetUsers() []User {
	users := []User{}
	user := User{}

	s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(boltBucketUsers)

		return b.ForEach(func(k, v []byte) error {
			json.Unmarshal(v, &user)
			users = append(users, user)
			return nil
		})
	})

	return users
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
	username := []byte(user.Username)

	err := s.db.View(func(tx *bolt.Tx) error {
		index := tx.Bucket(boltBucketUsernameIndex)
		if index.Get(username) != nil {
			return ErrUsernameUnavailable
		}
		return nil
	})
	if err != nil {
		return err
	}

	if user.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			return err
		}
		user.Password = string(hash)
	}

	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(boltBucketUsers)

		var err error
		user.ID, err = b.NextSequence()
		if err != nil {
			return err
		}

		v, err := json.Marshal(user)
		if err != nil {
			return err
		}

		index := tx.Bucket(boltBucketUsernameIndex)
		id := uint64Bytes(user.ID)
		err = index.Put(username, id)
		if err != nil {
			return err
		}

		return b.Put(id, v)
	})
}

func (s *BoltStore) RemoveUser(username string) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		index := tx.Bucket(boltBucketUsernameIndex)
		user := []byte(username)
		id := index.Get(user)
		if id == nil {
			return ErrUserNotFound
		}

		index.Delete(user)
		return tx.Bucket(boltBucketUsers).Delete(id)
	})
}

func (s *BoltStore) AddClient(userid uint64, client *Client) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(boltBucketUsers)
		id := uint64Bytes(userid)
		v := b.Get(id)
		if v == nil {
			return ErrUserNotFound
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
		b := tx.Bucket(boltBucketCasts)
		id := uint64Bytes(castid)
		v := b.Get(id)
		if v == nil {
			return ErrCastNotFound
		}

		b = tx.Bucket(boltBucketUsers)
		id = uint64Bytes(userid)
		v = b.Get(id)
		if v == nil {
			return ErrUserNotFound
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

func (s *BoltStore) RemoveSubscription(userid, castid uint64) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(boltBucketCasts)
		id := uint64Bytes(castid)
		v := b.Get(id)
		if v == nil {
			return ErrCastNotFound
		}

		b = tx.Bucket(boltBucketUsers)
		id = uint64Bytes(userid)
		v = b.Get(id)
		if v == nil {
			return ErrUserNotFound
		}

		var user User
		json.Unmarshal(v, &user)

		for i, subid := range user.Subscriptions {
			if castid == subid {
				user.Subscriptions = append(user.Subscriptions[:i], user.Subscriptions[i+1:]...)

				v, err := json.Marshal(user)
				if err != nil {
					return err
				}

				return b.Put(id, v)
			}
		}

		return ErrSubsctiptionNotFound
	})
}
