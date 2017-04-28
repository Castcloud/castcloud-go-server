package api

import (
	"strconv"
	"strings"

	"github.com/labstack/echo"

	. "github.com/Castcloud/castcloud-go-server/api/schema"
)

type addedLabel struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

//
// GET /library/labels
//
func getLabels(c echo.Context) error {
	user := c.Get("user").(*User)
	return c.JSON(200, store.GetLabels(user.ID))
}

//
// POST /library/labels
//
func addLabel(c echo.Context) error {
	name := strings.TrimPrefix(form(c, "name"), "label/")
	if name == "" {
		return c.String(400, "Name too short")
	}

	label := Label{
		Name:     name,
		Content:  form(c, "content"),
		Expanded: form(c, "expanded") == "true",
	}

	user := c.Get("user").(*User)
	err := store.SaveLabel(&label, user.ID)
	if err == ErrLabelExists {
		return c.String(409, "Existing label")
	} else if err != nil {
		return err
	}

	return c.JSON(200, addedLabel{
		ID:   label.ID,
		Name: label.Name,
	})
}

//
// PUT /library/labels/:id
//
func updateLabel(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.NoContent(400)
	}

	label := store.GetLabel(id)
	if label == nil {
		return c.String(404, "Label not found")
	}

	if name := form(c, "name"); name != "" {
		label.Name = name
	}
	if content := form(c, "content"); content != "" {
		label.Content = content
	}
	if expanded := form(c, "expanded"); expanded != "" {
		label.Expanded = expanded == "true"
	}

	user := c.Get("user").(*User)
	return store.SaveLabel(label, user.ID)
}

//
// DELETE /library/labels/:id
//
func removeLabel(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.NoContent(400)
	}

	user := c.Get("user").(*User)
	return store.RemoveLabel(id, user.ID)
}
