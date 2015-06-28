package api

import (
	"github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/github.com/labstack/echo"
)

func form(c *echo.Context, key string) string {
	return c.Request().PostFormValue(key)
}

func formContains(c *echo.Context, keys ...string) bool {
	for _, key := range keys {
		if form(c, key) == "" {
			return false
		}
	}
	return true
}
