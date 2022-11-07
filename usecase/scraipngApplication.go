package usecase

import (
	"github.com/scarlet0725/prism-api/model"
)

type ScrapingApplication interface {
	Execute(*model.ScrapingRequest) (model.ScrapingResult, error)
}
