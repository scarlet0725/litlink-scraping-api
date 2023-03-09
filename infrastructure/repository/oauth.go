package repository

import (
	"context"

	"github.com/scarlet0725/prism-api/model"
)

type OAuth interface {
	SaveGoogleOAuthState(context.Context, *model.GoogleOAuthState) (*model.GoogleOAuthState, error)
	GetGoogleOAuthState(ctx context.Context, state string) (*model.GoogleOAuthState, error)
	SaveGoogleOAuthToken(context.Context, *model.GoogleOAuthToken) (*model.GoogleOAuthToken, error)
}
