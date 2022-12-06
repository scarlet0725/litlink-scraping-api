package usecase

import (
	"github.com/scarlet0725/prism-api/cmd"
	"github.com/scarlet0725/prism-api/model"
	"github.com/scarlet0725/prism-api/infrastructure/repository"
)

type Artist interface {
	CreateArtist(artist *model.Artist) (*model.Artist, error)
	//UpdateArtist(artist *model.Artist) (*model.Artist, error)
	//GetArtistByName(name string) ([]*model.Artist, error)
	GetArtistByID(id string) (*model.Artist, error)
}

type artistUsecase struct {
	db repository.DB
}

func NewArtistUsecase(db repository.DB) Artist {
	return &artistUsecase{
		db: db,
	}
}

func (a *artistUsecase) CreateArtist(artist *model.Artist) (*model.Artist, error) {
	id := cmd.MakeRamdomID(artistIDLength)

	artist.ArtistID = id

	result, err := a.db.CreateArtist(artist)
	if err != nil {
		return nil, &model.AppError{
			Msg:  "Failed to create artist",
			Code: 500,
		}
	}
	return result, nil
}

func (a *artistUsecase) GetArtistByID(id string) (*model.Artist, error) {
	result, err := a.db.GetArtistByID(id)
	if err != nil {
		return nil, &model.AppError{
			Msg:  "Failed to get artist",
			Code: 500,
		}
	}
	return result, nil
}
