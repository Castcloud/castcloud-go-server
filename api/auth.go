package api

import (
	"github.com/khlieng/castcloud-go/Godeps/_workspace/src/github.com/labstack/echo"
)

func auth() echo.HandlerFunc {
	return func(c *echo.Context) error {
		if c.Request().URL.Path != "/account/login" {
			token := c.Request().Header.Get("Authorization")
			if token == "" {
				return echo.NewHTTPError(401)
			}

			user := store.GetUserByToken(token)
			if user == nil {
				return echo.NewHTTPError(401)
			}

			c.Set("user", user)
		}

		return nil
	}
}
