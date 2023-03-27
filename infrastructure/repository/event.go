package repository

import (
	"context"

	"github.com/scarlet0725/prism-api/model"
)

type Event interface {
	CreateEvent(context.Context, *model.Event) (*model.Event, error)
	CreateEvents(context.Context, []*model.Event) ([]*model.Event, error)
	UpdateEvent(context.Context, *model.Event) (*model.Event, error)
	DeleteEvent(context.Context, *model.Event) error
	GetEventsByArtistID(context.Context, string) ([]*model.Event, error)
	GetEventByID(context.Context, string) (*model.Event, error)
	GetRyzmEventsByUUDIDs(context.Context, []string) ([]*model.RyzmEvent, error)
	MergeEvents(ctx context.Context, base *model.Event, target *model.Event) (*model.Event, error)
	SearchEvents(context.Context, *model.EventSearchQuery) ([]*model.EventSearchResult, error)
}
