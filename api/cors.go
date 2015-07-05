package api

import (
	"github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/github.com/labstack/echo"
)

func cors() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			if c.Request().Header.Get("Origin") == "" {
				return h(c)
			}

			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			if c.Request().Method == "OPTIONS" {
				c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
				c.Response().Header().Set("Access-Control-Allow-Headers", "Authorization")
				return c.NoContent(200)
			}

			return h(c)
		}
	}
}
