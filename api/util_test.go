package api

import (
	"crypto/md5"
	"net/http"
	"testing"

	"github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/github.com/labstack/echo"
	"github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

func TestFormContains(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	c := echo.NewContext(req, echo.NewResponse(nil), echo.New())
	assert.False(t, formContains(c, "a", "b"))

	c.Request().Form.Set("a", "val")
	c.Request().Form.Set("b", "val")
	assert.True(t, formContains(c, "a", "b"))
}

func TestMD5(t *testing.T) {
	assert.Len(t, md5Hash("stuff"), md5.Size*2)
}
