package cache

import (
	"blogging-platform-api/domain"
	log "blogging-platform-api/logger"
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
)

type (
	ICache interface {
		Set(data *domain.StructName) error
		Get() (*domain.StructName, error)
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

func (c *Cache) Set(data *domain.StructName) error {
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

func (c *Cache) Get() (*domain.StructName, error) {
	response := new(domain.StructName)

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
