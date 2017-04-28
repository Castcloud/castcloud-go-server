package api

import (
	"github.com/labstack/echo"
)

func cors() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().Header.Get("Origin") == "" {
				return next(c)
			}

			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			if c.Request().Method == "OPTIONS" {
				c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
				c.Response().Header().Set("Access-Control-Allow-Headers", "Authorization")
				return c.NoContent(200)
			}

			return next(c)
		}
	}
}
