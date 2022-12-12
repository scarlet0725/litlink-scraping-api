package usecase_test

import (
	"testing"

	"github.com/scarlet0725/prism-api/framework"
	"github.com/scarlet0725/prism-api/infrastructure"
	"github.com/scarlet0725/prism-api/model"
	"github.com/scarlet0725/prism-api/schema"
	"github.com/scarlet0725/prism-api/usecase"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestArtist(t *testing.T) {
	sqlite := sqlite.Open("file::memory:?cache=shared")

	db, err := gorm.Open(sqlite, &gorm.Config{})

	db.AutoMigrate(
		&schema.User{},
		&schema.Artist{},
		&schema.Event{},
		&schema.Venue{},
		&schema.RyzmEvent{},
		&schema.UnStructuredEventInformation{},
		&schema.Role{},
		&model.GoogleOAuthState{},
		&model.GoogleOAuthToken{},
	)

	if err != nil {
		t.Fatal(err)
	}

	gorm := infrastructure.NewGORMClient(db)

	random := framework.NewRamdomIDGenerator()

	usecase := usecase.NewArtistUsecase(gorm, random)

	t.Run(
		"Artistを作成できる",
		func(t *testing.T) {
			artist := &model.Artist{
				Name: "test",
			}
			_, err := usecase.CreateArtist(artist)

			if err != nil {
				t.Errorf("error: %v", err)
			}
		},
	)

	t.Run(
		"存在するArtistIDを元にArtistを取得できる",
		func(t *testing.T) {
			artist := &model.Artist{
				Name: "test",
			}
			createdArtist, err := usecase.CreateArtist(artist)

			if err != nil {
				t.Errorf("error: %v", err)
			}

			_, err = usecase.GetArtistByID(createdArtist.ArtistID)

			if err != nil {
				t.Errorf("error: %v", err)
			}
		},
	)

	t.Run(
		"存在しないArtistIDを元にArtistを取得できない",
		func(t *testing.T) {
			_, err := usecase.GetArtistByID("NotExistsArtistID")

			if err == nil {
				t.Errorf("error: %v", err)
			}
		},
	)

}
