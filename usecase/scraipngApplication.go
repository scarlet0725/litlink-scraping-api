package usecase

import (
	"github.com/scarlet0725/litlink-scraping-api/model"
)

type ScrapingApplication interface {
	Execute(*model.ScrapingRequest) (model.ScrapingResult, error)
}
