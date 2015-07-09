package api

import (
	"log"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/github.com/labstack/echo"
	mw "github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/github.com/labstack/echo/middleware"

	. "github.com/Castcloud/castcloud-go-server/api/schema"
)

var (
	store     APIStore
	authCache *memAuthCache
	crawl     *crawler
	config    *Config
)

type Config struct {
	Port  int
	Dir   string
	Debug bool

	CrawlInterval          time.Duration
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
	authCache = newMemAuthCache()

	crawl = newCrawler(config.CrawlInterval)
	crawl.start(config.MaxDownloadConnections)

	if user := store.GetUser("test"); user == nil {
		store.AddUser(&User{
			Username: "test",
			Password: "pass",
		})
		store.AddClient(1, &Client{
			Token: "token",
			UUID:  "real_unique",
			Name:  "Castcloud",
		})
	}

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

	r.SetHTTPErrorHandler(errorHandler)

	r.Use(mw.Logger())
	r.Use(mw.Recover())
	r.Use(cors())
	r.Use(mw.StripTrailingSlash())
	r.Use(auth())

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

	events := r.Group("/library/events")
	events.Get("", getEvents)
	events.Post("", addEvents)

	labels := r.Group("/library/labels")
	labels.Get("", getLabels)
	labels.Post("", addLabel)
	labels.Put("/:id", updateLabel)
	labels.Delete("/:id", removeLabel)

	/*opml := r.Group("/library")
	opml.Get("/casts.opml")
	opml.Post("/casts.opml")*/

	return r
}

func errorHandler(err error, c *echo.Context) {
	code := http.StatusInternalServerError
	msg := http.StatusText(code)
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code()
		msg = he.Error()
	}
	c.Response().Header().Set("Content-Type", "text/plain; charset=utf-8")
	c.Response().WriteHeader(code)
	c.Response().Write([]byte(msg))
}
