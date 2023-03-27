package translator_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/scarlet0725/prism-api/ent"
	"github.com/scarlet0725/prism-api/infrastructure/translator"
	"github.com/scarlet0725/prism-api/model"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestTranslator(t *testing.T) {
	var (
	//EntUserType = reflect.TypeOf(ent.User{})
	)
	tests := []struct {
		name  string
		input interface{}
		want  interface{}
	}{
		{
			name: "Ent Schema To Model: User",
			input: ent.User{
				ID:        1,
				UserID:    "test",
				Username:  "test",
				FirstName: "frist",
				LastName:  "last",
				Email:     "example@example.com",
				Password:  []byte("test"),
				APIKey:    "test",
				CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				DeletedAt: nil},
			want: model.User{
				ID:         1,
				UserID:     "test",
				Username:   "test",
				FamilyName: "frist",
				GivenName:  "last",
				Email:      "example@example.com",
				Password:   []byte("test"),
				APIKey:     "test",
				CreatedAt:  time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt:  time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				DeletedAt:  nil,
			},
		},
		{
			name: "Ent Schema To Model: ExternalCalendar",
			input: ent.ExternalCalendar{
				ID:          1,
				Name:        "test",
				Description: "test",
				CalendarID:  "testcalendar@googlegroup.com",
				SourceType:  "google",
				UserID:      1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				result interface{}
				diff   string
			)

			//	v := reflect.ValueOf(tt.input).Type()

			switch value := tt.input.(type) {
			case ent.User:
				result = *translator.UserFromEnt(&value)

			default:
				t.Errorf("TestTranslator() Input Type Error: %v", reflect.ValueOf(tt.input).Type().Name())
			}

			if diff = cmp.Diff(tt.want, result); diff != "" {
				t.Errorf("TestTranslator() mismatch (-want +got):\n%s", diff)
			}
		},
		)
	}
}
