package usecase

import (
	"github.com/scarlet0725/prism-api/gateway"
	"github.com/scarlet0725/prism-api/model"
)

type EventApplication interface {
	CreateEvent(*model.Event) error
	GetEvent(string) (*model.Event, error)
	GetEventsByName(string) ([]*model.Event, error)
}

type eventApplication struct {
	s gateway.Client
}

func (a *eventApplication) CreateEvent(e *model.Event) error {
	return nil
}

func (a *eventApplication) GetEventsByName(name string) ([]*model.Event, error) {
	return nil, nil
}
