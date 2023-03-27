package translator

import (
	"time"

	"github.com/scarlet0725/prism-api/ent"
	"github.com/scarlet0725/prism-api/model"
)

func UserFromEnt(source *ent.User) *model.User {

	user := &model.User{
		ID:               source.ID,
		UserID:           source.UserID,
		Username:         source.Username,
		FamilyName:       source.FirstName,
		GivenName:        source.LastName,
		Email:            source.Email,
		Password:         source.Password,
		APIKey:           source.APIKey,
		IsAdminVerified:  source.IsAdminVerified,
		DeleteProtected:  source.DeleteProtected,
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
		DeletedAt:        &time.Time{},
		Events:           []*model.Event{},
		Roles:            []*model.Role{},
		GoogleToken:      &model.GoogleOAuthToken{},
		GoogleOAuthState: &model.GoogleOAuthState{},
		ExternalCalendar: &model.ExternalCalendar{},
	}

	if source.Edges.Roles != nil {
		for _, role := range source.Edges.Roles {
			user.Roles = append(user.Roles, RoleFromEnt(role))
		}
	}

	return user
}

func RoleFromEnt(source *ent.Role) *model.Role {
	if source == nil {
		return nil
	}

	role := &model.Role{
		ID:          source.ID,
		RoleID:      source.RoleID,
		Name:        source.Name,
		Description: source.Description,
	}

	return role
}
