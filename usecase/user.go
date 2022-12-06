package usecase

import (
	"crypto/sha512"
	"encoding/hex"

	"github.com/scarlet0725/prism-api/framework"
	"github.com/scarlet0725/prism-api/infrastructure/repository"
	"github.com/scarlet0725/prism-api/model"
)

type User interface {
	GetUserByUserID(id string) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	DeleteUser(*model.User) error
	GetUserByAPIKey(apiKey string) (*model.User, error)
	VerifyAccount(userID string) (*model.User, error)
}

type userUsecase struct {
	db     repository.DB
	random framework.RandomID
}

func NewUserUsecase(db repository.DB, r framework.RandomID) User {
	return &userUsecase{
		db:     db,
		random: r,
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
	id := a.random.Generate(userIDLength)

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

func (a *userUsecase) VerifyAccount(userID string) (*model.User, error) {
	user, err := a.db.GetUser(userID)
	if err != nil {
		return nil, &model.AppError{
			Code: 404,
			Msg:  "User not found",
		}
	}

	user.IsAdminVerified = true
	user, err = a.db.UpdateUser(user)

	if err != nil {
		return nil, &model.AppError{
			Code: 400,
			Msg:  "Failed to update user",
		}
	}
	return user, nil
}

func (a *userUsecase) GetUserByUserID(userID string) (*model.User, error) {
	user, err := a.db.GetUser(userID)

	if err != nil {
		return nil, &model.AppError{
			Code: 400,
			Msg:  "Failed to update user",
		}
	}

	return user, nil
}
