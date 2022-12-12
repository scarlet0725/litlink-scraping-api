package repository

import (
	"errors"

	"github.com/scarlet0725/prism-api/model"
)

var (
	ErrInvalidEvent             = errors.New("invalid event")
	ErrCalendarPermissionDenied = errors.New("calendar permission denied")
)

type ExternalCalendar interface {
	CreateEvent(*model.CalenderEvent) (*model.Event, error)
	UpdateEvent(*model.CalenderEvent) (*model.Event, error)
	DeleteEvent(*model.CalenderEvent) error
	GetEvent(calenderID string, eventID string) (*model.Event, error)
	CreateCalendar(*model.ExternalCalendar) (*model.ExternalCalendar, error)
}
