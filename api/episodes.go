package api

import (
	"strconv"

	"github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/github.com/labstack/echo"
)

func getNewEpisodes(c *echo.Context) error {
	return nil
}

func getEpisodes(c *echo.Context) error {
	castid, err := strconv.ParseUint(c.Param("castid"), 10, 64)
	if err != nil {
		return err
	}

	return c.JSON(200, store.GetEpisodesByCast(castid))
}

func getEpisode(c *echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}

	return c.JSON(200, store.GetEpisode(id))
}

func getEpisodesByLabel(c *echo.Context) error {
	return nil
}
