package controller

import (
	"github.com/scarlet0725/prism-api/model"
	"github.com/scarlet0725/prism-api/repository"
)

type FetchController interface {
	Fetch(*model.ScrapingRequest) (*model.ScrapingResult, error)
}

type fetchController struct {
	cache repository.CacheRepository
	http  repository.HTTPRepository
}

func NewFetchController(http repository.HTTPRepository, cache repository.CacheRepository) FetchController {
	return &fetchController{
		http:  http,
		cache: cache,
	}
}

func (f *fetchController) Fetch(r *model.ScrapingRequest) (*model.ScrapingResult, error) {
	var (
		cache *model.CacheData
		ok    error
	)

	cache, ok = f.cache.GetByKey(r.URL)

	switch ok {
	case nil:
		result := &model.ScrapingResult{
			Data:    cache.Value,
			Request: r,
		}
		return result, nil

	default:
		b, err := f.http.Get(r.URL, r.Option.HTTPHeader, r.Option.HTTPParams)
		if err != nil {
			return &model.ScrapingResult{}, err
		}
		cacheData := model.CacheData{
			Key:   r.URL,
			Value: b,
		}

		f.cache.Set(&cacheData, 600)
		//TODO: キャッシュを書き込み失敗したらロギングする

		result := &model.ScrapingResult{
			Request: r,
			Data:    b,
		}

		return result, nil
	}
}
