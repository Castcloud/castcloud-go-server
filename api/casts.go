package api

import (
	"github.com/labstack/echo"
)

//
// GET /library/casts
//
func getCasts(c *echo.Context) error {
	user := c.Get("user").(*User)
	return c.JSON(200, store.GetCasts(user.Subscriptions))
}

//
// POST /library/casts
//
func addCast(c *echo.Context) error {
	user := c.Get("user").(*User)
	url := form(c, "feedurl")
	cast := store.GetCastByURL(url)
	if cast == nil {
		cast = <-crawl.fetch(url)
	}

	if cast != nil {
		err := store.AddSubscription(user.ID, cast.ID)
		if err != nil && err != ErrSubscriptionExists {
			return err
		}
	}

	return c.JSON(200, cast)
}
