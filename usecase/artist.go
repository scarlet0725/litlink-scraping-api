package usecase

import (
	"context"

	"github.com/scarlet0725/prism-api/framework"
	"github.com/scarlet0725/prism-api/infrastructure/repository"
	"github.com/scarlet0725/prism-api/model"
)

type Artist interface {
	CreateArtist(ctx context.Context, artist *model.Artist) (*model.Artist, error)
	//UpdateArtist(artist *model.Artist) (*model.Artist, error)
	//GetArtistByName(name string) ([]*model.Artist, error)
	GetArtistByID(ctx context.Context, id string) (*model.Artist, error)
}

type artistUsecase struct {
	db     repository.Artist
	random framework.RandomID
}

func NewArtistUsecase(db repository.Artist, r framework.RandomID) Artist {
	return &artistUsecase{
		db:     db,
		random: r,
	}
}

func (a *artistUsecase) CreateArtist(ctx context.Context, artist *model.Artist) (*model.Artist, error) {
	id := a.random.Generate(artistIDLength)

	artist.ArtistID = id

	result, err := a.db.CreateArtist(ctx, artist)
	if err != nil {
		return nil, &model.AppError{
			Msg:  "Failed to create artist",
			Code: 500,
		}
	}
	return result, nil
}

func (a *artistUsecase) GetArtistByID(ctx context.Context, id string) (*model.Artist, error) {
	result, err := a.db.GetArtistByID(ctx, id)
	if err != nil {
		return nil, &model.AppError{
			Msg:  "Failed to get artist",
			Code: 500,
		}
	}
	return result, nil
}
