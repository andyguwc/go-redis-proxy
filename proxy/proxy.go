package proxy 

import (
	// "context"
	"fmt"
	"net/http"
	
	"github.com/andyguwc/go-redis-cache/cache"
	"github.com/julienschmidt/httprouter"
	redis "github.com/go-redis/redis"
)


// Proxy service - contains cacheDB and redis instance
type Proxy struct {
	CacheDB *cache.CacheDB
	RInstance *redis.Client
}

func New(RInstance *redis.Client, capacity int, globalExpiry int) *Proxy {
	c := cache.New(capacity, globalExpiry)
	return &Proxy{
		CacheDB: c,
		RInstance: RInstance,
	}
}

// Proxy Serve method which passes to a Redis pass through handler 
func (p *Proxy) Serve(port string) error {
	Addr := ":" + port
	router := httprouter.New()
	router.GET("/GET/:key", p.GetHandler)
	fmt.Printf("Listening on port %v\n", port)
	return http.ListenAndServe(Addr, router)
}

// Handler first checks CacheDB, if not get from Redis instance
func (p *Proxy) GetHandler (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key := ps.ByName("key")
	// Get value from cache first
	if val, ok := p.CacheDB.Get(key); ok {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(val))
		return 
	}

	// Get the value from redis, cache it and return
	val, err := p.RInstance.Get(key).Result()
	if err == redis.Nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404: Requested key not found: " + key))
		return 
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500: Internal Service Error"))
		return 
	} 
        
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(val))
	p.CacheDB.Add(key, val)
	return 
}