package repository

import "github.com/scarlet0725/prism-api/model"

type Venue interface {
	CreateVenue(*model.Venue) (*model.Venue, error)
	//GetVenueByName(name string) (*model.Venue, error)
	GetVenueByID(id string) (*model.Venue, error)
	//GetVenuesByNames(names []string) ([]*model.Venue, error)
	//UpdateVenue(*model.Venue) (*model.Venue, error)
}
