package model

import "time"

type User struct {
	ID              uint   `json:"-" gorm:"primary_key;unique;not null;auto_increment"`
	UserID          string `json:"user_id" gorm:"unique;not null"`
	Username        string `json:"username" gorm:"unique;not null"`
	FamilyName      string `json:"family_name"`
	GivenName       string `json:"given_name"`
	Email           string `json:"email" gorm:"unique;not null"`
	Password        []byte `json:"-" gorm:"not null"`
	APIKey          string `json:"-"`
	IsAdminVerified bool   `json:"-" gorm:"not null"`
	DeleteProtected bool   `json:"-" gorm:"not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time `sql:"index" json:"-"`
}
