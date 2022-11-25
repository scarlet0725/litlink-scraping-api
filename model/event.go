package model

import (
	"time"
)

type Event struct {
	ID          uint       `json:"-" gorm:"primary_key;unique;not null;auto_increment"`
	EventID     string     `json:"event_id" gorm:"unique;not null"`
	UUID        string     `json:"uuid"`
	Name        string     `json:"name" gorm:"not null"`
	Date        *time.Time `json:"date"`
	OpenTime    *time.Time `json:"open_time"`
	StartTime   *time.Time `json:"start_time"`
	EndTime     *time.Time `json:"end_time"`
	Description string     `json:"description"`
	VenueName   string     `json:"venue_name"`
	Venue       *Venue     `json:"venue"`
	VenueID     uint       `json:"-"`
	Url         string     `json:"url"`
	TicketURL   string     `json:"ticket_url"`
	ArtistName  string     `json:"artist_name"`
	Artists     []*Artist  `json:"artists" gorm:"many2many:events_artists"`
}
