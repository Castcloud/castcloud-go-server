package api

type APIStore interface {
	Close() error

	GetUser(username string) *User
	GetUsers() []User
	GetUserByToken(token string) *User
	AddUser(user *User) error
	RemoveUser(username string) error
	AddClient(userid uint64, client *Client) error
	AddSubscription(userid, castid uint64) error
	RemoveSubscription(userid, castid uint64) error

	GetCast(id uint64) *Cast
	GetCasts(ids []uint64) []Cast
	GetCastByURL(url string) *Cast
	SaveCast(cast *Cast) error
}

type User struct {
	ID            uint64
	Username      string
	Password      string
	Clients       []*Client
	Subscriptions []uint64
}

type Client struct {
	Token string
	UUID  string
	Name  string
}

type Cast struct {
	ID   uint64 `json:"id"`
	URL  string `json:"url"`
	Name string `json:"name"`
	Feed []byte `json:"feed"`
}

type Episode struct {
	ID        uint64 `json:"id"`
	CastID    uint64 `json:"castid"`
	LastEvent Event  `json:"lastevent"`
	Feed      []byte `json:"feed"`
	CrawlTS   uint64 `json:"crawlts"`
}

type Event struct {
}