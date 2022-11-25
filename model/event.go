package model

import (
	"time"
)

type Event struct {
	ID                uint         `json:"-" gorm:"primary_key;unique;not null;auto_increment"`
	EventID           string       `json:"event_id" gorm:"unique;not null"`
	UUID              string       `json:"-"`
	Name              string       `json:"name" gorm:"not null"`
	Date              *time.Time   `json:"date"`
	OpenTime          *time.Time   `json:"open_time"`
	StartTime         *time.Time   `json:"start_time"`
	EndTime           *time.Time   `json:"end_time"`
	Description       string       `json:"description"`
	Url               string       `json:"url"`
	TicketURL         string       `json:"ticket_url"`
	ArtistName        string       `json:"artist_name"`
	Artists           []*Artist    `json:"artists" gorm:"many2many:events_artists"`
	VenueName         string       `json:"venue_name"`
	Venue             *Venue       `json:"venue"`
	VenueID           *uint        `json:"-"`
	RelatedRyzmEvents []*RyzmEvent `json:"-" gorm:"many2many:related_ryzm_events"`
	Users             []*User      `json:"-" gorm:"many2many:user_events"`
	UserCreated       *User        `json:"-" gorm:"foreignkey:UserCreatedID"`
	UserCreatedID     *uint        `json:"-"`
}

type RyzmEvent struct {
	ID   uint `gorm:"primary_key;unique;not null;auto_increment"`
	UUID string
}
