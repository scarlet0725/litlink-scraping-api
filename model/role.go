package model

type Role struct {
	ID          int     `json:"-"`
	RoleID      string  `json:"role_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Users       []*User `json:"-"`
}
