package repository

import (
	"context"

	"github.com/scarlet0725/prism-api/model"
)

type User interface {
	CreateUser(context.Context, *model.User) (*model.User, error)
	GetUser(context.Context, string) (*model.User, error)
	UpdateUser(context.Context, *model.User) (*model.User, error)
	GetUserByAPIKey(context.Context, string) (*model.User, error)
	DeleteUser(context.Context, *model.User) error
	GetUserCalendarByUserID(context.Context, int) (*model.ExternalCalendar, error)
	GetGoogleOAuthToken(context.Context, int) (*model.GoogleOAuthToken, error)
	SaveExternalCalendar(context.Context, *model.ExternalCalendar) (*model.ExternalCalendar, error)
	AddRegistrationEvent(context.Context, *model.User, *model.Event) error
	RemoveRegistrationEvent(context.Context, *model.User, *model.Event) error
	GetGoogleCalendarConfig(context.Context, int) (*model.GoogleCalenderConfig, error)
	GetUserByUsername(context.Context, string) (*model.User, error)
}
