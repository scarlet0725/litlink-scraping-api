package model

import "time"

type User struct {
	ID               uint   `json:"-" gorm:"primary_key;unique;not null;auto_increment"`
	UserID           string `json:"user_id" gorm:"unique;not null"`
	Username         string `json:"username" gorm:"unique;not null"`
	FamilyName       string `json:"family_name"`
	GivenName        string `json:"given_name"`
	Email            string `json:"email" gorm:"unique;not null"`
	Password         []byte `json:"-" gorm:"not null"`
	APIKey           string `json:"-"`
	IsAdminVerified  bool   `json:"-" gorm:"not null"`
	DeleteProtected  bool   `json:"-" gorm:"not null"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time        `gorm:"index" json:"-"`
	Events           []*Event          `json:"-" gorm:"many2many:user_events"`
	Roles            []*Role           `json:"-" gorm:"many2many:user_roles"`
	GoogleToken      *GoogleOAuthToken `json:"-"`
	GoogleOAuthState *GoogleOAuthState `json:"-"`
	ExternalCalendar *ExternalCalendar `json:"-"`
}

type GoogleOAuthToken struct {
	ID           uint   `json:"-" gorm:"primary_key;unique;not null;auto_increment"`
	UserID       *uint  `json:"user_id" gorm:"unique;not null"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	Expiry       time.Time
}

func (g GoogleOAuthToken) TableName() string {
	return "google_oauth_tokens"
}

type OAuthURLResponse struct {
	AuthURL string `json:"auth_url"`
	State   string `json:"state"`
}

type GoogleOAuthState struct {
	ID     uint   `gorm:"primary_key;unique;not null;auto_increment"`
	UserID *uint  `gorm:"unique;not null"`
	State  string `gorm:"unique;not null"`
}

func (g GoogleOAuthState) TableName() string {
	return "google_oauth_states"
}
