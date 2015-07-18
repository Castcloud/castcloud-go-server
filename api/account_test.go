package api

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/github.com/stretchr/testify/assert"

	. "github.com/Castcloud/castcloud-go-server/api/schema"
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

func TestGetSettings(t *testing.T) {
	r := createRouter()

	// It should return 200 and a list of settings
	req := testRequest(r, "GET", "/account/settings", nil)
	req.Header.Set("Authorization", "token")
	res := req.send()
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, testJSON([]Setting{}), res.Body.String())
}

func TestSetSettings(t *testing.T) {
	r := createRouter()

	// It should return 400 when there is no json in the body
	req := testRequest(r, "POST", "/account/settings", nil)
	req.Header.Set("Authorization", "token")
	assert.Equal(t, 400, req.send().Code)

	// It should return 400 if the json is invalid
	req.PostForm = url.Values{}
	req.PostForm.Set("json", "not_json")
	assert.Equal(t, 400, req.send().Code)

	// It should return 200 when adding a new setting
	req.PostForm.Set("json", testJSON([]Setting{
		Setting{
			Name:  "test",
			Value: "a",
		},
	}))
	assert.Equal(t, 200, req.send().Code)

	// It should have added the setting
	settings := checkSettings(t, req)
	assert.Contains(t, settings, Setting{
		ID:    1,
		Name:  "test",
		Value: "a",
	})

	// It should return 200 when updating a setting
	req.Method = "POST"
	req.PostForm.Set("json", testJSON([]Setting{
		Setting{
			Name:  "test",
			Value: "b",
		},
	}))
	assert.Equal(t, 200, req.send().Code)

	// It should have updated the setting
	settings = checkSettings(t, req)
	assert.Contains(t, settings, Setting{
		ID:    1,
		Name:  "test",
		Value: "b",
	})

	// It should return 200 when doing a clientspecific setting update
	req.Method = "POST"
	req.PostForm.Set("json", testJSON([]Setting{
		Setting{
			Name:           "test",
			Value:          "c",
			ClientSpecific: true,
		},
	}))
	assert.Equal(t, 200, req.send().Code)

	// It should have updated the setting and should not return the
	// old non-specific version of it
	settings = checkSettings(t, req)
	assert.Contains(t, settings, Setting{
		ID:             2,
		Name:           "test",
		Value:          "c",
		ClientSpecific: true,
	})
	assert.NotContains(t, settings, Setting{
		ID:    1,
		Name:  "test",
		Value: "b",
	})

	// It should return the old non-specific setting to other clients
	req.Header.Set("Authorization", "evtest1")
	settings = checkSettings(t, req)
	assert.Contains(t, settings, Setting{
		ID:    1,
		Name:  "test",
		Value: "b",
	})
	assert.NotContains(t, settings, Setting{
		ID:             2,
		Name:           "test",
		Value:          "c",
		ClientSpecific: true,
	})
}

func checkSettings(t *testing.T, req testReq) []Setting {
	req.Method = "GET"
	res := req.send()
	assert.Equal(t, 200, res.Code)
	settings := []Setting{}
	err := json.Unmarshal(res.Body.Bytes(), &settings)
	assert.Nil(t, err)
	return settings
}

func TestRemoveSetting(t *testing.T) {
	r := createRouter()

	// It should return 400 if the id is invalid
	req := testRequest(r, "DELETE", "/account/settings/cake", nil)
	req.Header.Set("Authorization", "token")
	assert.Equal(t, 400, req.send().Code)

	// It should return 404 when the setting does not exist
	req.URL.Path = "/account/settings/1881"
	assert.Equal(t, 404, req.send().Code)

	// It should return 200 when removing a setting
	req.URL.Path = "/account/settings/1"
	assert.Equal(t, 200, req.send().Code)

	// It should have removed the setting
	req.Method = "GET"
	req.URL.Path = "/account/settings"
	res := req.send()
	assert.Equal(t, 200, res.Code)
	settings := []Setting{}
	err := json.Unmarshal(res.Body.Bytes(), &settings)
	assert.Nil(t, err)

	for _, setting := range settings {
		if setting.ID == 1 {
			t.Error("Setting with ID 1 did not get removed")
		}
	}
}

func TestCreateToken(t *testing.T) {
	token, err := createToken(32)
	assert.Nil(t, err)
	assert.Len(t, token, 64)
}

func BenchmarkPing(b *testing.B) {
	r := createRouter()
	req := testRequest(r, "GET", "/account/ping", nil)
	req.Header.Set("Authorization", "token")

	for i := 0; i < b.N; i++ {
		req.send()
	}
}
