package translator

import (
	"github.com/scarlet0725/prism-api/ent"
	"github.com/scarlet0725/prism-api/model"
)

func ArtistFromEnt(source *ent.Artist) *model.Artist {
	artist := &model.Artist{
		ID:       source.ID,
		ArtistID: source.ArtistID,
		Name:     source.Name,
		URL:      source.URL,
	}

	return artist
}

func ExternalCalendarFromEnt(source *ent.ExternalCalendar) *model.ExternalCalendar {
	externalCalendar := &model.ExternalCalendar{
		ID:          source.ID,
		Name:        source.Name,
		Description: source.Description,
		CalendarID:  source.CalendarID,
		Type:        source.SourceType,
		CreatedAt:   source.CreatedAt,
		UpdatedAt:   source.UpdatedAt,
		DeletedAt:   source.DeletedAt,
	}

	if source.Edges.User != nil {
		externalCalendar.UserID = source.Edges.User.ID
	}

	return externalCalendar
}

func EventFromModel(source *model.Event) *ent.Event {
	event := &ent.Event{
		ID:          int(source.ID),
		EventID:     source.EventID,
		Name:        source.Name,
		Date:        source.Date,
		OpenTime:    source.OpenTime,
		StartTime:   source.StartTime,
		EndTime:     source.EndTime,
		Description: source.Description,
		URL:         source.Url,
		TicketURL:   source.TicketURL,
	}

	return event
}

func EventFromEnt(source *ent.Event) *model.Event {
	event := &model.Event{
		ID:          source.ID,
		EventID:     source.EventID,
		Name:        source.Name,
		Date:        source.Date,
		OpenTime:    source.OpenTime,
		StartTime:   source.StartTime,
		EndTime:     source.EndTime,
		Description: source.Description,
		Url:         source.URL,
		TicketURL:   source.TicketURL,
	}

	return event
}

func ExternalCalendarFromModel(source *model.ExternalCalendar) *ent.ExternalCalendar {
	externalcalendar := &ent.ExternalCalendar{
		Name:        source.Name,
		Description: source.Description,
		CalendarID:  source.CalendarID,
		SourceType:  source.Type,
	}

	return externalcalendar
}

func GoogleOAuthTokenFromEnt(source *ent.GoogleOauthToken) *model.GoogleOAuthToken {
	googleOAuthToken := &model.GoogleOAuthToken{
		RefreshToken: source.RefreshToken,
		AccessToken:  source.AccessToken,
		Expiry:       source.Expiry,
	}

	return googleOAuthToken
}

func EventFromEnts(source *ent.Event) *model.Event {
	events := model.Event{}

	events.ID = source.ID
	events.EventID = source.EventID
	events.Name = source.Name
	events.Date = source.Date
	events.OpenTime = source.OpenTime
	events.StartTime = source.StartTime
	events.EndTime = source.EndTime
	events.Description = source.Description
	events.Url = source.URL
	events.TicketURL = source.TicketURL

	return &events
}

func VenueFromEnt(source *ent.Venue) *model.Venue {
	venue := &model.Venue{
		ID:          uint(source.ID),
		VenueID:     source.VenueID,
		Name:        source.Name,
		Description: source.Description,
		WebSite:     source.WebSite,
		Postcode:    source.Postcode,
		Prefecture:  source.Prefecture,
		City:        source.City,
		Street:      source.Street,
		IsOpen:      source.IsOpen,
		Events:      []*model.Event{},
	}

	if source.Edges.Events == nil {
		return venue
	}

	e := make([]*model.Event, len(source.Edges.Events))

	for _, v := range source.Edges.Events {
		e = append(e, EventFromEnt(v))
	}

	venue.Events = e

	return venue
}
