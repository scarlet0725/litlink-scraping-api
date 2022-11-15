package model

import (
	"time"
)

type Event struct {
	ID          string    `json:"id"`
	UUID        string    `json:"uuid"`
	Name        string    `json:"name"`
	Date        time.Time `json:"date"`
	OpenTime    time.Time `json:"open_time"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Description string    `json:"description"`
	Venue       string    `json:"venue"`
	Url         string    `json:"url"`
	TicketURL   string    `json:"ticket_url"`
	Artist      string    `json:"artist"`
}
