package schema

import (
	"github.com/scarlet0725/prism-api/model"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	model.User
}

type Artist struct {
	gorm.Model
	model.Artist
}

type Event struct {
	gorm.Model
	model.Event
}

type EventsArtists struct {
	EventID  int `gorm:"primaryKey"`
	ArtistID int `gorm:"primaryKey"`
}
