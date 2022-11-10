package controller

import (
	"github.com/scarlet0725/prism-api/gateway"
	"github.com/scarlet0725/prism-api/model"
)

type FetchController interface {
	Fetch(*model.ScrapingRequest) (model.ScrapingResult, error)
}

type fetchController struct {
	s gateway.Client
	c gateway.Cache
}

func NewFetchController(s gateway.Client, c gateway.Cache) FetchController {
	return &fetchController{
		s: s,
		c: c,
	}
}

func (f *fetchController) Fetch(r *model.ScrapingRequest) (model.ScrapingResult, error) {
	var err error
	var s model.ScrapingResult

	cache, ok := f.c.GetByKey(r.URL)

	switch ok {
	case nil:
		s = model.ScrapingResult{
			Data:    cache.Value,
			Request: r,
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
		//TODO: キャッシュを書き込み失敗したらロギングする

		s.Request = r

	}

	if err != nil {
		return model.ScrapingResult{}, err
	}

	return s, nil
}
