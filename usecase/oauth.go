package usecase

import (
	"context"

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
	jwt    framework.JWTHandler
}

func NewOAuthUsecase(oauth repository.OAuth, ramdom framework.RandomID, google framework.GoogleOAuth) OAuth {
	return &oauthUsecase{
		oauth:  oauth,
		ramdom: ramdom,
		google: google,
	}
}

func (a *oauthUsecase) GoogleLinkage(ctx context.Context, user *model.User) (*model.OAuthURLResponse, error) {
	state, err := a.jwt.CreateStateToken(user)

	if err != nil {
		return nil, err
	}
	u := a.google.GenerateAuthURL(state)

	o := &model.OAuthURLResponse{
		AuthURL: u,
		State:   state,
	}

	return o, nil
}

func (a *oauthUsecase) GoogleOAuthCallback(ctx context.Context, callback *model.GoogleOauthCallback) error {
	s, err := a.jwt.ValidateToken(callback.State)

	if err != nil {
		return err
	}

	t, err := a.google.GetToken(callback.Code)

	if err != nil {
		return err
	}
	u := &model.User{}
	t.User = u
	t.User.ID = s.ID
	t.User.UserID = s.UserID

	_, err = a.oauth.SaveGoogleOAuthToken(ctx, t)

	if err != nil {
		return err
	}

	return nil
}
