package api

import (
	"log"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"

	. "github.com/Castcloud/castcloud-go-server/api/schema"
)

var (
	store     APIStore
	authCache *memAuthCache
	crawl     *crawler
	config    *Config
)

type Config struct {
	Port      int
	Dir       string
	Debug     bool
	LogFormat string

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

	port := ":" + strconv.Itoa(config.Port)
	createRouter(false).Start(port)
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

func createRouter(silent bool) *echo.Echo {
	r := echo.New()

	r.HTTPErrorHandler = errorHandler

	if !silent {
		r.Use(mw.LoggerWithConfig(mw.LoggerConfig{
			Format: config.LogFormat,
		}))
	}
	r.Use(mw.Recover())
	r.Use(cors())
	r.Use(mw.RemoveTrailingSlash())
	r.Use(auth)

	account := r.Group("/account")
	account.POST("/login", login)
	account.GET("/ping", ping)
	account.GET("/settings", getSettings)
	account.POST("/settings", setSettings)
	account.DELETE("/settings/:id", removeSetting)
	//account.Get("/takeout")

	casts := r.Group("/library/casts")
	casts.GET("", getCasts)
	casts.POST("", addCast)
	casts.PUT("/:id", renameCast)
	casts.DELETE("/:id", removeCast)

	episodes := r.Group("/library")
	episodes.GET("/newepisodes", getNewEpisodes)
	episodes.GET("/episodes/:castid", getEpisodes)
	episodes.GET("/episode/:id", getEpisode)
	episodes.GET("/episodes/label/:label", getEpisodesByLabel)

	events := r.Group("/library/events")
	events.GET("", getEvents)
	events.POST("", addEvents)

	labels := r.Group("/library/labels")
	labels.GET("", getLabels)
	labels.POST("", addLabel)
	labels.PUT("/:id", updateLabel)
	labels.DELETE("/:id", removeLabel)

	/*opml := r.Group("/library")
	opml.GET("/casts.opml")
	opml.POST("/casts.opml")*/

	// Perhaps:
	// /newepisodes           -> /episodes?since=0
	// /episodes/label/:label -> /episodes?label=label

	return r
}

func errorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	msg := http.StatusText(code)
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Error()
	}
	c.Response().Header().Set("Content-Type", "text/plain; charset=utf-8")
	c.Response().WriteHeader(code)
	c.Response().Write([]byte(msg))
}
