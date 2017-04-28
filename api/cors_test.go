package api

import (
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestCORS(t *testing.T) {
	mw := cors()
	req := testRequest(nil, "GET", "/", nil)
	res := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req.Request, echo.NewResponse(res, e))
	called := false

	next := func(c echo.Context) error {
		called = true
		return nil
	}

	// It calls the next middleware when no origin is set
	h := mw(next)
	h(c)
	assert.True(t, called)
	assert.Empty(t, res.Header().Get("Access-Control-Allow-Origin"))

	// It sets CORS headers and calls the next middleware when the origin is set
	req.Header.Set("Origin", "china")
	called = false
	h(c)
	assert.True(t, called)
	assert.NotEmpty(t, res.Header().Get("Access-Control-Allow-Origin"))

	// It sets CORS headers, ends the middleware chain and
	// returns 200 when receiving a preflight request
	req.Method = "OPTIONS"
	res = httptest.NewRecorder()
	c = e.NewContext(req.Request, echo.NewResponse(res, e))
	res.Code = 0
	called = false
	h(c)
	assert.False(t, called)
	assert.Equal(t, 200, res.Code)
	assert.NotEmpty(t, res.Header().Get("Access-Control-Allow-Origin"))
	assert.NotEmpty(t, res.Header().Get("Access-Control-Allow-Methods"))
	assert.NotEmpty(t, res.Header().Get("Access-Control-Allow-Headers"))
}
