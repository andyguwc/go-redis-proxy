// this one has the cachestore struct 
// also get and set and expire based on LRU

package cache

import (
	"time"
	"sync"
	// "fmt"

	lru "github.com/hashicorp/golang-lru"

)

// Element contains value and Expiry 
type Element struct {
	Value string
	Expiry int64
}

// Cache DB / Store 
type CacheDB struct {
	Cache *lru.Cache
	GlobalExpiry int64
	mutex *sync.Mutex
}


// Initiate Instance of CacheDB
func New(capacity int, globalExpiry int) *CacheDB {
	cache, _ := lru.New(capacity)
	return &CacheDB{
		Cache: cache,
		GlobalExpiry: int64(globalExpiry),
		mutex: &sync.Mutex{},
	}
}

func (c *CacheDB) Get(key string) (val string, ok bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if e, ok :=c.Cache.Get(key); ok {
		element := e.(*Element)
		// if expired
		if c.IsExpired(element) {
			c.Cache.Remove(key) 
			return "", false
		}

		return element.Value, true
	}

	return "", false

}

// Add key, val to LRU cache
func (c *CacheDB) Add(key string, val string) {
	now := time.Now()
    duration := time.Millisecond * time.Duration(c.GlobalExpiry)
	expiry := int64(now.Add(duration).UnixNano())
	
	element := &Element{
		Value: val,
		Expiry: expiry,
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.Cache.Add(key, element)
}


func (c *CacheDB) IsExpired(e *Element) bool {
	now := int64(time.Now().UnixNano()) 
    if e.Expiry - now >= 0 {
        return false
    }

	return true
}





