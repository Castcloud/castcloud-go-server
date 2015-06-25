package api

import (
	"log"
	"path"
	"strconv"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

var (
	store Store
	crawl *crawler
)

type Config struct {
	Port    int
	DataDir string
	Debug   bool

	MaxDownloadConnections int
}

func Serve(cfg *Config) {
	var err error
	store, err = NewBoltStore(path.Join(cfg.DataDir, "store"))
	if err != nil {
		log.Fatal(err)
	}

	crawl = newCrawler()
	crawl.start(cfg.MaxDownloadConnections)

	cast := <-crawl.fetch("http://feeds.feedburner.com/BsdNowHd")
	log.Println(cast.Name, ":", cast.URL)

	r := createRouter()
	r.Run(":" + strconv.Itoa(cfg.Port))
}

func createRouter() *echo.Echo {
	r := echo.New()

	r.Use(mw.Logger())
	r.Use(mw.Recover())
	r.Use(mw.StripTrailingSlash())
	r.Use(auth())
	r.Use(mw.Gzip())

	account := r.Group("/account")
	account.Post("/login", login)
	account.Get("/ping", ping)
	/*account.Get("/settings", nil)
	account.Post("/settings", nil)
	account.Delete("/settings/:id", nil)
	account.Get("/takeout", nil)*/

	casts := r.Group("/library/casts")
	casts.Get("", getCasts)
	casts.Post("", addCast)
	/*casts.Put("/:id", nil)
	casts.Delete("/:id", nil)

	episodes := r.Group("/library")
	episodes.Get("/newepisodes", nil)
	episodes.Get("/episodes/:castid", nil)
	episodes.Get("/episode/:id", nil)
	episodes.Get("/episodes/label/:label", nil)

	events := r.Group("/library/events")
	events.Get("/", nil)
	events.Post("/", nil)

	labels := r.Group("/library/labels")
	labels.Get("/", nil)
	labels.Post("/", nil)
	labels.Put("/:id", nil)
	labels.Delete("/:id", nil)

	opml := r.Group("/library")
	opml.Get("/casts.opml", nil)
	opml.Post("/casts.opml", nil)*/

	return r
}
