package cmd

import (
	"errors"
	"fmt"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

func GetGoogleOAuthConfig(host string) (*oauth2.Config, error) {
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		return nil, errors.New("env is not set")
	}

	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  fmt.Sprintf("%s/v1/oauth/google/callback", host),
		Scopes: []string{
			calendar.CalendarEventsScope,
			calendar.CalendarScope,
		},
		Endpoint: google.Endpoint,
	}

	return conf, nil
}
