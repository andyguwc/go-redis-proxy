package proxy_test

import (
	"testing"
	"net/http"
	"net/http/httptest"
	
	"github.com/andyguwc/go-redis-cache/proxy"
	"github.com/julienschmidt/httprouter"
	redis "github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

var capacity = 10
var globalExpiry = 100
	
func TestHandler(t *testing.T) {
	RInstance := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",	
	})

	p := proxy.New(RInstance, capacity, globalExpiry)
	p.CacheDB.Add("foo", "bar")

	router := httprouter.New()
	router.GET("/GET/:key", p.GetHandler)
	
    req, _ := http.NewRequest("GET", "/GET/foo", nil)
    rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "wrong status")
}

