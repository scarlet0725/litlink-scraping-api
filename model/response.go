package model

type UserResponse struct {
	OK   bool   `json:"ok"`
	User *User  `json:"user,omitempty"`
	Err  string `json:"err,omitempty"`
}
