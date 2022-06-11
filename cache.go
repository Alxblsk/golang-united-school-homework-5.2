package cache

import (
	"time"
)

type Cache struct {
	data map[string]string
}

func NewCache() Cache {
	return Cache{
		data: map[string]string{},
	}
}

func (c *Cache) Get(key string) (string, bool) {
	val, ok := c.data[key]
	return val, ok
}

func (c *Cache) Put(key string, value string) {
	c.data[key] = value
}

func (c *Cache) Keys() []string {
	keysList := make([]string, len(c.data))

	for _, val := range c.data {
		keysList = append(keysList, val)
	}

	return keysList
}

func (c *Cache) PutTill(key string, value string, deadline time.Time) {
	c.data[key] = value

	timeDiff := deadline.Sub(time.Now())

	// Should be a cleanup before every other operation instead; but let's start with a sync method
	time.Sleep(timeDiff)
	delete(c.data, key)
}
