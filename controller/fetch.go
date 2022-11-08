package controller

import (
	"github.com/scarlet0725/prism-api/cache"
	"github.com/scarlet0725/prism-api/model"
	"github.com/scarlet0725/prism-api/scraping"
)

type FetchController interface {
	Fetch(model.ScrapingRequest) (model.ScrapingResult, error)
}

type fetchController struct {
	s scraping.Client
	c cache.Cache
}

func NewFetchController(s scraping.Client, c cache.Cache) FetchController {
	return &fetchController{
		s: s,
		c: c,
	}
}

func (f *fetchController) Fetch(r model.ScrapingRequest) (model.ScrapingResult, error) {
	var err error
	var s model.ScrapingResult

	cache, ok := f.c.GetByKey(r.URL)

	switch ok {
	case nil:
		s = model.ScrapingResult{
			Data: cache.Value,
		}
	default:
		s, err = f.s.Execute(r.URL)
		if err != nil {
			break
		}
		cacheData := model.CacheData{
			Key:   r.URL,
			Value: s.Data,
		}

		f.c.Set(&cacheData, 600)

	}

	if err != nil {
		return model.ScrapingResult{}, err
	}

	return s, nil
}
