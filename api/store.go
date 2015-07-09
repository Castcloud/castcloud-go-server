package api

import (
	. "github.com/Castcloud/castcloud-go-server/api/schema"
)

type APIStore interface {
	Close() error

	GetUser(username string) *User
	GetUsers() []User
	GetUserByToken(token string) *User
	AddUser(user *User) error
	RemoveUser(username string) error
	AddClient(userid uint64, client *Client) error
	AddSubscription(userid, castid uint64) (*User, error)
	RemoveSubscription(userid, castid uint64) (*User, error)

	GetCast(id uint64) *Cast
	GetCasts() []Cast
	GetCastsByID(ids []uint64) []Cast
	GetCastByURL(url string) *Cast
	SaveCast(cast *Cast) error

	GetEpisode(id uint64) *Episode
	GetEpisodesByCast(castid uint64) []Episode
	GetEpisodesSince(ts int64) []Episode
	SaveEpisode(episode *Episode) error
	SaveEpisodes(episodes []Episode) error

	GetEvents(userid, since uint64, excludeUUID string) []Event
	AddEvents(events []Event, userid uint64, clientUUID string) error

	GetLabel(id uint64) *Label
	GetLabels(userid uint64) []Label
	SaveLabel(label *Label, userid uint64) error
	RemoveLabel(id, userid uint64) error
}
