package infrastructure

import (
	"context"
	"errors"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/scarlet0725/prism-api/ent"
	"github.com/scarlet0725/prism-api/ent/externalcalendar"
	"github.com/scarlet0725/prism-api/ent/googleoauthtoken"
	"github.com/scarlet0725/prism-api/ent/user"
	"github.com/scarlet0725/prism-api/infrastructure/repository"
	"github.com/scarlet0725/prism-api/infrastructure/translator"
	"github.com/scarlet0725/prism-api/model"
)

type userRepository struct {
	db *ent.Client
}

func NewUserRepository(db *ent.Client) repository.User {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) GetUser(ctx context.Context, id string) (*model.User, error) {
	result, err := u.db.User.Query().Where(
		user.And(
			user.UserID(id),
			user.DeletedAtIsNil(),
		),
	).First(ctx)

	if ent.IsNotFound(err) {
		return nil, errors.New("user not found")
	}

	if err != nil {
		return nil, err
	}

	user := translator.UserFromEnt(result)

	return user, nil

}

func (u *userRepository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	result, err := u.db.User.Create().
		SetUserID(user.UserID).
		SetUsername(user.Username).
		SetFirstName(user.FamilyName).
		SetLastName(user.GivenName).
		SetEmail(user.Email).
		SetPassword(user.Password).
		SetAPIKey(user.APIKey).
		SetIsAdminVerified(user.IsAdminVerified).
		SetDeleteProtected(user.DeleteProtected).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	createdUser := translator.UserFromEnt(result)

	return createdUser, nil

}

func (u *userRepository) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	result, err := u.db.User.UpdateOneID(int(user.ID)).
		SetUsername(user.Username).
		SetFirstName(user.FamilyName).
		SetLastName(user.GivenName).
		SetEmail(user.Email).
		SetIsAdminVerified(user.IsAdminVerified).
		SetDeleteProtected(user.DeleteProtected).
		SetAPIKey(user.APIKey).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	updatedUser := translator.UserFromEnt(result)

	return updatedUser, nil

}

func (u *userRepository) DeleteUser(ctx context.Context, deleteUser *model.User) error {
	if deleteUser.DeleteProtected {
		return errors.New("user is protected")
	}

	tx, err := u.db.Tx(ctx)

	if err != nil {
		return err
	}

	c, err := tx.GoogleOauthToken.Delete().
		Where(
			googleoauthtoken.HasUserWith(
				user.ID(int(deleteUser.ID)),
			),
		).Exec(ctx)

	if err != nil || c >= 2 {
		tx.Rollback()
	}

	err = tx.User.UpdateOneID(int(deleteUser.ID)).
		ClearEvents().
		ClearExternalCalendars().
		SetDeletedAt(time.Now()).
		Exec(ctx)

	if err != nil {
		return err
	}

	err = tx.Commit()

	if err != nil {
		return err
	}

	//TODO: 論理削除か物理削除か決める

	return nil
}

func (u *userRepository) GetUserByAPIKey(ctx context.Context, apiKey string) (*model.User, error) {
	result, err := u.db.User.Query().Where(user.APIKey(apiKey)).First(ctx)

	if err != nil {
		return nil, err
	}

	user := translator.UserFromEnt(result)

	return user, nil
}

func (u *userRepository) GetUserCalendarByUserID(ctx context.Context, id int) (*model.ExternalCalendar, error) {
	result, err := u.db.ExternalCalendar.Query().Where(externalcalendar.HasUserWith(user.ID(id))).First(ctx)

	if err != nil {
		return nil, err
	}

	calendar := translator.ExternalCalendarFromEnt(result)

	return calendar, nil

}

func (u *userRepository) GetGoogleCalendarConfig(ctx context.Context, id int) (*model.GoogleCalenderConfig, error) {
	var config *model.GoogleCalenderConfig

	err := u.db.GoogleOauthToken.Query().Modify(
		func(s *sql.Selector) {
			t := sql.Table("external_calendars")
			s.FullJoin(t).
				On(
					t.C("user_id"),
					s.C("user_id"),
				).
				Where(
					sql.EQ(
						t.C("source_type"),
						"google"),
				)
		},
	).
		Scan(ctx, &config)

	if err != nil {
		return nil, err
	}

	return config, nil

}

func (u *userRepository) GetGoogleOAuthToken(ctx context.Context, id int) (*model.GoogleOAuthToken, error) {
	result, err := u.db.GoogleOauthToken.Query().Where(googleoauthtoken.HasUserWith(user.ID(id))).First(ctx)

	if err != nil {
		return nil, err
	}

	token := translator.GoogleOAuthTokenFromEnt(result)

	return token, nil

}

func (u *userRepository) SaveExternalCalendar(ctx context.Context, cal *model.ExternalCalendar) (*model.ExternalCalendar, error) {
	result, err := u.db.ExternalCalendar.Create().
		SetName(cal.Name).
		SetCalendarID(cal.CalendarID).
		SetSourceType(cal.Type).
		SetDescription(cal.Description).
		SetUserID(cal.UserID).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	createdCal := translator.ExternalCalendarFromEnt(result)

	return createdCal, nil
}

func (u *userRepository) AddRegistrationEvent(ctx context.Context, user *model.User, event *model.Event) error {
	_, err := u.db.User.UpdateOneID(int(user.ID)).AddEvents(translator.EventFromModel(event)).Save(ctx)

	return err
}

func (u *userRepository) RemoveRegistrationEvent(ctx context.Context, user *model.User, event *model.Event) error {
	e := translator.EventFromModel(event)
	_, err := u.db.User.UpdateOneID(int(user.ID)).RemoveEvents(e).Save(ctx)
	return err
}

func (u *userRepository) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	result, err := u.db.User.Query().Where(user.Username(username)).First(ctx)

	if err != nil {
		return nil, err
	}

	user := translator.UserFromEnt(result)

	return user, nil
}

func (u *userRepository) VerifyUser(ctx context.Context, ID int) error {
	err := u.db.User.UpdateOneID(ID).SetIsAdminVerified(true).Exec(ctx)

	return err
}
