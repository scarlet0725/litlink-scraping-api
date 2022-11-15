package repository

import (
	"github.com/scarlet0725/prism-api/model"
)

type DB interface {
	GetArtistByName(name string) (*model.Artist, error)
	//GetUserByAPIKey(apiKey string) (*model.User, error)
}
