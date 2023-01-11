package translator

import (
	"time"

	"github.com/scarlet0725/prism-api/ent"
	"github.com/scarlet0725/prism-api/model"
)

func ArtistFromEnt(source *ent.Artist) *model.Artist {
	artist := &model.Artist{
		ID:       uint(source.ID),
		ArtistID: source.ArtistID,
		Name:     source.Name,
		URL:      source.URL,
	}

	return artist
}

func UserFromEnt(source *ent.User) *model.User {
	user := &model.User{
		ID:               uint(source.ID),
		UserID:           source.UserID,
		Username:         source.Username,
		FamilyName:       source.FirstName,
		GivenName:        source.LastName,
		Email:            source.Email,
		Password:         source.Password,
		APIKey:           source.APIKey,
		IsAdminVerified:  source.IsAdminVerified,
		DeleteProtected:  source.DeleteProtected,
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
		DeletedAt:        &time.Time{},
		Events:           []*model.Event{},
		Roles:            []*model.Role{},
		GoogleToken:      &model.GoogleOAuthToken{},
		GoogleOAuthState: &model.GoogleOAuthState{},
		ExternalCalendar: &model.ExternalCalendar{},
	}

	return user
}

func ExternalCalendarFromEnt(source *ent.ExternalCalendar) *model.ExternalCalendar {
	externalCalendar := &model.ExternalCalendar{
		ID:          source.ID,
		Name:        source.Name,
		Description: source.Description,
		CalendarID:  source.CalendarID,
		Type:        source.SourceType,
		UserID:      source.UserID,
		CreatedAt:   source.CreatedAt,
		UpdatedAt:   source.UpdatedAt,
		DeletedAt:   source.DeletedAt,
	}

	return externalCalendar
}
