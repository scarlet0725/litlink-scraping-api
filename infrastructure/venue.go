package infrastructure

import (
	"context"

	"github.com/scarlet0725/prism-api/ent"
	"github.com/scarlet0725/prism-api/ent/venue"
	"github.com/scarlet0725/prism-api/infrastructure/repository"
	"github.com/scarlet0725/prism-api/infrastructure/translator"
	"github.com/scarlet0725/prism-api/model"
)

type venueRepository struct {
	db *ent.Client
}

func NewVenueRepository(db *ent.Client) repository.Venue {
	return &venueRepository{
		db: db,
	}
}

func (v *venueRepository) CreateVenue(venue *model.Venue) (*model.Venue, error) {
	result, err := v.db.Venue.Create().
		SetVenueID(venue.VenueID).
		SetName(venue.Name).
		SetDescription(venue.Description).
		SetWebSite(venue.WebSite).
		SetIsOpen(venue.IsOpen).
		SetPostcode(venue.Postcode).
		SetPrefecture(venue.Prefecture).
		SetCity(venue.City).
		SetStreet(venue.Street).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	venue = translator.VenueFromEnt(result)

	return venue, nil
}

func (v *venueRepository) GetVenueByID(id string) (*model.Venue, error) {
	result, err := v.db.Venue.Query().Where(venue.VenueID(id)).Only(context.Background())

	if err != nil {
		return nil, err
	}

	venue := translator.VenueFromEnt(result)

	return venue, nil
}
