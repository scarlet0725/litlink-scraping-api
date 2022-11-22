package repository

import (
	"github.com/scarlet0725/prism-api/model"
)

type DB interface {
	//ユーザー関連の操作
	GetUser(id string) (*model.User, error)
	CreateUser(*model.User) (*model.User, error)
	GetUserByAPIKey(apiKey string) (*model.User, error)
	UpdateUser(*model.User) (*model.User, error)

	//イベント関連の操作
	CreateEvent(*model.Event) (*model.Event, error)
	UpdateEvent(*model.Event) (*model.Event, error)
	DeleteEvent(*model.Event) error
	GetEventsByArtistID(artistID string) ([]*model.Event, error)
	GetEventByID(ID string) (*model.Event, error)

	//アーティスト関連の操作
	CreateArtist(*model.Artist) (*model.Artist, error)
	GetArtistByName(name string) (*model.Artist, error)
	GetArtistByID(id string) (*model.Artist, error)
	GetArtistsByIDs(ids []string) ([]*model.Artist, error)
}
