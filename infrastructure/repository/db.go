package repository

import (
	"github.com/scarlet0725/prism-api/model"
)

type DB interface {
	GetUser(id string) (*model.User, error)
	CreateUser(*model.User) (*model.User, error)
	GetUserByAPIKey(apiKey string) (*model.User, error)
	UpdateUser(*model.User) (*model.User, error)
	DeleteUser(*model.User) error

	CreateEvents([]*model.Event) ([]*model.Event, error)
	CreateEvent(*model.Event) (*model.Event, error)
	UpdateEvent(*model.Event) (*model.Event, error)
	DeleteEvent(*model.Event) error
	GetEventsByArtistID(artistID string) ([]*model.Event, error)
	GetEventByID(ID string) (*model.Event, error)
	GetRyzmEventsByUUDIDs(IDs []string) ([]*model.RyzmEvent, error)
	MergeEvents(base *model.Event, target *model.Event) (*model.Event, error)

	CreateArtist(*model.Artist) (*model.Artist, error)
	GetArtistByName(name string) (*model.Artist, error)
	GetArtistByID(id string) (*model.Artist, error)
	GetArtistsByIDs(ids []string) ([]*model.Artist, error)

	CreateVenue(*model.Venue) (*model.Venue, error)
	GetVenueByName(name string) (*model.Venue, error)
	GetVenueByID(id string) (*model.Venue, error)
	GetVenuesByNames(names []string) ([]*model.Venue, error)
	UpdateVenue(*model.Venue) (*model.Venue, error)

	SaveGoogleOAuthState(*model.GoogleOAuthState) (*model.GoogleOAuthState, error)
	GetGoogleOAuthStateByState(state string) (*model.GoogleOAuthState, error)
	SaveGoogleOAuthToken(*model.GoogleOAuthToken) (*model.GoogleOAuthToken, error)
}
