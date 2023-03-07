package repository

import (
	"github.com/scarlet0725/prism-api/model"
)

type Event interface {
	//CreateEvents([]*model.Event) ([]*model.Event, error)
	CreateEvent(*model.Event) (*model.Event, error)
	UpdateEvent(*model.Event) (*model.Event, error)
	DeleteEvent(*model.Event) error
	GetEventsByArtistID(artistID string) ([]*model.Event, error)
	GetEventByID(ID string) (*model.Event, error)
	GetRyzmEventsByUUDIDs(IDs []string) ([]*model.RyzmEvent, error)
	MergeEvents(base *model.Event, target *model.Event) (*model.Event, error)
	SearchEvents(*model.EventSearchQuery) ([]*model.EventSearchResult, error)
}
