package api

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/github.com/labstack/echo"
	"github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/golang.org/x/crypto/bcrypt"
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
