package cache

import (
	"time"

	"github.com/jellydator/ttlcache/v3"
)

type Cache[T any] struct {
	storage *ttlcache.Cache[string, T]
}

func New[T any]() *Cache[T] {
	return &Cache[T]{
		storage: ttlcache.New(ttlcache.WithDisableTouchOnHit[string, T]()),
	}
}

func (c *Cache[T]) Start() {
	c.storage.Start()
}

func (c *Cache[T]) Set(key string, value T, ttl time.Duration) {
	c.storage.Set(key, value, ttl)
}

func (c *Cache[T]) Get(key string) (T, bool) {
	item := c.storage.Get(key)
	if item == nil {
		return *new(T), false
	}
	return item.Value(), true
}

func (c *Cache[T]) Stop() {
	c.storage.Stop()
}
