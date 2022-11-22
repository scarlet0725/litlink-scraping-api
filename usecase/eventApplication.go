package usecase

import (
	"fmt"

	"github.com/scarlet0725/prism-api/cmd"
	"github.com/scarlet0725/prism-api/controller"
	"github.com/scarlet0725/prism-api/model"
	"github.com/scarlet0725/prism-api/parser"
	"github.com/scarlet0725/prism-api/repository"
	"github.com/scarlet0725/prism-api/selializer"
)

type EventApplication interface {
	CreateEvent(*model.CreateEvent) (*model.Event, error)
	DeleteEvent(*model.Event) error
	//GetEvent(string) (*model.Event, error)
	GetEventsByArtistName(string) ([]model.Event, *model.AppError)
	UpdateArtistLatestEventInformation(string) (*model.Event, error)
	GetEventByID(string) (*model.Event, error)
}

type eventApplication struct {
	db         repository.DB
	fetch      controller.FetchController
	parser     parser.DocParser
	selializer selializer.ResponseSerializer
	json       parser.JsonParser
}

func NewEventApplication(db repository.DB, fetch controller.FetchController, parser parser.DocParser, selializer selializer.ResponseSerializer, json parser.JsonParser) EventApplication {
	return &eventApplication{
		db:         db,
		fetch:      fetch,
		parser:     parser,
		selializer: selializer,
		json:       json,
	}
}

func (a *eventApplication) CreateEvent(e *model.CreateEvent) (*model.Event, error) {
	id := cmd.MakeRamdomID(eventIDLength)

	artists, _ := a.db.GetArtistsByIDs(e.ArtistIDs)

	fmt.Println(artists)

	event := &model.Event{
		EventID:     id,
		Name:        e.Name,
		Date:        e.Date,
		Description: e.Description,
		OpenTime:    e.OpenTime,
		StartTime:   e.StartTime,
		EndTime:     e.EndTime,
		Url:         e.Url,
		TicketURL:   e.TicketURL,
		Artists:     artists,
	}
	fmt.Println(event)

	return a.db.CreateEvent(event)
}

func (a *eventApplication) GetEventsByName(name string) ([]model.Event, error) {
	return nil, nil
}

func (a *eventApplication) GetEventsByArtistName(name string) ([]model.Event, *model.AppError) {
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
			return []model.Event{}, &model.AppError{
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

func (a *eventApplication) UpdateArtistLatestEventInformation(id string) (*model.Event, error) {
	return nil, nil
}

func (a *eventApplication) GetEventByID(ID string) (*model.Event, error) {
	return a.db.GetEventByID(ID)
}

func (a *eventApplication) DeleteEvent(event *model.Event) error {
	return a.db.DeleteEvent(event)
}
