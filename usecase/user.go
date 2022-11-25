package usecase

import (
	"crypto/sha512"
	"encoding/hex"

	"github.com/scarlet0725/prism-api/cmd"
	"github.com/scarlet0725/prism-api/model"
	"github.com/scarlet0725/prism-api/repository"
)

type User interface {
	//GetUser(id string) (*model.User, *model.AppError)
	CreateUser(user *model.User) (*model.User, error)
	//UpdateUser(user *model.User) (*model.User, *model.AppError)
	DeleteUser(*model.User) error
	//CreateAPIKey(id int) (*model.User, *model.AppError)
	//DeleteAPIKey(id int) (*model.User, *model.AppError)
	GetUserByAPIKey(apiKey string) (*model.User, error)
}

type userUsecase struct {
	db repository.DB
}

func NewUserApplication(db repository.DB) User {
	return &userUsecase{
		db: db,
	}
}

func (a *userUsecase) GetUser(id string) (*model.User, *model.AppError) {
	user, err := a.db.GetUser(id)

	if err != nil {
		return nil, &model.AppError{
			Code: 404,
			Msg:  "user_not_found",
		}
	}

	return user, nil
}

func (a *userUsecase) CreateUser(user *model.User) (*model.User, error) {
	id := cmd.MakeRamdomID(userIDLength)

	user.UserID = id

	user, err := a.db.CreateUser(user)

	if err != nil {
		return nil, &model.AppError{
			Code: 400,
			Msg:  "Failed to create user",
		}
	}

	return user, nil

}

func (a *userUsecase) DeleteUser(user *model.User) error {
	if user.ID == 0 {
		return &model.AppError{
			Code: 404,
			Msg:  "User not found",
		}
	}
	err := a.db.DeleteUser(user)
	return err
}

func (a *userUsecase) GetUserByAPIKey(key string) (*model.User, error) {
	sha512 := sha512.Sum512([]byte(key))
	k := hex.EncodeToString(sha512[:])
	user, err := a.db.GetUserByAPIKey(string(k))

	if err != nil {
		return nil, &model.AppError{
			Code: 404,
			Msg:  "user_not_found",
		}
	}

	return user, nil

}
