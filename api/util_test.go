package api

import (
	"net/http"
	"testing"

	"github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/github.com/labstack/echo"
	"github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

func TestFormContains(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	c := echo.NewContext(req, echo.NewResponse(nil), echo.New())
	assert.False(t, formContains(c, "a", "b"))

	c.Request().PostForm.Set("a", "val")
	c.Request().PostForm.Set("b", "val")
	assert.True(t, formContains(c, "a", "b"))
}
