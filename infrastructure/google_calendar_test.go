package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/scarlet0725/prism-api/model"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

const ()

type testCredential struct {
	ClientID     string    `json:"client_id"`
	ClientSecret string    `json:"client_secret"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	Expiry       time.Time `json:"expiry"`
}

func TestGoogleCalender(t *testing.T) {
	//TODO: テストを書く

	dir, err := os.Getwd()

	if err != nil {
		t.Fatal(err)
	}

	dir = dir + "/testdata/config.json"

	config, err := os.ReadFile(dir)

	var credential testCredential

	if err := json.Unmarshal(config, &credential); err != nil {
		t.Fatal(err)
	}

	conf := &oauth2.Config{
		ClientID:     credential.ClientID,
		ClientSecret: credential.ClientSecret,
		RedirectURL:  fmt.Sprintf("%s/v1/oauth/google/callback", "http://localhost:8080"),
		Scopes: []string{
			calendar.CalendarScope,
			calendar.CalendarEventsScope,
		},
		Endpoint: google.Endpoint,
	}

	token := &oauth2.Token{
		AccessToken:  credential.AccessToken,
		TokenType:    "Bearer",
		RefreshToken: credential.RefreshToken,
		Expiry:       credential.Expiry,
	}

	t.Log(credential.Expiry.Format(time.RFC3339))

	client := conf.Client(context.TODO(), token)

	srv := NewGoogleCalenderClient(client)

	createdCalendar := &model.ExternalCalendar{}

	t.Run(
		"カレンダーを作成できる",
		func(t *testing.T) {
			createdCalendar, err = srv.CreateCalendar(&model.ExternalCalendar{
				Name: "test",
			})
			if err != nil {
				t.Fatal(err)
			}
		},
	)

	t.Logf("Created Calendar: %v", *createdCalendar)

}
