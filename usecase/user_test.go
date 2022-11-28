package usecase

import (
	"github.com/scarlet0725/prism-api/model"
	"github.com/scarlet0725/prism-api/repository"
)

type userTestDB struct {
	user   map[string]*model.User
	artist map[string]*model.Artist
	event  map[string]*model.Event
}

func NewUserTestDB() repository.DB {
	return &userTestDB{
		user:   map[string]*model.User{},
		artist: map[string]*model.Artist{},
		event:  map[string]*model.Event{},
	}
}

func (u *userTestDB) GetUser(id string) (*model.User, error) {
	return u.user[id], nil
}

func (u *userTestDB) CreateUser(user *model.User) (*model.User, error) {
	u.user[user.UserID] = user
	return user, nil
}

func (u *userTestDB) GetArtistByName(name string) (*model.Artist, error) {
	return nil, nil
}

func (u *userTestDB) CreateArtist(artist *model.Artist) (*model.Artist, error) {
	u.artist[artist.ArtistID] = artist
	return artist, nil
}

func (u *userTestDB) CreateEvent(event *model.Event) (*model.Event, error) {
	u.event[event.EventID] = event
	return event, nil
}

func (u *userTestDB) DeleteEvent(event *model.Event) error {
	return nil
}

func (u *userTestDB) GetEventsByArtistID(artistID string) ([]*model.Event, error) {
	return nil, nil
}

func (u *userTestDB) GetArtistByID(ID string) (*model.Artist, error) {
	return nil, nil
}

func (u *userTestDB) GetArtistsByIDs(ids []string) ([]*model.Artist, error) {
	return nil, nil
}

func (u *userTestDB) GetEventByID(ID string) (*model.Event, error) {
	return nil, nil
}

func (u *userTestDB) GetUserByAPIKey(apiKey string) (*model.User, error) {
	return nil, nil
}

func (u *userTestDB) UpdateUser(user *model.User) (*model.User, error) {
	return nil, nil
}

func (u *userTestDB) UpdateEvent(event *model.Event) (*model.Event, error) {
	return nil, nil
}
