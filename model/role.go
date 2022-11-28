package model

type Role struct {
	ID          uint    `json:"-" gorm:"primary_key;unique;not null;auto_increment"`
	RoleID      string  `json:"role_id" gorm:"unique;not null"`
	Name        string  `json:"name" gorm:"unique;not null"`
	Description string  `json:"description"`
	Users       []*User `json:"-" gorm:"many2many:user_roles"`
}
