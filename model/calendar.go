package model

import "time"

type UserCalendar struct {
	ID                 int    `json:"-"`
	UserID             uint   `json:"-"`
	CalendarID         string `json:"calendar_id"`
	ExternalCalendarID string `json:"external_id"`
}

type ArtistCalendar struct {
	ID                 int     `json:"-"`
	ExternalCalendarID string  `json:"external_id"`
	SorceType          string  `json:"source_type"`
	Artist             *Artist `json:"artist"`
	ArtistID           *uint   `json:"-"`
}

type CalendarEvent struct {
	ExternalEventID    string `json:"external_event_id"`
	ExternalCalendarID string `json:"external_calender_id"`
	Public             bool   `json:"public"`
	Event              *Event `json:"event"`
}

type ExternalCalendar struct {
	ID          int        `json:"-"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	CalendarID  string     `json:"calendar_id"`
	Type        string     `json:"type"`
	UserID      int        `json:"-"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `gorm:"index" json:"-"`
}

type GoogleCalenderConfig struct {
	ExternalCalendar
	GoogleOAuthToken
}
