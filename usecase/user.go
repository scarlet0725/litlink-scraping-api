package usecase

import (
	"context"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"net/http"

	"github.com/scarlet0725/prism-api/framework"
	"github.com/scarlet0725/prism-api/infrastructure/repository"
	"github.com/scarlet0725/prism-api/model"
	"golang.org/x/crypto/bcrypt"
)

type User interface {
	GetUserByUserID(ctx context.Context, id string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	DeleteUser(context.Context, *model.User) error
	GetUserByAPIKey(ctx context.Context, apiKey string) (*model.User, error)
	VerifyAccount(ctx context.Context, userID string) (*model.User, error)
	CreateCalendar(context.Context, *model.ExternalCalendar) (*model.ExternalCalendar, error)
	RegistrationEvent(context.Context, *model.User, *model.Event) (*model.RegistrationEventResponse, error)
	CreateAPIKey(context.Context, *model.LoginRequest) (*model.CreateAPIKeyResponse, error)
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

func (a *userUsecase) GetUser(ctx context.Context, id string) (*model.User, *model.AppError) {
	user, err := a.user.GetUser(ctx, id)

	if err != nil {
		return nil, &model.AppError{
			Code: 404,
			Msg:  "user_not_found",
		}
	}

	return user, nil
}

func (a *userUsecase) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	id := a.random.Generate(userIDLength)

	user.UserID = id

	user, err := a.user.CreateUser(ctx, user)

	if err != nil {
		return nil, &model.AppError{
			Code: 400,
			Msg:  "Failed to create user",
		}
	}

	return user, nil

}

func (a *userUsecase) DeleteUser(ctx context.Context, user *model.User) error {
	if user.ID == 0 {
		return &model.AppError{
			Code: 404,
			Msg:  "User not found",
		}
	}
	err := a.user.DeleteUser(ctx, user)
	return err
}

func (a *userUsecase) GetUserByAPIKey(ctx context.Context, key string) (*model.User, error) {
	sha512 := sha512.Sum512([]byte(key))
	k := hex.EncodeToString(sha512[:])
	user, err := a.user.GetUserByAPIKey(ctx, string(k))

	if err != nil {
		return nil, &model.AppError{
			Code: 404,
			Msg:  "user_not_found",
		}
	}

	return user, nil

}

func (a *userUsecase) VerifyAccount(ctx context.Context, userID string) (*model.User, error) {
	user, err := a.user.GetUser(ctx, userID)
	if err != nil {
		return nil, &model.AppError{
			Code: 404,
			Msg:  "User not found",
		}
	}

	user.IsAdminVerified = true
	user, err = a.user.UpdateUser(ctx, user)

	if err != nil {
		return nil, &model.AppError{
			Code: 400,
			Msg:  "Failed to update user",
		}
	}
	return user, nil
}

func (a *userUsecase) GetUserByUserID(ctx context.Context, userID string) (*model.User, error) {
	user, err := a.user.GetUser(ctx, userID)

	if err != nil {
		return nil, &model.AppError{
			Code: 400,
			Msg:  "Failed to update user",
		}
	}

	return user, nil
}

func (a *userUsecase) CreateCalendar(ctx context.Context, req *model.ExternalCalendar) (*model.ExternalCalendar, error) {

	_, err := a.user.GetUserCalendarByUserID(ctx, req.UserID)

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
	token, err := a.user.GetGoogleOAuthToken(ctx, req.UserID)

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

	a.user.SaveExternalCalendar(ctx, result)

	return result, err
}

func (a *userUsecase) RegistrationEvent(ctx context.Context, user *model.User, event *model.Event) (*model.RegistrationEventResponse, error) {

	userID := int(user.ID)
	err := a.user.AddRegistrationEvent(ctx, user, event)

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

	config, err := a.user.GetGoogleCalendarConfig(ctx, userID)

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

func (u *userUsecase) CreateAPIKey(ctx context.Context, req *model.LoginRequest) (*model.CreateAPIKeyResponse, error) {

	user, err := u.user.GetUserByUsername(ctx, req.Username)

	if err != nil {
		return nil, &model.AppError{
			Code: 404,
			Msg:  "User not found",
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		return nil, &model.AppError{
			Code: 400,
			Msg:  "Invalid password",
		}
	}

	key, err := u.random.GenerateUUID4()

	if err != nil {
		return nil, &model.AppError{
			Code: 500,
			Msg:  "Failed to generate api key",
		}
	}

	hash := sha512.Sum512([]byte(key))
	user.APIKey = hex.EncodeToString(hash[:])
	user, err = u.user.UpdateUser(ctx, user)

	result := &model.CreateAPIKeyResponse{
		APIKey: key,
		User:   user,
	}
	return result, err

}
