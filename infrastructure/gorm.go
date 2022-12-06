package infrastructure

import (
	"errors"

	"github.com/scarlet0725/prism-api/infrastructure/repository"
	"github.com/scarlet0725/prism-api/model"
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
	err := g.db.Preload(clause.Associations).Where("user_id = ?", id).First(&user).Error
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
	err := g.db.Preload(clause.Associations).Where("name = ?", name).First(&artist).Error
	if err != nil {
		return nil, err
	}
	return &artist, nil
}

func (g *gormDB) GetUserByAPIKey(apiKey string) (*model.User, error) {
	var user model.User
	err := g.db.Preload(clause.Associations).Where("api_key = ?", apiKey).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (g *gormDB) UpdateUser(user *model.User) (*model.User, error) {
	var u schema.User
	u.User = *user
	err := g.db.Save(&u).Error
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
	err := g.db.Preload(clause.Associations).Where("event_id = ?", id).First(&event).Error
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
	if event.ID == 0 {
		return nil, errors.New("invalid id")
	}
	var e schema.Event
	e.Event = *event
	err := g.db.Model(&e).Where("id = ?", e.Event.ID).Updates(&e).Error
	if err != nil {
		return nil, err
	}
	return &e.Event, nil
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
	err := g.db.Preload("Artist").Preload("RyzmEvent").Where("event_id = ?", ID).First(&event).Error
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (g *gormDB) GetRyzmEventsByUUDIDs(ids []string) ([]*model.RyzmEvent, error) {
	var ryzmEvents []*model.RyzmEvent
	if err := g.db.Where("uuid IN ?", ids).Find(&ryzmEvents).Error; err != nil {
		return nil, err
	}
	return ryzmEvents, nil
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

func (g *gormDB) GetVenueByName(name string) (*model.Venue, error) {
	var venue *model.Venue
	var v schema.Venue
	v.Venue = *venue
	err := g.db.Where("name = ?", name).Take(&v).Error
	return venue, err
}

func (g *gormDB) GetVenuesByNames(names []string) ([]*model.Venue, error) {
	var venues []*model.Venue
	err := g.db.Where("name IN ?", names).Find(&venues).Error
	if err != nil {
		return nil, err
	}
	return venues, nil
}

func (g *gormDB) MergeEvents(base *model.Event, target *model.Event) (*model.Event, error) {
	tx := g.db.Begin()
	var e schema.Event
	e.Event = *base
	e.Model.ID = base.ID

	if tx.Updates(&e).Error != nil {
		tx.Rollback()
		return nil, &model.AppError{
			Code: 404,
			Msg:  "Base event not found",
		}
	}

	if err := tx.Model(target).Association("Artists").Clear(); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Model(target).Association("Users").Clear(); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Model(target).Association("RelatedRyzmEvents").Clear(); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Model(target).Association("Venue").Clear(); err != nil {
		tx.Rollback()
		return nil, err
	}

	if tx.Where("id = ?", target.ID).Delete(&schema.Event{Event: *target}).Error != nil {
		tx.Rollback()
		return nil, errors.New("failed to delete target event")
	}

	return base, tx.Commit().Error
}

func (g *gormDB) GetGoogleOAuthStateByState(state string) (*model.GoogleOAuthState, error) {
	var s *model.GoogleOAuthState
	err := g.db.Where("state = ?", state).First(&s).Error
	return s, err
}

func (g *gormDB) SaveGoogleOAuthState(s *model.GoogleOAuthState) (*model.GoogleOAuthState, error) {
	err := g.db.Clauses(
		clause.OnConflict{
			UpdateAll: true,
		},
	).Create(s).Error

	return s, err
}

func (g *gormDB) SaveGoogleOAuthToken(s *model.GoogleOAuthToken) (*model.GoogleOAuthToken, error) {
	err := g.db.Clauses(
		clause.OnConflict{
			UpdateAll: true,
		},
	).Create(s).Error

	return s, err
}
