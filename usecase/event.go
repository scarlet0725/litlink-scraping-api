package usecase

import (
	"context"
	"time"

	"github.com/scarlet0725/prism-api/controller"
	"github.com/scarlet0725/prism-api/framework"
	"github.com/scarlet0725/prism-api/infrastructure/repository"
	"github.com/scarlet0725/prism-api/model"
	"github.com/scarlet0725/prism-api/parser"
	"github.com/scarlet0725/prism-api/selializer"
)

type Event interface {
	CreateEvent(context.Context, *model.CreateEvent) (*model.Event, error)
	DeleteEvent(context.Context, *model.Event) error
	UpdateEvent(context.Context, *model.UpdateEvent) (*model.Event, error)
	//GetEvent(string) (*model.Event, error)
	GetEventByID(context.Context, string) (*model.Event, error)
	MergeEvents(context.Context, *model.MergeEvent) (*model.Event, error)
	SearchEvents(context.Context, *model.EventSearchQuery) ([]*model.Event, error)
}

type eventUsecase struct {
	fetch      controller.FetchController
	parser     parser.DocParser
	selializer selializer.ResponseSerializer
	json       parser.JsonParser
	random     framework.RandomID
	event      repository.Event
	artist     repository.Artist
	venue      repository.Venue
}

func NewEventUsecase(
	event repository.Event,
	artist repository.Artist,
	venue repository.Venue,
	fetch controller.FetchController,
	parser parser.DocParser,
	selializer selializer.ResponseSerializer,
	json parser.JsonParser,
	r framework.RandomID,
) Event {
	return &eventUsecase{
		event:      event,
		artist:     artist,
		fetch:      fetch,
		parser:     parser,
		selializer: selializer,
		json:       json,
		random:     r,
	}
}

func (e *eventUsecase) CreateEvent(ctx context.Context, event *model.CreateEvent) (*model.Event, error) {
	id := e.random.Generate(eventIDLength)

	artists, _ := e.artist.GetArtistsByIDs(ctx, event.ArtistIDs)
	venue, _ := e.venue.GetVenueByID(ctx, event.VenueID)
	jst, _ := time.LoadLocation(Locale)

	date, err := time.ParseInLocation("2006-01-02", event.Date, jst)

	if err != nil {
		return nil, err
	}

	newEvent := &model.Event{
		EventID:     id,
		Name:        event.Name,
		Date:        &date,
		Description: event.Description,
		OpenTime:    event.OpenTime,
		StartTime:   event.StartTime,
		EndTime:     event.EndTime,
		Url:         event.Url,
		TicketURL:   event.TicketURL,
		Artists:     artists,
		Venue:       venue,
	}

	return e.event.CreateEvent(ctx, newEvent)
}

func (e *eventUsecase) GetEventsByName(name string) ([]model.Event, error) {
	return nil, nil
}

func (e *eventUsecase) GetEventByID(ctx context.Context, ID string) (*model.Event, error) {
	return e.event.GetEventByID(ctx, ID)
}

func (e *eventUsecase) DeleteEvent(ctx context.Context, event *model.Event) error {
	return e.event.DeleteEvent(ctx, event)
}

func (e *eventUsecase) UpdateEvent(ctx context.Context, event *model.UpdateEvent) (*model.Event, error) {

	if event.EventID == "" {
		return nil, &model.AppError{
			Code: 400,
			Msg:  "Event id is required",
		}
	}

	req, err := e.event.GetEventByID(ctx, event.EventID)

	if err != nil {
		return nil, &model.AppError{
			Code: 404,
			Msg:  "Event not found",
		}
	}

	var venue *model.Venue

	if event.VenueID != "" {
		venue, err = e.venue.GetVenueByID(ctx, event.VenueID)
		if err != nil {
			return nil, &model.AppError{
				Code: 404,
				Msg:  "Venue not found",
			}
		}
	}

	artists, err := e.artist.GetArtistsByIDs(ctx, event.ArtistIDs)
	if err != nil {
		return nil, &model.AppError{
			Code: 404,
			Msg:  "Artist ID is invalid",
		}
	}

	req.Artists = append(artists, req.Artists...)

	req.Name = event.Name
	req.Description = event.Description
	req.Date = event.Date
	req.OpenTime = event.OpenTime
	req.StartTime = event.StartTime
	req.EndTime = event.EndTime
	req.Venue = venue
	req.RelatedRyzmEvents = nil
	req.UnStructuredInformation = nil
	req.Artists = append(artists, req.Artists...)

	return e.event.UpdateEvent(ctx, req)
}

func (e *eventUsecase) MergeEvents(ctx context.Context, req *model.MergeEvent) (*model.Event, error) {
	base, err := e.event.GetEventByID(ctx, req.EventID)

	if err != nil {
		return nil, &model.AppError{
			Code: 404,
			Msg:  "Event not found",
		}
	}

	merge, err := e.event.GetEventByID(ctx, req.MergeTargetEventID)

	if err != nil {
		return nil, &model.AppError{
			Code: 404,
			Msg:  "Event not found",
		}
	}

	base.Artists = append(base.Artists, merge.Artists...)
	base.RelatedRyzmEvents = append(base.RelatedRyzmEvents, merge.RelatedRyzmEvents...)

	return e.event.MergeEvents(ctx, base, merge)

}

func (e *eventUsecase) SearchEvents(ctx context.Context, q *model.EventSearchQuery) ([]*model.Event, error) {
	result, err := e.event.SearchEvents(ctx, q)

	if err != nil {
		return nil, &model.AppError{
			Code: 500,
			Msg:  "Failed to search events",
		}
	}

	var events []*model.Event

	for _, r := range result {
		r.Event.Artists = []*model.Artist{&r.Artist}
		events = append(events, &r.Event)
	}

	return events, nil

}
