package selializer

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/scarlet0725/prism-api/model"
)

func TestRyzmResponseSelializer(t *testing.T) {
	parsedPublishesAt, _ := time.Parse(time.RFC3339, "2022-11-10T04:58:00Z")
	parsedCreatedAt, _ := time.Parse(time.RFC3339, "2022-11-10T05:00:32Z")
	parsedUpdatedAt, _ := time.Parse(time.RFC3339, "2022-11-10T05:18:22Z")

	input := model.RyzmAPIResponse{
		Data: []model.RyzmLiveData{
			{
				ID:              "6c16f274-aa55-4ea0-9994-64020f25629d",
				Status:          "publish",
				EventDate:       "2022-11-17",
				EventDateStatus: "status_confirmed",
				CoverImage: struct {
					ID       string      "json:\"id\""
					URL      string      "json:\"url\""
					MimeType string      "json:\"mime_type\""
					FileName string      "json:\"file_name\""
					Width    string      "json:\"width\""
					Height   string      "json:\"height\""
					AltText  interface{} "json:\"alt_text\""
					Title    interface{} "json:\"title\""
				}{
					ID:       "a679ddb2-d547-492b-b259-3c5cdc48db98",
					URL:      "https://ryzm.imgix.net/sites/06ef9f50-4256-4b18-af8f-d3042a6fa7c3/images/9b3c2493-912f-4cd9-9e06-ef9d037aec68",
					MimeType: "image/jpeg",
					FileName: "D2CDA96A-5EEF-4F73-8725-5C4CFC266418.jpeg",
					Width:    "1500",
					Height:   "2121",
				},
				Category: struct {
					ID       int    "json:\"id\""
					Name     string "json:\"name\""
					Slug     string "json:\"slug\""
					Position int    "json:\"position\""
				}{
					ID:       349,
					Name:     "LIVE",
					Slug:     "live",
					Position: 1,
				},
				Venue:           "Shibuya Milkyway",
				Title:           "【単独無銭】プリズムセン",
				Artist:          "PRSMIN",
				DoorsStartsTime: "open/start 19:30/20:00",
				Price:           "¥0+1D",
				ReservationSetting: struct {
					TicketReservationType              string      "json:\"ticket_reservation_type\""
					WebReservationMaxQuantity          interface{} "json:\"web_reservation_max_quantity\""
					WebReservationMaxQuantityPerPerson interface{} "json:\"web_reservation_max_quantity_per_person\""
					Platforms                          []struct {
						ID  string "json:\"id\""
						URL string "json:\"url\""
					} "json:\"platforms\""
				}{
					TicketReservationType: "platform",
					Platforms: []struct {
						ID  string "json:\"id\""
						URL string "json:\"url\""
					}{
						{
							ID:  "other",
							URL: "https://t.livepocket.jp/e/g6fq1",
						},
					},
				},
				Body:        "<p>心が広い方向けの公演です</p><p><br></p><p>11/10 22:00~販売開始</p>",
				PublishesAt: parsedPublishesAt,
				Archived:    0,
				CreatedAt:   parsedCreatedAt,
				UpdatedAt:   parsedUpdatedAt,
			},
		},
		Links: struct {
			First string      "json:\"first\""
			Last  string      "json:\"last\""
			Prev  interface{} "json:\"prev\""
			Next  interface{} "json:\"next\""
		}{
			First: "https://api.ryzm.jp/public/lives?page=1",
			Last:  "https://api.ryzm.jp/public/lives?page=1",
		},
		Meta: struct {
			CurrentPage int "json:\"current_page\""
			From        int "json:\"from\""
			LastPage    int "json:\"last_page\""
			Links       []struct {
				URL    interface{} "json:\"url\""
				Label  string      "json:\"label\""
				Active bool        "json:\"active\""
			} "json:\"links\""
			Path    string "json:\"path\""
			PerPage int    "json:\"per_page\""
			To      int    "json:\"to\""
			Total   int    "json:\"total\""
		}{
			CurrentPage: 1,
			From:        1,
			LastPage:    1,
			Links: []struct {
				URL    interface{} "json:\"url\""
				Label  string      "json:\"label\""
				Active bool        "json:\"active\""
			}{
				{
					Label:  "&laquo; Previous",
					Active: false,
				},
				{
					URL:    "https://api.ryzm.jp/public/lives?page=1",
					Label:  "1",
					Active: true,
				},
				{
					Label:  "Next &raquo;",
					Active: false,
				},
			},
			Path:    "https://api.ryzm.jp/public/lives",
			PerPage: 12,
			To:      8,
			Total:   8,
		},
	}

	s := NewResponseSerializer()

	result, err := s.SelializeRyzmData(input)

	if err != nil {
		t.Fatal(err)
	}

	expectedDate, _ := time.Parse(time.RFC3339, "2022-11-17")
	expected := []model.Event{
		{
			UUID:       "6c16f274-aa55-4ea0-9994-64020f25629d",
			Name:       "【単独無銭】プリズムセン",
			Date:       &expectedDate,
			ArtistName: "PRSMIN",
			VenueName:  "Shibuya Milkyway",
			TicketURL:  "https://t.livepocket.jp/e/g6fq1",
		},
	}

	diff := cmp.Diff(expected, result)

	if diff != "" {
		t.Errorf("RyzmResponseSelializer differs: (-got +want)\n%s", diff)
	}

}
