package api

import (
	"testing"

	"github.com/khlieng/castcloud-go/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {
	err := store.AddUser(&User{
		Username: "added",
		Password: "pass",
	})
	assert.Nil(t, err)

	user := store.GetUser("added")
	assert.NotNil(t, user)
	assert.Equal(t, "added", user.Username)
	assert.NotEmpty(t, user.Password)

	err = store.AddUser(&User{Username: "added"})
	assert.Equal(t, ErrUsernameUnavailable, err)
}

func TestRemoveUser(t *testing.T) {
	err := store.AddUser(&User{Username: "remove_me"})
	assert.Nil(t, err)
	err = store.RemoveUser("remove_me")
	assert.Nil(t, err)
	assert.Nil(t, store.GetUser("remove_me"))

	err = store.RemoveUser("not_a_user")
	assert.Equal(t, ErrUserNotFound, err)
}
