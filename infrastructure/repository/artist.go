package repository

import "github.com/scarlet0725/prism-api/model"

type Artist interface {
	CreateArtist(*model.Artist) (*model.Artist, error)
	GetArtistByName(name string) (*model.Artist, error)
	GetArtistByID(id string) (*model.Artist, error)
	GetArtistsByIDs(ids []string) ([]*model.Artist, error)
}
