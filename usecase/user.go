package usecase

type User interface {
	GetUser(id int) (*User, error)
	CreateUser(user *User) (*User, error)
	UpdateUser(user *User) (*User, error)
	DeleteUser(id int) (*User, error)
	CreateAPIKey(id int) (*User, error)
	DeleteAPIKey(id int) (*User, error)
}
