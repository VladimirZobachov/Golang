package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	data      map[string]cacheItem
	mutex     sync.RWMutex
	ttl       time.Duration
	shutdown  chan struct{}
	gcRunning bool
}

type cacheItem struct {
	value  interface{}
	expiry time.Time
}

func NewCache(ttl time.Duration) *Cache {
	cache := &Cache{
		data:     make(map[string]cacheItem),
		mutex:    sync.RWMutex{},
		ttl:      ttl,
		shutdown: make(chan struct{}),
	}
	// Start goroutine for garbage collection
	go cache.startGC()
	return cache
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	item, ok := c.data[key]
	if !ok {
		return nil, false
	}
	if time.Now().After(item.expiry) {
		c.delete(key)
		return nil, false
	}
	return item.value, true
}

func (c *Cache) Set(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[key] = cacheItem{
		value:  value,
		expiry: time.Now().Add(c.ttl),
	}
}

func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.delete(key)
}

func (c *Cache) delete(key string) {
	delete(c.data, key)
}

func (c *Cache) startGC() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.gcRunning {
		return
	}
	c.gcRunning = true

	ticker := time.NewTicker(c.ttl)
	for {
		select {
		case <-ticker.C:
			c.mutex.Lock()
			for key := range c.data {
				if time.Now().After(c.data[key].expiry) {
					c.delete(key)
				}
			}
			c.mutex.Unlock()
		case <-c.shutdown:
			ticker.Stop()
			c.gcRunning = false
			return
		}
	}
}

func (c *Cache) Shutdown() {
	c.shutdown <- struct{}{}
}

func main() {
	// Create a cache with TTL of 1 minute
	cache := NewCache(time.Second)

	// Set data
	cache.Set("key1", "value1")
	cache.Set("key2", 42)

	// Get data
	value1, ok1 := cache.Get("key1")
	fmt.Println("key1:", value1, ok1) // Output: key1: value1 true

	value2, ok2 := cache.Get("key2")
	fmt.Println("key2:", value2, ok2) // Output: key2: 42 true

	// Wait for TTL to expire
	time.Sleep(time.Second)

	// Get expired data
	value1, ok1 = cache.Get("key1")
	fmt.Println("key1:", value1, ok1) // Output: key1: <nil> false

	// Delete data
	cache.Delete("key2")

	// Check if data is deleted
	value2, ok2 = cache.Get("key2")
	fmt.Println("key2:", value2, ok2) // Output: key2: <nil> false

	// Shutdown the cache
	cache.Shutdown()
}
