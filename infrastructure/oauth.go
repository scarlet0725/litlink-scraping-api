package infrastructure

import (
	"context"

	"github.com/scarlet0725/prism-api/ent"
	"github.com/scarlet0725/prism-api/infrastructure/repository"
	"github.com/scarlet0725/prism-api/model"
)

type oauth struct {
	db *ent.Client
}

func NewOAuthRepository(db *ent.Client) repository.OAuth {
	return &oauth{
		db: db,
	}
}

func (o *oauth) SaveGoogleOAuthState(ctx context.Context, state *model.GoogleOAuthState) (*model.GoogleOAuthState, error) {
	result, err := o.db.GoogleOauthState.Create().SetUserID(int(state.UserID)).SetState(state.State).Save(ctx)

	if err != nil {
		return nil, err
	}

	s := model.GoogleOAuthState{
		ID:     result.ID,
		UserID: result.Edges.User.ID,
		State:  result.State,
	}

	return &s, nil
}

func (o *oauth) GetGoogleOAuthState(ctx context.Context, state string) (*model.GoogleOAuthState, error) {
	result, err := o.db.GoogleOauthState.Query().Where().Only(ctx)

	if err != nil {
		return nil, err
	}

	s := model.GoogleOAuthState{
		ID:     result.ID,
		UserID: result.Edges.User.ID,
		State:  result.State,
	}

	return &s, nil
}

func (o *oauth) SaveGoogleOAuthToken(ctx context.Context, token *model.GoogleOAuthToken) (*model.GoogleOAuthToken, error) {
	result, err := o.db.GoogleOauthToken.Create().
		SetUserID(int(token.UserID)).
		SetAccessToken(token.AccessToken).
		SetRefreshToken(token.RefreshToken).
		SetExpiry(token.Expiry).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	t := model.GoogleOAuthToken{
		ID:           result.ID,
		UserID:       result.Edges.User.ID,
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
		Expiry:       result.Expiry,
	}

	return &t, nil
}
