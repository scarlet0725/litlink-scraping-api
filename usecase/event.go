package usecase

import (
	"time"

	"github.com/scarlet0725/prism-api/cmd"
	"github.com/scarlet0725/prism-api/controller"
	"github.com/scarlet0725/prism-api/infrastructure/repository"
	"github.com/scarlet0725/prism-api/model"
	"github.com/scarlet0725/prism-api/parser"
	"github.com/scarlet0725/prism-api/selializer"
)

type Event interface {
	CreateEvent(*model.CreateEvent) (*model.Event, error)
	DeleteEvent(*model.Event) error
	UpdateEvent(*model.UpdateEvent) (*model.Event, error)
	//GetEvent(string) (*model.Event, error)
	GetEventsByArtistName(string) ([]*model.Event, *model.AppError)
	CreateArtistEventsFromCrawlData(id string) ([]*model.Event, error)
	GetEventByID(string) (*model.Event, error)
	MergeEvents(*model.MergeEvent) (*model.Event, error)
}

type eventUsecase struct {
	db         repository.DB
	fetch      controller.FetchController
	parser     parser.DocParser
	selializer selializer.ResponseSerializer
	json       parser.JsonParser
}

func NewEventApplication(db repository.DB, fetch controller.FetchController, parser parser.DocParser, selializer selializer.ResponseSerializer, json parser.JsonParser) Event {
	return &eventUsecase{
		db:         db,
		fetch:      fetch,
		parser:     parser,
		selializer: selializer,
		json:       json,
	}
}

func (a *eventUsecase) CreateEvent(e *model.CreateEvent) (*model.Event, error) {
	id := cmd.MakeRamdomID(eventIDLength)

	artists, _ := a.db.GetArtistsByIDs(e.ArtistIDs)
	venue, _ := a.db.GetVenueByID(e.VenueID)
	jst, _ := time.LoadLocation(Locale)

	date, err := time.ParseInLocation("2006-01-02", e.Date, jst)

	if err != nil {
		return nil, err
	}

	event := &model.Event{
		EventID:     id,
		Name:        e.Name,
		Date:        &date,
		Description: e.Description,
		OpenTime:    e.OpenTime,
		StartTime:   e.StartTime,
		EndTime:     e.EndTime,
		Url:         e.Url,
		TicketURL:   e.TicketURL,
		Artists:     artists,
		Venue:       venue,
	}

	return a.db.CreateEvent(event)
}

func (a *eventUsecase) GetEventsByName(name string) ([]model.Event, error) {
	return nil, nil
}

func (a *eventUsecase) GetEventsByArtistName(name string) ([]*model.Event, *model.AppError) {
	artist, err := a.db.GetArtistByName(name)

	if err != nil {
		return nil, &model.AppError{
			Code: 404,
			Msg:  "artist_not_found",
		}
	}

	req := &model.ScrapingRequest{
		Host: artist.CrawlTargetURL,
		URL:  artist.CrawlTargetURL,
		Type: artist.CrawlSiteType,
		Option: model.FetchOptions{
			IsUseCache: true,
		},
	}

	if req.Type == "ryzm" {
		req.Option = model.FetchOptions{
			HTTPHeader: map[string]string{
				"X-RYZM-HOST": artist.RyzmHost,
			},
			HTTPParams: map[string]string{
				"archived": "0",
			},
		}
	}

	res, err := a.fetch.Fetch(req)

	if err != nil {
		return nil, &model.AppError{
			Code: 500,
			Msg:  "fetch_error",
		}
	}

	switch artist.CrawlSiteType {
	case "ryzm":
		json, err := a.json.Ryzm(res.Data)
		if err != nil {
			return nil, &model.AppError{
				Code: 500,
				Msg:  "failed_to_parse_json",
			}
		}
		result, err := a.selializer.SelializeRyzmData(json)

		if err != nil {
			return nil, &model.AppError{
				Code: 500,
				Msg:  "failed_to_parse_json",
			}
		}

		return result, nil
	default:
		return nil, &model.AppError{
			Code: 500,
			Msg:  "error!",
		}

	}
}

func (a *eventUsecase) CreateArtistEventsFromCrawlData(id string) ([]*model.Event, error) {
	artist, err := a.db.GetArtistByID(id)

	if err != nil {
		return nil, &model.AppError{
			Code: 404,
			Msg:  "artist_not_found",
		}
	}

	if artist.CrawlTargetURL == "" || artist.CrawlSiteType == "" || artist.RyzmHost == "" {
		return nil, &model.AppError{
			Code: 400,
			Msg:  "This artist is not supported auto update",
		}
	}

	req := &model.ScrapingRequest{
		Host: artist.CrawlTargetURL,
		URL:  artist.CrawlTargetURL,
		Type: artist.CrawlSiteType,
		Option: model.FetchOptions{
			IsUseCache: true,
		},
	}

	if req.Type == "ryzm" {
		req.Option = model.FetchOptions{
			HTTPHeader: map[string]string{
				"X-RYZM-HOST": artist.RyzmHost,
			},
			HTTPParams: map[string]string{
				"archived": "0",
			},
			CacheKey: artist.RyzmHost,
		}
	}

	res, err := a.fetch.Fetch(req)

	if err != nil {
		return nil, &model.AppError{
			Code: 500,
			Msg:  "fetch_error",
		}
	}

	switch artist.CrawlSiteType {
	case "ryzm":

		json, err := a.json.Ryzm(res.Data)
		if err != nil {
			return nil, &model.AppError{
				Code: 500,
				Msg:  "failed_to_parse_json",
			}
		}
		result, err := a.selializer.SelializeRyzmData(json)

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

		registeredEvents, err := a.db.GetRyzmEventsByUUDIDs(fetchedRyzmEventUUIDs)

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
				event.EventID = cmd.MakeRamdomID(eventIDLength)
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
		venues, err := a.db.GetVenuesByNames(venueNames)
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

		return a.db.CreateEvents(registrationExpectedEvents)

	default:
		return nil, &model.AppError{
			Code: 500,
			Msg:  "error!",
		}

	}

}

func (a *eventUsecase) GetEventByID(ID string) (*model.Event, error) {
	return a.db.GetEventByID(ID)
}

func (a *eventUsecase) DeleteEvent(event *model.Event) error {
	return a.db.DeleteEvent(event)
}

func (a *eventUsecase) UpdateEvent(event *model.UpdateEvent) (*model.Event, error) {

	if event.EventID == "" {
		return nil, &model.AppError{
			Code: 400,
			Msg:  "Event id is required",
		}
	}

	req, err := a.db.GetEventByID(event.EventID)

	if err != nil {
		return nil, &model.AppError{
			Code: 404,
			Msg:  "Event not found",
		}
	}

	var venue *model.Venue

	if event.VenueID != "" {
		venue, err = a.db.GetVenueByID(event.VenueID)
		if err != nil {
			return nil, &model.AppError{
				Code: 404,
				Msg:  "Venue not found",
			}
		}
	}

	req.Name = event.Name
	req.Description = event.Description
	req.Date = event.Date
	req.OpenTime = event.OpenTime
	req.StartTime = event.StartTime
	req.EndTime = event.EndTime
	req.Venue = venue

	return a.db.UpdateEvent(req)
}

func (a *eventUsecase) MergeEvents(req *model.MergeEvent) (*model.Event, error) {
	base, err := a.db.GetEventByID(req.EventID)

	if err != nil {
		return nil, &model.AppError{
			Code: 404,
			Msg:  "Event not found",
		}
	}

	merge, err := a.db.GetEventByID(req.MergeTargetEventID)

	if err != nil {
		return nil, &model.AppError{
			Code: 404,
			Msg:  "Event not found",
		}
	}

	base.Artists = append(base.Artists, merge.Artists...)
	base.RelatedRyzmEvents = append(base.RelatedRyzmEvents, merge.RelatedRyzmEvents...)

	return a.db.MergeEvents(base, merge)

}
