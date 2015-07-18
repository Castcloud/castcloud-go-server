package api

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"strconv"

	"github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/github.com/labstack/echo"
	"github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/golang.org/x/crypto/bcrypt"

	. "github.com/Castcloud/castcloud-go-server/api/schema"
)

type token struct {
	Token string `json:"token"`
}

//
// POST /account/login
//
func login(c *echo.Context) error {
	if !formContains(c, "username", "password", "uuid", "clientname") {
		return echo.NewHTTPError(400)
	}

	user := store.GetUser(form(c, "username"))
	if user == nil {
		return echo.NewHTTPError(401)
	}

	if correctPassword(user.Password, form(c, "password")) {
		uuid := form(c, "uuid")
		for _, client := range user.Clients {
			if uuid == client.UUID {
				return c.JSON(200, token{client.Token})
			}
		}

		t, err := createToken(32)
		if err != nil {
			return err
		}

		err = store.AddClient(user.ID, &Client{
			Token: t,
			UUID:  uuid,
			Name:  form(c, "clientname"),
		})
		if err != nil {
			return err
		}

		return c.JSON(200, token{t})
	}

	return echo.NewHTTPError(401)
}

//
// GET /account/ping
//
func ping(c *echo.Context) error {
	// This function getting called means there was a valid token
	return nil
}

//
// GET /account/settings
//
func getSettings(c *echo.Context) error {
	user := c.Get("user").(*User)
	uuid := c.Get("uuid").(string)
	return c.JSON(200, store.GetSettings(user.ID, uuid))
}

//
// POST /account/settings
//
func setSettings(c *echo.Context) error {
	data := form(c, "json")
	if data == "" {
		return c.NoContent(400)
	}

	settings := []Setting{}
	err := json.Unmarshal([]byte(data), &settings)
	if err != nil {
		return c.NoContent(400)
	}

	user := c.Get("user").(*User)
	uuid := c.Get("uuid").(string)
	return store.SaveSettings(settings, user.ID, uuid)
}

//
// DELETE /account/settings/:id
//
func removeSetting(c *echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.NoContent(400)
	}

	user := c.Get("user").(*User)
	err = store.RemoveSetting(id, user.ID)
	if err == ErrSettingNotFound {
		return c.String(404, "Setting not found")
	}

	return err
}

func createToken(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

func correctPassword(hash, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass)) == nil
}
