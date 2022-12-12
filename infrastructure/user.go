package infrastructure

import (
	"errors"

	"github.com/scarlet0725/prism-api/infrastructure/repository"
	"github.com/scarlet0725/prism-api/model"
	"github.com/scarlet0725/prism-api/schema"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.User {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) GetUser(id string) (*model.User, error) {
	var user model.User
	err := u.db.Preload(clause.Associations).Where("user_id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) CreateUser(user *model.User) (*model.User, error) {
	var schema schema.User
	schema.User = *user

	result := u.db.Create(&schema)

	err := result.Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) UpdateUser(user *model.User) (*model.User, error) {
	var schema schema.User
	schema.User = *user
	err := u.db.Save(&u).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) DeleteUser(user *model.User) error {
	return u.db.Delete(user).Error
}

func (u *userRepository) GetUserByAPIKey(apiKey string) (*model.User, error) {
	var user model.User
	err := u.db.Preload(clause.Associations).Where("api_key = ?", apiKey).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) GetUserCalendarByUserID(id int) (*model.ExternalCalendar, error) {
	var calendar model.ExternalCalendar
	err := u.db.Where("user_id = ?", id).First(&calendar).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return &calendar, nil
}

func (u *userRepository) GetGoogleCalendarConfig(id int) (*model.GoogleCalenderConfig, error) {
	var config *model.GoogleCalenderConfig

	err := u.db.Model(&model.GoogleOAuthToken{}).Joins("INNER JOIN external_calendar ON google_oauth_tokens.user_id == external_calendar.user_id").Where("user_id = ?", id).First(&config).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	return config, nil

}

func (u *userRepository) GetGoogleOAuthToken(id int) (*model.GoogleOAuthToken, error) {
	token := &model.GoogleOAuthToken{}

	err := u.db.First(token, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	return token, nil
}
