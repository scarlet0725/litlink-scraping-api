package repository

import (
	"context"

	"github.com/scarlet0725/prism-api/model"
)

type Venue interface {
	CreateVenue(context.Context, *model.Venue) (*model.Venue, error)
	//GetVenueByName(name string) (*model.Venue, error)
	GetVenueByID(context.Context, string) (*model.Venue, error)
	GetVenuesByNames(ctx context.Context, names []string) ([]*model.Venue, error)
	//UpdateVenue(*model.Venue) (*model.Venue, error)
}
