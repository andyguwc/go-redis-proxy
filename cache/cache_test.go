package cache_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/andyguwc/go-redis-cache/cache"
)

var capacity = 10
var globalExpiry = 100

func TestAdd(t *testing.T) {
	// t.Error() // to indicate test failed
	c := cache.New(capacity, globalExpiry)
	c.Add("foo","bar")
	val, ok := c.Get("foo")
	assert.Equal(t, true, ok, "key not existing")
	assert.Equal(t, "bar", val, "value not set correctly")
}

func TestGetNullKey(t *testing.T) {
	c := cache.New(capacity, globalExpiry)
	val, ok := c.Get("foo3")
	assert.Equal(t, false, ok, "key not existing")
	assert.Equal(t, "", val, "value not set correctly")
}