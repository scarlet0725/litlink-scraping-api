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
	CreateArtistEventsFromCrawlData(ctx context.Context, id string) ([]*model.Event, error)
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

func (e *eventUsecase) CreateArtistEventsFromCrawlData(ctx context.Context, id string) ([]*model.Event, error) {
	artist, err := e.artist.GetArtistByID(ctx, id)

	if err != nil {
		return nil, &model.AppError{
			Code: 404,
			Msg:  "artist_not_found",
		}
	}

	if artist.RyzmCrawlConfig == nil {
		return nil, &model.AppError{
			Code: 400,
			Msg:  "This artist is not supported auto update",
		}
	}

	req := &model.ScrapingRequest{
		Host: artist.RyzmCrawlConfig.CrawlTargetURL,
		URL:  artist.RyzmCrawlConfig.CrawlTargetURL,
		Type: artist.RyzmCrawlConfig.CrawlSiteType,
		Option: model.FetchOptions{
			IsUseCache: true,
		},
	}

	if req.Type == "ryzm" {
		req.Option = model.FetchOptions{
			HTTPHeader: map[string]string{
				"X-RYZM-HOST": artist.RyzmCrawlConfig.RyzmHost,
			},
			HTTPParams: map[string]string{
				"archived": "0",
			},
			CacheKey: artist.RyzmCrawlConfig.RyzmHost,
		}
	}

	res, err := e.fetch.Fetch(req)

	if err != nil {
		return nil, &model.AppError{
			Code: 500,
			Msg:  "fetch_error",
		}
	}

	switch artist.RyzmCrawlConfig.CrawlSiteType {
	case "ryzm":

		json, err := e.json.Ryzm(res.Data)
		if err != nil {
			return nil, &model.AppError{
				Code: 500,
				Msg:  "failed_to_parse_json",
			}
		}
		result, err := e.selializer.SelializeRyzmData(json)

		if err != nil {
			return nil, &model.AppError{
				Code: 500,
				Msg:  "failed_to_parse_json",
			}
		}

		//Ryzmから取得したイベントのUUIDをsliceに格納しすでに登録されているかを確認する
		fetchedRyzmEventUUIDs := []string{}
		for _, event := range result {
			for _, v := range event.RelatedRyzmEvents {
				fetchedRyzmEventUUIDs = append(fetchedRyzmEventUUIDs, v.UUID)
			}
		}

		registeredEvents, err := e.event.GetRyzmEventsByUUDIDs(ctx, fetchedRyzmEventUUIDs)

		if err != nil {
			return nil, &model.AppError{
				Code: 500,
				Msg:  "failed_to_get_events",
			}
		}

		//すでに登録されているイベントのUUIDをmapに格納
		existedEventUUID := map[string]bool{}
		for _, event := range registeredEvents {
			existedEventUUID[event.UUID] = true
		}

		registrationExpectedEvents := []*model.Event{}

		//登録されているかどうかmapから確認し、登録されていないイベントを登録する

		var c int
		for _, event := range result {
			c = 0
			for _, ryzm := range event.RelatedRyzmEvents {
				_, ok := existedEventUUID[ryzm.UUID]
				if ok {
					c++
				}
			}
			if c == 0 {
				event.EventID = e.random.Generate(eventIDLength)
				event.Artists = append(event.Artists, artist)
				registrationExpectedEvents = append(registrationExpectedEvents, event)
			}
		}

		if len(registrationExpectedEvents) <= 0 {
			return []*model.Event{}, nil
		}

		//会場を登録出来たらする
		//会場の登録は会場の名前で行う

		//会場の名前をsliceに格納
		venueNames := []string{}
		for _, event := range registrationExpectedEvents {
			venueNames = append(venueNames, event.UnStructuredInformation.VenueName)
		}

		//会場の名前から会場を取得
		venues, err := e.venue.GetVenuesByNames(ctx, venueNames)
		if err != nil {
			return nil, &model.AppError{
				Code: 500,
				Msg:  "Failed to get venues",
			}
		}
		venueMap := map[string]*model.Venue{}

		for _, venue := range venues {
			venueMap[venue.Name] = venue
		}

		//会場が登録されていない場合は会場を登録する

		for _, event := range registrationExpectedEvents {
			venue, ok := venueMap[event.UnStructuredInformation.VenueName]
			if !ok {
				continue
			}
			event.Venue = venue
		}

		return e.event.CreateEvents(ctx, registrationExpectedEvents)

	default:
		return nil, &model.AppError{
			Code: 500,
			Msg:  "error!",
		}

	}

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
