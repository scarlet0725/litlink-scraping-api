package repository

import (
	"context"

	"github.com/scarlet0725/prism-api/model"
)

type OAuth interface {
	SaveGoogleOAuthToken(context.Context, *model.GoogleOAuthToken) (*model.GoogleOAuthToken, error)
}
