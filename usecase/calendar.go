package usecase

import "github.com/scarlet0725/prism-api/model"

type Calendar interface {
	CreateEvent(*model.Event) (*model.Event, error)
}

type calendarUsecase struct {
}

func NewCalendarApplication() Calendar {
	return &calendarUsecase{}
}

func (a *calendarUsecase) CreateEvent(event *model.Event) (*model.Event, error) {
	return nil, nil
}