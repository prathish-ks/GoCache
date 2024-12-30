package cache

import (
	"sync"
	"time"
)

// MemoryCache is an in-memory implementation of the Cache interface.
type MemoryCache struct {
	data  map[string]cacheItem
	mutex sync.RWMutex
}

type cacheItem struct {
	value      interface{}
	expiration int64
}

// NewMemoryCache creates a new instance of MemoryCache.
func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		data: make(map[string]cacheItem),
	}
}

// Set adds a value to the cache with a specified expiration time.
func (m *MemoryCache) Set(key string, value interface{}, duration time.Duration) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.data[key] = cacheItem{
		value:      value,
		expiration: time.Now().Add(duration).UnixNano(),
	}
}

// Get retrieves a value from the cache. It returns nil if the key does not exist or has expired.
func (m *MemoryCache) Get(key string) interface{} {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	item, found := m.data[key]
	if !found || time.Now().UnixNano() > item.expiration {
		return nil
	}
	return item.value
}

// Delete removes a value from the cache.
func (m *MemoryCache) Delete(key string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	delete(m.data, key)
}