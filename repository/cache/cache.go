package cache

import (
	"blogging-platform-api/domain/dto"
	log "blogging-platform-api/logger"
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
)

type (
	ICache interface {
		Set(data *dto.Post) error
		Get() (*dto.Post, error)
	}

	Cache struct {
		redis *redis.Client
		ttl   time.Duration
	}
)

func NewCache(redis *redis.Client, ttl time.Duration) ICache {
	return &Cache{
		redis: redis,
		ttl:   ttl,
	}
}

func (c *Cache) getKey() string {
	return "<cache key>"
}

func (c *Cache) Set(data *dto.Post) error {
	val, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := c.redis.Set(c.getKey(), val, c.ttl).Err(); err != nil {
		return err
	}

	log.Info("cache created", "key", c.getKey(), "ttl", c.ttl)

	return nil
}

func (c *Cache) Get() (*dto.Post, error) {
	response := new(dto.Post)

	val, err := c.redis.Get(c.getKey()).Bytes()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(val, &response)
	if err != nil {
		return nil, err
	}

	log.Info("load data from cache")

	return response, nil
}
