package main

import (
	"log"
	"os"
	"strconv"

	"github.com/andyguwc/go-redis-cache/proxy"
	redis "github.com/go-redis/redis"

)

func main() {
	capacity := getEnvInt("CAPACITY", 10)
	globalExpiry := getEnvInt("GLOBAL_EXPIRY", 500)
	port := getEnv("PORT", "8080")
	
	RInstance := redis.NewClient(&redis.Options{
		Addr: getEnv("REDIS_ADDR", "localhost:6379"),	
	})

	p := proxy.New(RInstance, capacity, globalExpiry)
	log.Fatal(p.Serve(port))

}

// getEnv get key environment variable if exist otherwise return defalutValue
func getEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return defaultValue
    }
    return value
}

// also need one for int 
func getEnvInt(key string, defaultValue int) int {
    value, err := strconv.Atoi(os.Getenv(key))
    if err != nil {
        return defaultValue
    }
    return value
}