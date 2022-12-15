package usecase

import (
	"strings"

	"github.com/scarlet0725/prism-api/infrastructure/repository"
	"github.com/scarlet0725/prism-api/model"
)

const (
	defaultCalendarName = "prism calendar"
	defaultDescription  = "prismによる自動同期カレンダー"
)

type Calendar interface {
	CreateEvent(*model.Event) (*model.Event, error)
	CreateCalender(*model.ExternalCalendar) (*model.ExternalCalendar, error)
}

type calendarUsecase struct {
	cal repository.ExternalCalendar
	db  repository.DB
}

func NewCalendarApplication(cal repository.ExternalCalendar, db repository.DB) Calendar {
	return &calendarUsecase{
		cal: cal,
		db:  db,
	}
}

func (a *calendarUsecase) CreateEvent(event *model.Event) (*model.Event, error) {
	//TODO: 実装する
	return nil, nil
}

func (a *calendarUsecase) CreateCalender(cal *model.ExternalCalendar) (*model.ExternalCalendar, error) {

	if strings.EqualFold(cal.Name, "default") || cal.Name == "" {
		cal.Name = defaultCalendarName
	}

	if cal.Description == "" {
		cal.Description = defaultDescription
	}

	result, err := a.cal.CreateCalendar(cal)

	return result, err
}
