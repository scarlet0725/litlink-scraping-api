package cache

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/scarlet0725/prism-api/model"
)

type Cache interface {
	Set(*model.CacheData, int64) error
	Get(*model.CacheData) (*model.CacheData, error)
	GetByKey(string) (*model.CacheData, error)
}

type redisCache struct {
	Client *redis.Client
}

func CreateRedisManager(c *redis.Client) Cache {
	return &redisCache{
		Client: c,
	}
}

func (c *redisCache) Set(d *model.CacheData, ttl int64) error {
	cmd := c.Client.Set(d.Key, d.Value, time.Duration(ttl)*time.Second)
	err := cmd.Err()
	return err
}

func (c *redisCache) Get(d *model.CacheData) (*model.CacheData, error) {
	cmd := c.Client.Get(d.Key)
	err := cmd.Err()
	if err != nil {
		return nil, err
	}
	data, err := cmd.Bytes()

	if err != nil {
		return nil, err
	}

	return &model.CacheData{
		Key:   d.Key,
		Value: data,
	}, nil

}

func (c *redisCache) GetByKey(key string) (*model.CacheData, error) {
	cmd := c.Client.Get(key)
	err := cmd.Err()
	if err != nil {
		return nil, err
	}
	data, err := cmd.Bytes()

	if err != nil {
		return nil, err
	}

	return &model.CacheData{
		Key:   key,
		Value: data,
	}, nil
}
