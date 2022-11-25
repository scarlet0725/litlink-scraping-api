package infrastructure

import (
	"errors"

	"github.com/scarlet0725/prism-api/model"
	"github.com/scarlet0725/prism-api/repository"
	"github.com/scarlet0725/prism-api/schema"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type gormDB struct {
	db *gorm.DB
}

func NewGORMClient(db *gorm.DB) repository.DB {
	return &gormDB{db}
}

func (g *gormDB) GetUser(id string) (*model.User, error) {
	var user model.User
	err := g.db.Where("user_id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (g *gormDB) CreateUser(user *model.User) (*model.User, error) {
	var u schema.User
	u.User = *user

	result := g.db.Create(&u)

	err := result.Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (g *gormDB) GetArtistByName(name string) (*model.Artist, error) {
	var artist model.Artist
	err := g.db.Where("name = ?", name).First(&artist).Error
	if err != nil {
		return nil, err
	}
	return &artist, nil
}

func (g *gormDB) GetUserByAPIKey(apiKey string) (*model.User, error) {
	var user model.User
	err := g.db.Where("api_key = ?", apiKey).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (g *gormDB) UpdateUser(user *model.User) (*model.User, error) {
	var u schema.User
	u.User = *user
	err := g.db.Save(u).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (g *gormDB) DeleteEvent(model *model.Event) error {
	if model.ID == 0 {
		return errors.New("invalid id")
	}

	return g.db.Select(clause.Associations).Delete(model).Error
}

func (g *gormDB) GetEventByID(id string) (*model.Event, error) {
	var event model.Event
	err := g.db.Where("event_id = ?", id).First(&event).Error
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (g *gormDB) CreateEvent(event *model.Event) (*model.Event, error) {
	var e schema.Event
	e.Event = *event
	err := g.db.Create(&e).Error
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (g *gormDB) CreateEvents(events []*model.Event) ([]*model.Event, error) {
	var es []*schema.Event
	for _, event := range events {
		es = append(es, &schema.Event{Event: *event})
	}
	err := g.db.Create(&es).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (g *gormDB) UpdateEvent(event *model.Event) (*model.Event, error) {
	var e schema.Event
	e.Event = *event
	err := g.db.Save(e).Error
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (g *gormDB) GetEventsByArtistID(artistID string) ([]*model.Event, error) {
	var events []*model.Event
	err := g.db.Where("artist_id = ?", artistID).Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (g *gormDB) CreateArtist(artist *model.Artist) (*model.Artist, error) {
	var a schema.Artist
	a.Artist = *artist
	err := g.db.Create(&a).Error
	if err != nil {
		return nil, err
	}
	return artist, nil
}

func (g *gormDB) GetArtistByID(id string) (*model.Artist, error) {
	var artist model.Artist
	err := g.db.Where("artist_id = ?", id).First(&artist).Error
	if err != nil {
		return nil, err
	}
	return &artist, nil
}

func (g *gormDB) GetArtistsByIDs(ids []string) ([]*model.Artist, error) {
	var artists []*model.Artist
	err := g.db.Where("artist_id IN ?", ids).Find(&artists).Error
	if err != nil {
		return nil, err
	}
	return artists, nil
}

func (g *gormDB) GetEventsByID(ID string) (*model.Event, error) {
	var event *model.Event
	err := g.db.Preload("Artist").Where("event_id = ?", ID).First(&event).Error
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (g *gormDB) GetEventsByUUIDs(IDs []string) ([]*model.Event, error) {
	var events []*model.Event
	err := g.db.Where("uuid IN ?", IDs).Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (g *gormDB) CreateVenue(v *model.Venue) (*model.Venue, error) {
	var s schema.Venue
	s.Venue = *v
	err := g.db.Create(&s).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (g *gormDB) GetVenueByID(id string) (*model.Venue, error) {
	var venue *model.Venue
	err := g.db.Where("venue_id = ?", id).First(&venue).Error
	if err != nil {
		return nil, err
	}
	return venue, nil
}

func (g *gormDB) DeleteVenue(v *model.Venue) error {
	if v.ID == 0 {
		return errors.New("invalid id")
	}

	return g.db.Select(clause.Associations).Delete(v).Error
}

func (g *gormDB) UpdateVenue(v *model.Venue) (*model.Venue, error) {
	var s *schema.Venue
	s.Venue = *v
	err := g.db.Save(s).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (g *gormDB) DeleteUser(user *model.User) error {
	return g.db.Delete(user).Error
}
