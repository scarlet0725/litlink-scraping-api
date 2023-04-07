package model

import "time"

type User struct {
	ID               int               `json:"-"`
	UserID           string            `json:"user_id"`
	Username         string            `json:"username"`
	FamilyName       string            `json:"family_name"`
	GivenName        string            `json:"given_name"`
	Email            string            `json:"email"`
	Password         []byte            `json:"-"`
	APIKey           string            `json:"-"`
	IsAdminVerified  bool              `json:"-"`
	DeleteProtected  bool              `json:"-"`
	CreatedAt        time.Time         `json:"-"`
	UpdatedAt        time.Time         `json:"-"`
	DeletedAt        *time.Time        `json:"-"`
	Events           []*Event          `json:"-"`
	Roles            []*Role           `json:"-"`
	GoogleToken      *GoogleOAuthToken `json:"-"`
	GoogleOAuthState *GoogleOAuthState `json:"-"`
	ExternalCalendar *ExternalCalendar `json:"-"`
}

func (u *User) Verify() {
	u.IsAdminVerified = true
}

func (u *User) IsVerified() bool {
	return u.IsAdminVerified
}

func (u *User) Protect() {
	u.DeleteProtected = true
}

func (u *User) IsProtected() bool {
	return u.DeleteProtected
}

type GoogleOAuthToken struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	Expiry       time.Time
	User         *User `json:"-"`
}

type OAuthURLResponse struct {
	AuthURL string `json:"auth_url"`
	State   string `json:"state"`
}

type GoogleOAuthState struct {
	ID     int
	UserID int
	State  string
}
