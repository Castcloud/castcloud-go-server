package api

import (
	"crypto/md5"
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestFormContains(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	c := echo.New().NewContext(req, nil)
	assert.False(t, formContains(c, "a", "b"))

	c.Request().Form.Set("a", "val")
	c.Request().Form.Set("b", "val")
	assert.True(t, formContains(c, "a", "b"))
}

func TestMD5(t *testing.T) {
	assert.Len(t, md5Hash("stuff"), md5.Size*2)
}
