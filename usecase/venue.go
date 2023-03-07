package usecase

import (
	"context"

	"github.com/scarlet0725/prism-api/framework"
	"github.com/scarlet0725/prism-api/infrastructure/repository"
	"github.com/scarlet0725/prism-api/model"
)

type Venue interface {
	CreateVenue(context.Context, *model.Venue) (*model.Venue, error)
	//UpdateVenue(*model.Venue) (*model.Venue, error)
	GetVenueByID(context.Context, string) (*model.Venue, error)
}

type venueUsecase struct {
	db     repository.Venue
	ramdom framework.RandomID
}

func NewVenueUsecase(db repository.Venue, r framework.RandomID) Venue {
	return &venueUsecase{
		db:     db,
		ramdom: r,
	}
}

func (a *venueUsecase) CreateVenue(ctx context.Context, venue *model.Venue) (*model.Venue, error) {
	id := a.ramdom.Generate(venueIDLength)

	venue.VenueID = id
	venue.ID = 0
	venue.IsOpen = true

	venue, err := a.db.CreateVenue(ctx, venue)

	if err != nil {
		return nil, err
	}

	return venue, nil
}

func (a *venueUsecase) GetVenueByID(ctx context.Context, id string) (*model.Venue, error) {
	venue, err := a.db.GetVenueByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return venue, nil
}
