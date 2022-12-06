package repository

import (
	"github.com/scarlet0725/prism-api/model"
)

type CacheRepository interface {
	Set(*model.CacheData, int64) error
	Get(*model.CacheData) (*model.CacheData, error)
	GetByKey(string) (*model.CacheData, error)
}
