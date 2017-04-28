package api

import (
	"strconv"
	"time"

	"github.com/labstack/echo"

	. "github.com/Castcloud/castcloud-go-server/api/schema"
)

type newEpisodes struct {
	Timestamp int64     `json:"timestamp"`
	Episodes  []Episode `json:"episodes"`
}

//
// GET /library/newepisodes
//
func getNewEpisodes(c echo.Context) error {
	now := time.Now().Unix()
	user := c.Get("user").(*User)
	since := c.Request().URL.Query().Get("since")
	if since == "" {
		return c.JSON(200, newEpisodes{
			Timestamp: now,
			Episodes:  store.GetEpisodesSince(0, user.Subscriptions),
		})
	}

	ts, err := strconv.ParseInt(since, 10, 64)
	if err != nil {
		return c.NoContent(400)
	}

	return c.JSON(200, newEpisodes{
		Timestamp: now,
		Episodes:  store.GetEpisodesSince(ts, user.Subscriptions),
	})
}

//
// GET /library/episodes/:castid
//
func getEpisodes(c echo.Context) error {
	castid, err := strconv.ParseUint(c.Param("castid"), 10, 64)
	if err != nil {
		return c.NoContent(400)
	}

	return c.JSON(200, store.GetEpisodesByCast(castid))
}

//
// GET /library/episode/:id
//
func getEpisode(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.NoContent(400)
	}

	return c.JSON(200, store.GetEpisode(id))
}

//
// GET /library/episodes/label/:label
//
func getEpisodesByLabel(c echo.Context) error {
	return nil
}
