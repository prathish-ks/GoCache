package cache

type Cache interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Delete(key string) error
}

type CacheItem struct {
	Key   string
	Value interface{}
}

type Cache struct {
	items map[string]CacheItem
}

func NewCache() *Cache {
	return &Cache{
		items: make(map[string]CacheItem),
	}
}

func (c *Cache) Set(key string, value interface{}) error {
	c.items[key] = CacheItem{Key: key, Value: value}
	return nil
}

func (c *Cache) Get(key string) (interface{}, error) {
	item, exists := c.items[key]
	if !exists {
		return nil, nil
	}
	return item.Value, nil
}

func (c *Cache) Delete(key string) error {
	delete(c.items, key)
	return nil
}