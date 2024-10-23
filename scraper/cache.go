package utils

import (
    "github.com/patrickmn/go-cache"
    "time"
)

type Cache struct {
    cache *cache.Cache
}

func NewCache() *Cache {
    return &Cache{
        cache: cache.New(5*time.Minute, 10*time.Minute),
    }
}

func (c *Cache) Set(key string, value interface{}) {
    c.cache.Set(key, value, cache.DefaultExpiration)
}

func (c *Cache) Get(key string) (interface{}, bool) {
    return c.cache.Get(key)
}