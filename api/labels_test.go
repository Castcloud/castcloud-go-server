package api

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/Castcloud/castcloud-go-server/api/schema"
)

func TestGetLabels(t *testing.T) {
	r := createRouter(true)

	// It should return 200 and a list of labels with the root label
	// already added, which contains the test cast
	req := testRequest(r, "GET", "/library/labels", nil)
	req.Header.Set("Authorization", "token")
	res := req.send()
	assert.Equal(t, 200, res.Code)
	labels := []Label{}
	err := json.Unmarshal(res.Body.Bytes(), &labels)
	assert.Nil(t, err)
	assert.Len(t, labels, 1)
	assert.True(t, labels[0].Root)
}

func TestAddLabel(t *testing.T) {
	r := createRouter(true)

	// It should return 400 if the label name is not set
	req := testRequest(r, "POST", "/library/labels", nil)
	req.Header.Set("Authorization", "token")
	assert.Equal(t, 400, req.send().Code)
	req.PostForm = url.Values{}
	req.PostForm.Set("name", "label/")
	assert.Equal(t, 400, req.send().Code)

	// It should return 200 and the added label when adding a label
	req.PostForm.Set("name", "test")
	req.PostForm.Set("content", "apples")
	req.PostForm.Set("expanded", "true")
	res := req.send()
	assert.Equal(t, 200, res.Code)
	added := addedLabel{}
	err := json.Unmarshal(res.Body.Bytes(), &added)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, added.ID)
	assert.Equal(t, "test", added.Name)

	// It should return 409 if the label name already exists
	assert.Equal(t, 409, req.send().Code)

	// It should have added the label
	req.Method = "GET"
	res = req.send()
	assert.Equal(t, 200, res.Code)
	labels := []Label{}
	err = json.Unmarshal(res.Body.Bytes(), &labels)
	assert.Nil(t, err)
	assert.Len(t, labels, 2)
	assert.Contains(t, labels[0].Content, "label/2")
	assert.Equal(t, "test", labels[1].Name)
	assert.Equal(t, "apples", labels[1].Content)
	assert.True(t, labels[1].Expanded)
}

func TestUpdateLabel(t *testing.T) {
	r := createRouter(true)

	// It should return 400 if the ID is invalid
	req := testRequest(r, "PUT", "/library/labels/broken", nil)
	req.Header.Set("Authorization", "token")
	assert.Equal(t, 400, req.send().Code)

	// It should return 404 if the label doesnt exist
	req.URL.Path = "/library/labels/1881"
	assert.Equal(t, 404, req.send().Code)

	// It should return 200 when updating a label
	req.URL.Path = "/library/labels/2"
	req.PostForm = url.Values{}
	req.PostForm.Set("name", "bob")
	req.PostForm.Set("content", "cake")
	req.PostForm.Set("expanded", "no_wai")
	assert.Equal(t, 200, req.send().Code)

	// It should have updated the label
	req.Method = "GET"
	req.URL.Path = "/library/labels"
	res := req.send()
	assert.Equal(t, 200, res.Code)
	labels := []Label{}
	err := json.Unmarshal(res.Body.Bytes(), &labels)
	assert.Nil(t, err)
	assert.Len(t, labels, 2)
	assert.Equal(t, "bob", labels[1].Name)
	assert.Equal(t, "cake", labels[1].Content)
	assert.False(t, labels[1].Expanded)
}

func TestRemoveLabel(t *testing.T) {
	r := createRouter(true)

	// It should return 400 if the ID is invalid
	req := testRequest(r, "DELETE", "/library/labels/rawr", nil)
	req.Header.Set("Authorization", "token")
	assert.Equal(t, 400, req.send().Code)

	// It returns 200 when removing a label
	req.URL.Path = "/library/labels/2"
	assert.Equal(t, 200, req.send().Code)

	// It should have removed the label
	req.Method = "GET"
	req.URL.Path = "/library/labels"
	res := req.send()
	assert.Equal(t, 200, res.Code)
	labels := []Label{}
	err := json.Unmarshal(res.Body.Bytes(), &labels)
	assert.Nil(t, err)
	assert.Len(t, labels, 1)
	assert.NotContains(t, labels[0].Content, "label/2")
}
