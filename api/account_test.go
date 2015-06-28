package api

import (
	"encoding/json"
	"testing"

	"github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	r := createRouter()

	// It should return 400 if required params are missing
	req := testRequest(r, "POST", "/account/login", nil)
	assert.Equal(t, 400, req.send().Code)

	// It should return 200 and the correct token for an existing UUID
	req.PostForm.Set("username", "test")
	req.PostForm.Set("password", "pass")
	req.PostForm.Set("uuid", "real_unique")
	req.PostForm.Set("clientname", "Castcloud")
	res := req.send()
	assert.Equal(t, 200, res.Code)
	resb := token{}
	json.Unmarshal(res.Body.Bytes(), &resb)
	assert.Equal(t, "token", resb.Token)

	// It should return 200 and a new token if the UUID is new
	req.PostForm.Set("uuid", "nope")
	res = req.send()
	assert.Equal(t, 200, res.Code)
	json.Unmarshal(res.Body.Bytes(), &resb)
	assert.True(t, len(resb.Token) > 0)
	assert.NotEqual(t, "token", resb.Token)

	// It should return 401 if the user doesnt exist
	req.PostForm.Set("username", "bob")
	assert.Equal(t, 401, req.send().Code)

	// It should return 401 if the password is wrong
	req.PostForm.Set("username", "test")
	req.PostForm.Set("password", "not_correct")
	assert.Equal(t, 401, req.send().Code)
}

func TestPing(t *testing.T) {
	r := createRouter()

	// It should return 401 if no token is set
	req := testRequest(r, "GET", "/account/ping", nil)
	assert.Equal(t, 401, req.send().Code)

	// It should return 401 if the token is invalid
	req.Header.Set("Authorization", "apples")
	assert.Equal(t, 401, req.send().Code)

	// It should return 200 if the token is valid
	req.Header.Set("Authorization", "token")
	assert.Equal(t, 200, req.send().Code)
}

func TestCreateToken(t *testing.T) {
	token, err := createToken(32)
	assert.Nil(t, err)
	assert.Len(t, token, 64)
}
