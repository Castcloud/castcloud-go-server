package api

import (
	"log"
	"path"
	"strconv"

	"github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/github.com/labstack/echo"
	mw "github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/github.com/labstack/echo/middleware"
)

var (
	store  APIStore
	crawl  *crawler
	config *Config
)

type Config struct {
	Port  int
	Dir   string
	Debug bool

	MaxDownloadConnections int
}

func Store() APIStore {
	openStore(path.Join(config.Dir, "store"))
	return store
}

func Configure(cfg *Config) {
	config = cfg
}

func Serve() {
	openStore(path.Join(config.Dir, "store"))

	crawl = newCrawler()
	crawl.start(config.MaxDownloadConnections)

	store.AddUser(&User{
		Username: "test",
		Password: "pass",
	})
	store.AddClient(1, &Client{
		Token: "token",
		UUID:  "real_unique",
		Name:  "Castcloud",
	})

	log.Println("API listening on port", config.Port)
	createRouter().Run(":" + strconv.Itoa(config.Port))
}

func openStore(p string) {
	if store == nil {
		var err error
		store, err = NewBoltStore(p)
		if err != nil {
			log.Fatal(err)
		}
	}
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
	/*account.Get("/settings")
	account.Post("/settings")
	account.Delete("/settings/:id")
	account.Get("/takeout")*/

	casts := r.Group("/library/casts")
	casts.Get("", getCasts)
	casts.Post("", addCast)
	casts.Put("/:id", renameCast)
	casts.Delete("/:id", removeCast)

	// Perhaps:
	// /newepisodes           -> /episodes?since=0
	// /episodes/label/:label -> /episodes?label=label
	episodes := r.Group("/library")
	episodes.Get("/newepisodes", getNewEpisodes)
	episodes.Get("/episodes/:castid", getEpisodes)
	episodes.Get("/episode/:id", getEpisode)
	episodes.Get("/episodes/label/:label", getEpisodesByLabel)

	/*events := r.Group("/library/events")
	events.Get("")
	events.Post("")

	labels := r.Group("/library/labels")
	labels.Get("")
	labels.Post("")
	labels.Put("/:id")
	labels.Delete("/:id")

	opml := r.Group("/library")
	opml.Get("/casts.opml")
	opml.Post("/casts.opml")*/

	return r
}
