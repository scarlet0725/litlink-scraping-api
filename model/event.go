package model

import (
	"time"
)

type Event struct {
	ID                      uint                          `json:"-" gorm:"primary_key;unique;not null;auto_increment"`
	EventID                 string                        `json:"event_id" gorm:"unique;not null"`
	Name                    string                        `json:"name" gorm:"not null"`
	Date                    *time.Time                    `json:"date"`
	OpenTime                *time.Time                    `json:"open_time"`
	StartTime               *time.Time                    `json:"start_time"`
	EndTime                 *time.Time                    `json:"end_time"`
	Description             string                        `json:"description"`
	Url                     string                        `json:"url"`
	TicketURL               string                        `json:"ticket_url"`
	Artists                 []*Artist                     `json:"artists" gorm:"many2many:events_artists"`
	Venue                   *Venue                        `json:"venue"`
	VenueID                 *uint                         `json:"-"`
	RelatedRyzmEvents       []*RyzmEvent                  `json:"-" gorm:"foreignkey:EventID"`
	Users                   []*User                       `json:"-" gorm:"many2many:user_events"`
	UserCreated             *User                         `json:"-" gorm:"foreignkey:UserCreatedID"`
	UserCreatedID           *uint                         `json:"-"`
	UnStructuredInformation *UnStructuredEventInformation `json:"-" gorm:"foreignkey:EventID"`
}

type RyzmEvent struct {
	ID      uint `gorm:"primary_key;unique;not null;auto_increment"`
	EventID uint
	Event   *Event
	UUID    string
}

type UnStructuredEventInformation struct {
	ID         uint `gorm:"primary_key;unique;not null;auto_increment"`
	EventID    uint
	RyzmUUID   string
	VenueName  string
	ArtistName string
	Price      string
	Event      *Event
}

type EventSearchQuery struct {
	ArtistName string    `form:"artist_name"`
	VenueName  string    `form:"venue_name"`
	EventName  string    `form:"event_name"`
	ArtistID   string    `form:"artist_id"`
	VenueID    string    `form:"venue_id"`
	EventID    string    `form:"event_id"`
	DateAfter  time.Time `form:"date_after" time_format:"2006-01-02"  validate:"datetime=2006-01-02"`
	DateBefore time.Time `form:"date_before" time_format:"2006-01-02"  validate:"datetime=2006-01-02"`
}

type EventSearchResult struct {
	Event
	EventArtist
	Artist
	Venue
}

type EventArtist struct {
	EventID   int `gorm:"primaryKey"`
	ArtistID  int `gorm:"primaryKey"`
	CreatedAt time.Time
}
