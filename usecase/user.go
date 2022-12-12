package usecase

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"net/http"

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
	CreateCalendar(*model.ExternalCalendar) (*model.ExternalCalendar, error)
	RegistrationEvent(*model.User, *model.Event) (*model.RegistrationEventResponse, error)
}

type userUsecase struct {
	user       repository.User
	random     framework.RandomID
	google     framework.GoogleOAuth
	srvBuilder func(*http.Client) repository.ExternalCalendar
}

func NewUserUsecase(u repository.User, r framework.RandomID, g framework.GoogleOAuth, f func(*http.Client) repository.ExternalCalendar) User {
	return &userUsecase{
		user:       u,
		random:     r,
		google:     g,
		srvBuilder: f,
	}
}

func (a *userUsecase) GetUser(id string) (*model.User, *model.AppError) {
	user, err := a.user.GetUser(id)

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

	user, err := a.user.CreateUser(user)

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
	err := a.user.DeleteUser(user)
	return err
}

func (a *userUsecase) GetUserByAPIKey(key string) (*model.User, error) {
	sha512 := sha512.Sum512([]byte(key))
	k := hex.EncodeToString(sha512[:])
	user, err := a.user.GetUserByAPIKey(string(k))

	if err != nil {
		return nil, &model.AppError{
			Code: 404,
			Msg:  "user_not_found",
		}
	}

	return user, nil

}

func (a *userUsecase) VerifyAccount(userID string) (*model.User, error) {
	user, err := a.user.GetUser(userID)
	if err != nil {
		return nil, &model.AppError{
			Code: 404,
			Msg:  "User not found",
		}
	}

	user.IsAdminVerified = true
	user, err = a.user.UpdateUser(user)

	if err != nil {
		return nil, &model.AppError{
			Code: 400,
			Msg:  "Failed to update user",
		}
	}
	return user, nil
}

func (a *userUsecase) GetUserByUserID(userID string) (*model.User, error) {
	user, err := a.user.GetUser(userID)

	if err != nil {
		return nil, &model.AppError{
			Code: 400,
			Msg:  "Failed to update user",
		}
	}

	return user, nil
}

func (a *userUsecase) CreateCalendar(req *model.ExternalCalendar) (*model.ExternalCalendar, error) {

	_, err := a.user.GetUserCalendarByUserID(req.UserID)

	//エラーがNilの場合はユーザーのカレンダーが存在するのでエラーを返す
	if err == nil {
		return nil, &model.AppError{
			Code: 400,
			Msg:  "User already has a calendar",
		}
	}

	if !errors.Is(err, repository.ErrNotFound) {
		return nil, &model.AppError{
			Code: 500,
			Msg:  "Failed to get user calendar",
		}
	}

	//レコードがないエラーが発生した場合は、ユーザーのカレンダーが存在しないので、カレンダーを作成する
	token, err := a.user.GetGoogleOAuthToken(req.UserID)

	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, &model.AppError{
				Code: 400,
				Msg:  "Link your google account first",
			}
		}
		return nil, &model.AppError{
			Code: 500,
			Msg:  "Failed to get google oauth token",
		}
	}

	client := a.google.GetClient(token)

	srv := a.srvBuilder(client)

	result, err := srv.CreateCalendar(req)

	if err != nil {
		return nil, &model.AppError{
			Code: 500,
			Msg:  "Failed to create calendar",
		}
	}

	a.user.SaveExternalCalendar(result)

	return result, err
}

func (a *userUsecase) RegistrationEvent(user *model.User, event *model.Event) (*model.RegistrationEventResponse, error) {

	userID := int(user.ID)
	err := a.user.AddRegistrationEvent(user, event)

	if err != nil {
		return nil, &model.AppError{
			Code: 400,
			Msg:  "Failed to update user",
		}
	}

	result := &model.RegistrationEventResponse{
		CalenderAdded:   false,
		EventRegistered: true,
		EventID:         event.EventID,
	}

	config, err := a.user.GetGoogleCalendarConfig(userID)

	if err != nil {
		return result, &model.AppError{
			Code: 404,
			Msg:  "Google calendar config not found",
		}
	}

	client := a.google.GetClient(&config.GoogleOAuthToken)

	srv := a.srvBuilder(client)

	calendarEvent := &model.CalendarEvent{
		ExternalEventID:    event.EventID,
		ExternalCalendarID: config.ExternalCalendar.CalendarID,
		Event:              event,
		Public:             false,
	}

	_, err = srv.CreateEvent(calendarEvent)

	if err != nil {
		return result, nil
	}

	result.CalenderAdded = true

	return result, err

}
