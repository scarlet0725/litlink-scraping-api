package framework

import (
	"context"
	"net/http"

	"github.com/scarlet0725/prism-api/model"
	"golang.org/x/oauth2"
)

type GoogleOAuth interface {
	GenerateAuthURL(state string) string
	GetToken(code string) (*model.GoogleOAuthToken, error)
	GetClient(token *model.GoogleOAuthToken) *http.Client
}

type GoogleOAuthImpl struct {
	config *oauth2.Config
}

func NewGoogleOAuth(config *oauth2.Config) GoogleOAuth {
	return &GoogleOAuthImpl{
		config: config,
	}

}

func (g *GoogleOAuthImpl) GenerateAuthURL(state string) string {
	return g.config.AuthCodeURL(state, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
}

func (g *GoogleOAuthImpl) GetToken(code string) (*model.GoogleOAuthToken, error) {
	token, err := g.config.Exchange(context.TODO(), code)
	if err != nil {
		return nil, err
	}

	t := &model.GoogleOAuthToken{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
	}

	return t, nil
}

func (g *GoogleOAuthImpl) GetClient(token *model.GoogleOAuthToken) *http.Client {
	t := &oauth2.Token{
		AccessToken:  token.AccessToken,
		TokenType:    "Bearer",
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
	}

	c := g.config.Client(context.TODO(), t)

	return c
}
