package api

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/labstack/echo"
)

func form(c echo.Context, key string) string {
	return c.FormValue(key)
}

func formContains(c echo.Context, keys ...string) bool {
	for _, key := range keys {
		if form(c, key) == "" {
			return false
		}
	}
	return true
}

func md5Hash(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}
