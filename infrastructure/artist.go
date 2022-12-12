package infrastructure

import (
	"github.com/scarlet0725/prism-api/infrastructure/repository"
	"github.com/scarlet0725/prism-api/model"
	"github.com/scarlet0725/prism-api/schema"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type artist struct {
	db *gorm.DB
}

func NewArtistRepository(db *gorm.DB) repository.Artist {
	return &artist{
		db: db,
	}
}

func (a *artist) CreateArtist(artist *model.Artist) (*model.Artist, error) {
	var schema schema.Artist
	schema.Artist = *artist
	err := a.db.Create(&artist).Error
	if err != nil {
		return nil, err
	}
	return artist, nil
}

func (a *artist) GetArtistByName(name string) (*model.Artist, error) {
	var artist model.Artist
	err := a.db.Preload(clause.Associations).Where("name = ?", name).First(&artist).Error
	if err != nil {
		return nil, err
	}
	return &artist, nil
}

func (a *artist) GetArtistByID(id string) (*model.Artist, error) {
	var artist model.Artist
	err := a.db.Where("artist_id = ?", id).First(&artist).Error
	if err != nil {
		return nil, err
	}
	return &artist, nil
}

func (a *artist) GetArtistsByIDs(ids []string) ([]*model.Artist, error) {
	var artists []*model.Artist
	err := a.db.Where("artist_id IN ?", ids).Find(&artists).Error
	if err != nil {
		return nil, err
	}
	return artists, nil
}

func (a *artist) GetEventsByID(ID string) (*model.Event, error) {
	var event *model.Event
	err := a.db.Preload(clause.Associations).Where("event_id = ?", ID).First(&event).Error
	if err != nil {
		return nil, err
	}
	return event, nil
}
