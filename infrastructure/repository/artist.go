package repository

import (
	"context"

	"github.com/scarlet0725/prism-api/model"
)

type Artist interface {
	CreateArtist(context.Context, *model.Artist) (*model.Artist, error)
	GetArtistByName(context.Context, string) (*model.Artist, error)
	GetArtistByID(context.Context, string) (*model.Artist, error)
	GetArtistsByIDs(context.Context, []string) ([]*model.Artist, error)
}
