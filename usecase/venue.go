package usecase

import (
	"github.com/scarlet0725/prism-api/model"
	"github.com/scarlet0725/prism-api/repository"
)

type Venue interface {
	CreateVenue(*model.Venue) (*model.Venue, error)
	UpdateVenue(*model.Venue) (*model.Venue, error)
	//GetVenueByName(string) (*model.Venue, error)
	GetVenueByID(string) (*model.Venue, error)
}

type venueUsecase struct {
	db repository.DB
}

func NewVenueApplication(db repository.DB) Venue {
	return &venueUsecase{
		db: db,
	}
}

func (a *venueUsecase) CreateVenue(venue *model.Venue) (*model.Venue, error) {
	venue, err := a.db.CreateVenue(venue)

	if err != nil {
		return nil, err
	}

	return venue, nil
}

func (a *venueUsecase) UpdateVenue(venue *model.Venue) (*model.Venue, error) {
	venue, err := a.db.UpdateVenue(venue)

	if err != nil {
		return nil, err
	}

	return venue, nil
}

func (a *venueUsecase) GetVenueByID(id string) (*model.Venue, error) {
	venue, err := a.db.GetVenueByID(id)

	if err != nil {
		return nil, err
	}

	return venue, nil
}
