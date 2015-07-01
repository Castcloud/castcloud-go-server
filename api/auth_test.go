package api

import (
	"net/http"
	"testing"

	"github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/github.com/labstack/echo"
	"github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	mw := auth()
	req, _ := http.NewRequest("GET", "/not/allowed", nil)
	c := echo.NewContext(req, echo.NewResponse(nil), echo.New())

	// It should return a 401 error if no token is set
	err := mw(c).(*echo.HTTPError)
	assert.Equal(t, 401, err.Code())

	// It should return a 401 error if the token is invalid
	c.Request().Header.Set("Authorization", "stuff")
	err = mw(c).(*echo.HTTPError)
	assert.Equal(t, 401, err.Code())

	// It should return nil if the token is valid
	c.Request().Header.Set("Authorization", "token")
	assert.Nil(t, mw(c))

	// It should return nil if the path is /account/login
	c.Request().Header.Del("Authorization")
	c.Request().URL.Path = "/account/login"
	assert.Nil(t, mw(c))
}

func BenchmarkAuth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		store.GetUserByToken("token")
	}
}

func BenchmarkAuthCache(b *testing.B) {
	cache := newMemAuthCache()
	cache.set("token", store.GetUserByToken("token"))

	for i := 0; i < b.N; i++ {
		cache.get("token")
	}
}
