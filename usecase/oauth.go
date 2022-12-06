package usecase

import (
	"errors"

	"github.com/scarlet0725/prism-api/framework"
	"github.com/scarlet0725/prism-api/infrastructure/repository"
	"github.com/scarlet0725/prism-api/model"
)

type OAuth interface {
	GoogleLinkage(*model.User) (*model.OAuthURLResponse, error)
	GoogleOAuthCallback(*model.GoogleOauthCallback) error
}

type oauthUsecase struct {
	db     repository.DB
	ramdom framework.RamdomIDGenerator
	google framework.GoogleOAuth
}

func NewOAuthApplication(db repository.DB, ramdom framework.RamdomIDGenerator, google framework.GoogleOAuth) OAuth {
	return &oauthUsecase{
		ramdom: ramdom,
		db:     db,
		google: google,
	}
}

func (a *oauthUsecase) GoogleLinkage(user *model.User) (*model.OAuthURLResponse, error) {
	state, err := a.ramdom.GenerateUUID4()
	if err != nil {
		return nil, err
	}
	u := a.google.GenerateAuthURL(state)

	s := &model.GoogleOAuthState{
		UserID: &user.ID,
		State:  state,
	}

	_, err = a.db.SaveGoogleOAuthState(s)

	if err != nil {
		return nil, err
	}

	o := &model.OAuthURLResponse{
		AuthURL: u,
		State:   state,
	}

	return o, nil
}

func (a *oauthUsecase) GoogleOAuthCallback(callback *model.GoogleOauthCallback) error {
	s, err := a.db.GetGoogleOAuthStateByState(callback.State)
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

	_, err = a.db.SaveGoogleOAuthToken(t)

	if err != nil {
		return err
	}

	return nil
}
