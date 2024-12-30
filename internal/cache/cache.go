package cache

import "fmt"

type Cache interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Delete(key string) error
}

type CacheItem struct {
	Key   string
	Value interface{}
}

type MemoryCache struct {
	items map[string]CacheItem
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		items: make(map[string]CacheItem),
	}
}

func (c *MemoryCache) Set(key string, value interface{}) error {
	c.items[key] = CacheItem{Key: key, Value: value}
	return nil
}

func (c *MemoryCache) Get(key string) (interface{}, error) {
	item, exists := c.items[key]
	if (!exists) {
		return nil, fmt.Errorf("item not found")
	}
	return item.Value, nil
}

func (c *MemoryCache) Delete(key string) error {
	delete(c.items, key)
	return nil
}