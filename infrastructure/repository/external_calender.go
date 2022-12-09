package repository

import (
	"github.com/scarlet0725/prism-api/model"
)

type ExternalCalendar interface {
	CreateEvent(*model.CalenderEvent) (*model.Event, error)
	GetEventByExternalCalenderID(CalenderID string, EventID string) error
}
