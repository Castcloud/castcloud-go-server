package api

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/labstack/echo"
)

func TestMain(m *testing.M) {
	var err error
	store, err = NewBoltStore("/home/kenh/__cc_test_store")
	if err != nil {
		log.Fatal(err)
	}

	store.AddUser(&User{
		Username: "test",
		Password: "pass",
	})
	store.AddClient(1, &Client{
		Token: "token",
		UUID:  "real_unique",
		Name:  "Castcloud",
	})
	store.SaveCast(&Cast{
		URL:  "test.go",
		Name: "test",
	})
	store.AddSubscription(1, 1)

	crawl = newCrawler()
	crawl.start(4)

	code := m.Run()

	store.Close()
	os.Remove("/home/kenh/__cc_test_store")
	os.Exit(code)
}

type testReq struct {
	r *echo.Echo
	*http.Request
}

func testRequest(r *echo.Echo, method, url string, body io.Reader) testReq {
	req, _ := http.NewRequest(method, url, body)
	if method == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	return testReq{r: r, Request: req}
}

func (t testReq) send() *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	t.r.ServeHTTP(w, t.Request)
	return w
}

func testForm(f url.Values) *bytes.Buffer {
	return bytes.NewBufferString(f.Encode())
}

func testJSON(v interface{}) string {
	buf := &bytes.Buffer{}
	json.NewEncoder(buf).Encode(v)
	return buf.String()
}
