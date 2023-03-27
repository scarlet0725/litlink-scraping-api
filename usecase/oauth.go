package usecase

import (
	"context"
	"errors"

	"github.com/scarlet0725/prism-api/framework"
	"github.com/scarlet0725/prism-api/infrastructure/repository"
	"github.com/scarlet0725/prism-api/model"
)

type OAuth interface {
	GoogleLinkage(context.Context, *model.User) (*model.OAuthURLResponse, error)
	GoogleOAuthCallback(context.Context, *model.GoogleOauthCallback) error
}

type oauthUsecase struct {
	ramdom framework.RandomID
	google framework.GoogleOAuth
	oauth  repository.OAuth
}

func NewOAuthUsecase(oauth repository.OAuth, ramdom framework.RandomID, google framework.GoogleOAuth) OAuth {
	return &oauthUsecase{
		oauth:  oauth,
		ramdom: ramdom,
		google: google,
	}
}

func (a *oauthUsecase) GoogleLinkage(ctx context.Context, user *model.User) (*model.OAuthURLResponse, error) {
	state, err := a.ramdom.GenerateUUID4()
	if err != nil {
		return nil, err
	}
	u := a.google.GenerateAuthURL(state)

	s := &model.GoogleOAuthState{
		UserID: user.ID,
		State:  state,
	}

	_, err = a.oauth.SaveGoogleOAuthState(ctx, s)

	if err != nil {
		return nil, err
	}

	o := &model.OAuthURLResponse{
		AuthURL: u,
		State:   state,
	}

	return o, nil
}

func (a *oauthUsecase) GoogleOAuthCallback(ctx context.Context, callback *model.GoogleOauthCallback) error {
	s, err := a.oauth.GetGoogleOAuthState(ctx, callback.State)
	if err != nil {
		return err
	}

	if s.State != callback.State {
		return errors.New("invalid state")
	}

	t, err := a.google.GetToken(callback.Code)

	t.UserID = s.UserID

	if err != nil {
		return err
	}

	_, err = a.oauth.SaveGoogleOAuthToken(ctx, t)

	if err != nil {
		return err
	}

	return nil
}
