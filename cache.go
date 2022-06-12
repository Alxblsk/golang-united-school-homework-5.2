package cache

import (
	"time"
)

type Cache struct {
	expirations map[string]time.Time
	data        map[string]string
}

func NewCache() Cache {
	return Cache{
		expirations: map[string]time.Time{},
		data:        map[string]string{},
	}
}

func (c *Cache) Get(key string) (string, bool) {
	c.cleanup()

	val, ok := c.data[key]
	return val, ok
}

func (c *Cache) Put(key string, value string) {
	c.cleanup()

	c.data[key] = value
}

func (c *Cache) Keys() []string {
	c.cleanup()

	keysList := []string{}

	for key, _ := range c.data {
		keysList = append(keysList, key)
	}

	return keysList
}

func (c *Cache) PutTill(key string, value string, deadline time.Time) {
	c.cleanup()
	c.Put(key, value)

	c.expirations[key] = deadline
}

func (c *Cache) cleanup() {
	for key, _ := range c.expirations {
		c.removeExpired(key)
	}
}

func (c *Cache) removeExpired(key string) {
	if deadline, ok := c.expirations[key]; ok {
		timeDiff := deadline.Sub(time.Now())
		if timeDiff <= 0 {
			delete(c.data, key)
			delete(c.expirations, key)
		}
	}
}
