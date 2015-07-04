package api

import (
	"sync"

	"github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/github.com/labstack/echo"

	. "github.com/Castcloud/castcloud-go-server/api/schema"
)

func auth() echo.HandlerFunc {
	return func(c *echo.Context) error {
		if c.Request().URL.Path != "/account/login" {
			token := c.Request().Header.Get("Authorization")
			if token == "" {
				return echo.NewHTTPError(401)
			}

			user := authCache.get(token)
			if user != nil {
				c.Set("user", user)
				c.Set("token", token)
				return nil
			}

			user = store.GetUserByToken(token)
			if user == nil {
				return echo.NewHTTPError(401)
			}

			authCache.set(token, user)
			c.Set("user", user)
			c.Set("token", token)
		}

		return nil
	}
}

type memAuthCache struct {
	users map[string]*User
	lock  sync.Mutex
}

func newMemAuthCache() *memAuthCache {
	return &memAuthCache{
		users: make(map[string]*User),
	}
}

func (c *memAuthCache) get(token string) *User {
	c.lock.Lock()
	user := c.users[token]
	c.lock.Unlock()
	return user
}

func (c *memAuthCache) set(token string, user *User) {
	c.lock.Lock()
	c.users[token] = user
	c.lock.Unlock()
}
