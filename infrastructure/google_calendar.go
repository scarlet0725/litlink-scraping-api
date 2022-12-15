package infrastructure

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/scarlet0725/prism-api/infrastructure/repository"
	"github.com/scarlet0725/prism-api/model"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

const (
	timeZone   = "Asia/Tokyo"
	dateFormat = "2006-01-02"
)

type googleCalendar struct {
	client   *http.Client
	srv      *calendar.Service
	location *time.Location
}

func NewGoogleCalenderClient(client *http.Client) repository.ExternalCalendar {
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		loc = time.UTC
	}

	srv, err := calendar.NewService(context.TODO(), option.WithHTTPClient(client))
	if err != nil {
		return nil
	}

	return &googleCalendar{
		client:   client,
		location: loc,
		srv:      srv,
	}
}

func (g *googleCalendar) RegisterEventCalendar(eventCalendar *model.ArtistCalendar) (*model.ArtistCalendar, error) {
	return nil, nil
}

func (g *googleCalendar) buildGoogleEventStruct(event *model.Event) *calendar.Event {
	schedule := &calendar.Event{
		Id:          strings.ToLower(event.EventID),
		Summary:     event.Name,
		Description: event.Description,
	}

	if event.StartTime != nil && event.EndTime != nil {
		schedule.Start = &calendar.EventDateTime{
			DateTime: event.StartTime.Format(time.RFC3339),
		}
		schedule.End = &calendar.EventDateTime{
			DateTime: event.EndTime.Format(time.RFC3339),
		}
	} else {
		schedule.Start = &calendar.EventDateTime{
			Date:     event.Date.Format(dateFormat),
			TimeZone: timeZone,
		}
		schedule.End = &calendar.EventDateTime{
			Date:     event.Date.Format(dateFormat),
			TimeZone: timeZone,
		}
	}

	return schedule
}

func (g *googleCalendar) parseGoogleEvent(schedule *calendar.Event) *model.Event {
	event := &model.Event{}

	event.EventID = strings.ToUpper(schedule.Id)
	event.Name = schedule.Summary
	event.Description = schedule.Description

	switch schedule.Start.DateTime == "" {
	case true:
		date, _ := time.ParseInLocation(dateFormat, schedule.Start.Date, g.location)
		event.Date = &date
	case false:
		startTime, _ := time.ParseInLocation(time.RFC3339, schedule.Start.DateTime, g.location)
		endTime, _ := time.ParseInLocation(time.RFC3339, schedule.End.DateTime, g.location)
		event.StartTime = &startTime
		event.EndTime = &endTime
	}

	return event
}

func (g *googleCalendar) CreateEvent(event *model.CalendarEvent) (*model.Event, error) {
	//Nilチェック
	if event.Event == nil {
		return nil, repository.ErrInvalidEvent
	}

	cal, err := g.srv.Events.List(event.ExternalCalendarID).Do()

	if err != nil {
		return nil, repository.ErrCalendarPermissionDenied
	}

	if !(cal.AccessRole == "owner") || !(cal.AccessRole != "writer") || err != nil {
		return nil, repository.ErrCalendarPermissionDenied
	}

	schedule := g.buildGoogleEventStruct(event.Event)

	switch event.Public {
	case true:
		schedule.Visibility = "public"
	case false:
		schedule.Visibility = "private"
	}

	_, err = g.srv.Events.Insert(event.ExternalCalendarID, schedule).Do()

	if err != nil {
		return nil, err
	}

	return event.Event, nil
}

func (g *googleCalendar) GetEvent(calenderID string, eventID string) (*model.Event, error) {
	schedule, err := g.srv.Events.Get(calenderID, eventID).Do()

	if err != nil {
		return nil, err
	}
	event := g.parseGoogleEvent(schedule)

	return event, nil

}

func (g *googleCalendar) UpdateEvent(event *model.CalendarEvent) (*model.Event, error) {
	if event.Event == nil {
		return nil, repository.ErrInvalidEvent
	}
	cal, err := g.srv.Events.List(event.ExternalCalendarID).Do()

	if !(cal.AccessRole == "owner") || !(cal.AccessRole != "writer") || err != nil {
		return nil, repository.ErrCalendarPermissionDenied
	}

	schedule := g.buildGoogleEventStruct(event.Event)

	switch event.Public {
	case true:
		schedule.Visibility = "public"
	case false:
		schedule.Visibility = "private"
	}

	result, err := g.srv.Events.Update(event.ExternalCalendarID, schedule.Id, schedule).Do()

	e := g.parseGoogleEvent(result)

	if err != nil {
		return nil, err
	}

	return e, nil
}

func (g *googleCalendar) DeleteEvent(event *model.CalendarEvent) error {
	cal, err := g.srv.Events.List(event.ExternalCalendarID).Do()

	if !(cal.AccessRole == "owner") || !(cal.AccessRole != "writer") || err != nil {
		return errors.New("you don't have permission to delete event")
	}

	err = g.srv.Events.Delete(event.ExternalCalendarID, event.Event.EventID).Do()

	if err != nil {
		return err
	}

	return nil
}

func (g *googleCalendar) CreateCalendar(calender *model.ExternalCalendar) (*model.ExternalCalendar, error) {
	cal, err := g.srv.Calendars.Insert(&calendar.Calendar{
		Summary:     calender.Name,
		Description: calender.Description,
		TimeZone:    timeZone,
	}).Do()

	if err != nil {
		return nil, err
	}

	result := &model.ExternalCalendar{
		Name:        cal.Summary,
		Description: cal.Description,
		CalendarID:  cal.Id,
		Type:        "google",
		UserID:      calender.UserID,
	}

	return result, nil
}
