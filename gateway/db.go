package gateway

import (
	"github.com/scarlet0725/prism-api/model"
	"github.com/scarlet0725/prism-api/repository"
)

type InMemoryDB struct {
}

func (d *InMemoryDB) GetArtistByName(name string) (*model.Artist, error) {
	switch name {
	case "prsmin":
		return &model.Artist{
			Name:           "prsmin",
			URL:            "https://prsmin.com",
			RyzmHost:       "prsmin.com",
			CrawlTargetURL: "https://api.ryzm.jp/public/lives",
			CrawlSiteType:  "ryzm",
		}, nil
	case "onthetreatsuperseason":
		return &model.Artist{
			Name:           "onthetreatsuperseason",
			URL:            "https://onthetreatsuperseason.com",
			RyzmHost:       "onthetreatsuperseason.com",
			CrawlTargetURL: "https://api.ryzm.jp/public/lives",
			CrawlSiteType:  "ryzm",
		}, nil
	default:
		return nil, &model.AppError{
			Code: 404,
			Msg:  "artist_not_found",
		}
	}
}

func NewDevDB() repository.DB {
	return &InMemoryDB{}
}
