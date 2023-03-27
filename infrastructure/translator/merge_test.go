package translator_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/scarlet0725/prism-api/infrastructure/translator"
	"github.com/scarlet0725/prism-api/model"
)

func TestMerge(t *testing.T) {
	baseUser := model.User{
		ID:       1,
		UserID:   "NotChanged",
		Username: "NotChanged",
		Email:    "example@example.com",
		Password: []byte("test"),
		APIKey:   "test",
	}

	tests := []struct {
		name  string
		input interface{}
		want  interface{}
	}{
		{
			name: "Merge: User",
			input: model.User{
				FamilyName: "frist",
				GivenName:  "last",
			},
			want: model.User{
				ID:         1,
				UserID:     "NotChanged",
				Username:   "NotChanged",
				FamilyName: "frist",
				GivenName:  "last",
				Email:      "example@example.com",
				Password:   []byte("test"),
				APIKey:     "test",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result interface{}
			switch value := tt.input.(type) {
			case model.User:
				user := baseUser
				translator.MergeUser(&user, &value)
				result = user
			}

			if diff := cmp.Diff(tt.want, result, cmpopts.IgnoreFields(model.User{}, "CreatedAt", "UpdatedAt", "DeletedAt", "Events", "Roles", "GoogleToken", "GoogleOAuthState", "ExternalCalendar")); diff != "" {
				t.Errorf("Merge() mismatch (-want +got):\n%s", diff)
			}

		})
	}

}
