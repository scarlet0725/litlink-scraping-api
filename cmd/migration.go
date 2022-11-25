package cmd

import (
	"github.com/scarlet0725/prism-api/model"
	"github.com/scarlet0725/prism-api/schema"
	"gorm.io/gorm"
)

func MigrationDB(db *gorm.DB) {

	db.Debug().AutoMigrate(&schema.User{}, &schema.Artist{}, &schema.Event{}, &schema.Venue{})

	user := &schema.User{
		User: model.User{
			Username:        "Admin",
			UserID:          "7CAMWFMHF4YXS23P",
			Email:           "example@example.com",
			Password:        []byte("$2a$10$h2qZT.YdMDWYRZV36LkUQO8A6sB.coL8mzkMl25VA2eOAGfJTcGZ2"),
			APIKey:          "776f9e9c12eeffad240a488d12d5c8276c947ec3d67dcce5520be08580755f8edff66e5b502f27e7c400f5b96927e478426f44ee823b484951ba789e4ed1e070",
			IsAdminVerified: true,
			DeleteProtected: true,
		},
	}

	db.Create(&user)

	artists := []schema.Artist{
		{
			Artist: model.Artist{
				ArtistID:       "2MERWD724422E6D8",
				Name:           "prsmin",
				URL:            "https://prsmin.com",
				RyzmHost:       "prsmin.com",
				CrawlTargetURL: "https://api.ryzm.jp/public/lives",
				CrawlSiteType:  "ryzm",
			},
		},
		{
			Artist: model.Artist{
				ArtistID:       "7MHK8G565KEFQERZ",
				Name:           "onthetreatsuperseason",
				URL:            "https://onthetreatsuperseason.com",
				RyzmHost:       "onthetreatsuperseason.com",
				CrawlTargetURL: "https://api.ryzm.jp/public/lives",
				CrawlSiteType:  "ryzm",
			},
		},
	}

	db.Create(&artists)
}
