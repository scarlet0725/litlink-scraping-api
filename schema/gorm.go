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

type Venue struct {
	gorm.Model
	model.Venue
}

type RyzmEvent struct {
	model.RyzmEvent
}

type UnStructuredEventInformation struct {
	model.UnStructuredEventInformation
}

type Role struct {
	gorm.Model
	model.Role
}
