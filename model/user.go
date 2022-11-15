package model

type User struct {
	ID              int    `json:"id"`
	Username        string `json:"username"`
	FamilyName      string `json:"family_name"`
	GivenName       string `json:"given_name"`
	Email           string `json:"email"`
	Password        []byte `json:"-"`
	APIKey          []byte `json:"-"`
	IsAdminVerified bool   `json:"-"`
}
