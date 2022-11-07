package usecase

import (
	"github.com/scarlet0725/litlink-scraping-api/controller"
	"github.com/scarlet0725/litlink-scraping-api/model"
	"github.com/scarlet0725/litlink-scraping-api/scraping"
)

type EventApplication interface {
	CreateEvent(*model.Event) error
	GetEvent(string) (*model.Event, error)
	GetEvents() ([]*model.Event, error)
}

type eventApplication struct {
	s scraping.Client
	c controller.Controller
}

func (a *eventApplication) CreateEvent(e *model.Event) error {
	return nil
}

func (a *eventApplication) GetEvent(id string) (*model.Event, error) {
	return nil, nil
}
