package api

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/labstack/echo"

	. "github.com/Castcloud/castcloud-go-server/api/schema"
)

type events struct {
	Timestamp int64   `json:"timestamp"`
	Events    []Event `json:"events"`
}

//
// GET /library/events
//
func getEvents(c echo.Context) error {
	now := time.Now().Unix()
	var err error
	var ts uint64
	since := c.Request().URL.Query().Get("since")
	if since != "" {
		ts, err = strconv.ParseUint(since, 10, 64)
		if err != nil {
			return c.NoContent(400)
		}
	}

	user := c.Get("user").(*User)
	uuid := ""
	excludeSelf := c.Request().URL.Query().Get("exclude_self")
	if excludeSelf == "true" {
		uuid = c.Get("uuid").(string)
	}

	return c.JSON(200, events{
		Timestamp: now,
		Events:    store.GetEvents(user.ID, ts, uuid),
	})
}

//
// POST /library/events
//
func addEvents(c echo.Context) error {
	data := form(c, "json")
	if data == "" {
		return c.NoContent(400)
	}

	events := []Event{}
	err := json.Unmarshal([]byte(data), &events)
	if err != nil {
		return c.NoContent(400)
	}

	user := c.Get("user").(*User)
	uuid := c.Get("uuid").(string)
	return store.AddEvents(events, user.ID, uuid)
}
