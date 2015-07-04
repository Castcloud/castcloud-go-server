package api

import (
	"strconv"

	"github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/github.com/labstack/echo"

	. "github.com/Castcloud/castcloud-go-server/api/schema"
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
		if cast == nil {
			return c.String(500, "Could not fetch feed")
		}
	}

	user, err := store.AddSubscription(user.ID, cast.ID)
	if err != nil && err != ErrSubscriptionExists {
		return err
	}

	authCache.set(c.Get("token").(string), user)

	return c.JSON(200, cast)
}

//
// PUT /library/casts/:id
//
func renameCast(c *echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.NoContent(400)
	}

	cast := store.GetCast(id)
	if cast != nil {
		prev := cast.Name
		cast.Name = form(c, "name")
		if cast.Name != prev {
			return store.SaveCast(cast)
		}
	}

	return nil
}

//
// DELETE /library/casts/:id
//
func removeCast(c *echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.NoContent(400)
	}

	user := c.Get("user").(*User)
	user, err = store.RemoveSubscription(user.ID, id)
	authCache.set(c.Get("token").(string), user)
	return err
}
