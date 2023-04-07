package infrastructure

import (
	"context"
	"fmt"

	"github.com/scarlet0725/prism-api/ent"
	"github.com/scarlet0725/prism-api/infrastructure/repository"
	"github.com/scarlet0725/prism-api/model"
)

var userError = fmt.Errorf("invalid user")

type oauth struct {
	db *ent.Client
}

func NewOAuthRepository(db *ent.Client) repository.OAuth {
	return &oauth{
		db: db,
	}
}

func (o *oauth) SaveGoogleOAuthToken(ctx context.Context, token *model.GoogleOAuthToken) (*model.GoogleOAuthToken, error) {
	if token == nil || token.User == nil {
		return nil, userError
	}

	result, err := o.db.GoogleOauthToken.Create().
		SetUserID(int(token.User.ID)).
		SetAccessToken(token.AccessToken).
		SetRefreshToken(token.RefreshToken).
		SetExpiry(token.Expiry).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	t := model.GoogleOAuthToken{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
		Expiry:       result.Expiry,
		User:         token.User,
	}

	return &t, nil
}
