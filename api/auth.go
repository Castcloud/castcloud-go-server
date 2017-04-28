package api

import (
	"sync"

	"github.com/labstack/echo"

	. "github.com/Castcloud/castcloud-go-server/api/schema"
)

func auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().URL.Path != "/account/login" {
			token := c.Request().Header.Get("Authorization")
			if token == "" {
				return echo.NewHTTPError(401, "Bad token")
			}

			user := authCache.get(token)
			if user != nil {
				c.Set("user", user)
				c.Set("token", token)
				c.Set("uuid", user.UUID(token))
				return next(c)
			}

			user = store.GetUserByToken(token)
			if user == nil {
				return echo.NewHTTPError(401, "Bad token")
			}

			authCache.set(token, user)
			c.Set("user", user)
			c.Set("token", token)
			c.Set("uuid", user.UUID(token))
		}

		return next(c)
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
