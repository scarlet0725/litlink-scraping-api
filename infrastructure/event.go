package infrastructure

import (
	"context"
	"time"

	"github.com/scarlet0725/prism-api/ent"
	"github.com/scarlet0725/prism-api/ent/artist"
	"github.com/scarlet0725/prism-api/ent/event"
	"github.com/scarlet0725/prism-api/ent/ryzmevent"
	"github.com/scarlet0725/prism-api/infrastructure/repository"
	"github.com/scarlet0725/prism-api/infrastructure/translator"
	"github.com/scarlet0725/prism-api/model"
)

type Event struct {
	db *ent.Client
}

func NewEventRepository(db *ent.Client) repository.Event {
	return &Event{
		db: db,
	}
}

func (e *Event) CreateEvent(ctx context.Context, event *model.Event) (*model.Event, error) {
	result, err := e.db.Event.Create().
		SetEventID(event.EventID).
		SetName(event.Name).
		SetDescription(event.Description).
		SetNillableDate(event.Date).
		SetNillableStartTime(event.StartTime).
		SetNillableEndTime(event.EndTime).
		SetURL(event.Url).
		SetTicketURL(event.TicketURL).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	createdEvent := translator.EventFromEnt(result)

	return createdEvent, nil
}

func (e *Event) UpdateEvent(ctx context.Context, event *model.Event) (*model.Event, error) {
	result, err := e.db.Event.UpdateOneID(int(event.ID)).
		SetEventID(event.EventID).
		SetName(event.Name).
		SetNillableDescription(&event.Description).
		SetNillableDate(event.Date).
		SetNillableStartTime(event.StartTime).
		SetNillableEndTime(event.EndTime).
		SetURL(event.Url).
		SetTicketURL(event.TicketURL).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	updatedEvent := translator.EventFromEnt(result)

	return updatedEvent, nil
}

func (e *Event) DeleteEvent(ctx context.Context, event *model.Event) error {
	err := e.db.Event.DeleteOneID(int(event.ID)).Exec(ctx)

	return err
}

func (e *Event) GetEventsByArtistID(ctx context.Context, artistID string) ([]*model.Event, error) {
	result, err := e.db.Event.Query().
		Where(
			event.HasArtistsWith(artist.ArtistID(artistID)),
		).
		All(ctx)

	if err != nil {
		return nil, err
	}

	events := make([]*model.Event, 0, len(result))

	for _, event := range result {
		events = append(events, translator.EventFromEnt(event))
	}

	return events, nil
}

func (e *Event) GetEventByID(ctx context.Context, eventID string) (*model.Event, error) {
	result, err := e.db.Event.Query().
		Where(
			event.And(
				event.DeletedAtIsNil(),
				event.EventID(eventID),
			),
		).
		First(ctx)

	if err != nil {
		return nil, err
	}

	event := translator.EventFromEnt(result)

	return event, nil
}

func (e *Event) GetEvents(ctx context.Context) ([]*model.Event, error) {
	result, err := e.db.Event.Query().
		Where(
			event.DeletedAtNotNil(),
		).
		All(ctx)

	if err != nil {
		return nil, err
	}

	events := make([]*model.Event, 0, len(result))

	for _, event := range result {
		events = append(events, translator.EventFromEnt(event))
	}

	return events, nil
}

func (e *Event) MergeEvents(ctx context.Context, base *model.Event, target *model.Event) (*model.Event, error) {

	ids := make([]int, 0, (len(target.Artists) + len(base.Artists)))

	for _, artist := range target.Artists {
		ids = append(ids, int(artist.ID))
	}

	for _, artist := range base.Artists {
		ids = append(ids, int(artist.ID))
	}

	tx, err := e.db.Tx(ctx)

	if err != nil {
		return nil, err
	}

	//TODO: テスト厚めで
	result, err := tx.Event.UpdateOneID(int(base.ID)).
		AddArtistIDs(ids...).
		Save(ctx)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	//統合先のイベントに統合元のイベントの会場が設定されていない場合は統合先のイベントに統合元のイベントの会場を設定する
	if result.Edges.Venue == nil && target.Venue != nil {
		_, err = tx.Event.UpdateOneID(int(base.ID)).
			SetVenueID(int(target.Venue.ID)).
			Save(ctx)

		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	//統合元のイベントを削除する(ソフトデリート)
	err = tx.Event.UpdateOneID(int(target.ID)).
		SetDeletedAt(time.Now()).
		Exec(ctx)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	event := translator.EventFromEnt(result)

	return event, nil
}

func (e *Event) GetRyzmEventsByUUDIDs(ctx context.Context, IDs []string) ([]*model.RyzmEvent, error) {
	result, err := e.db.RyzmEvent.Query().Where(
		ryzmevent.UUIDIn(IDs...),
	).All(ctx)

	if err != nil {
		return nil, err
	}

	events := make([]*model.RyzmEvent, 0, len(result))

	for _, event := range result {
		events = append(events, translator.RyzmEventFromEnt(event))
	}

	return events, nil
}

func (e *Event) SearchEvents(ctx context.Context, query *model.EventSearchQuery) ([]*model.EventSearchResult, error) {
	//TODO: あとで実装
	return nil, nil
}
