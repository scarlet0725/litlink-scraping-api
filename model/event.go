package model

import (
	"time"
)

type Event struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Date      time.Time `json:"date"`
	OpenTime  time.Time `json:"open_time"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Location  string    `json:"location"`
	Url       string    `json:"url"`
}
