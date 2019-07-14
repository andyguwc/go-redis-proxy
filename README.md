# Redis Proxy Exercise

> This is implemented as a HTTP service which adds a local cache on top of the backing Redis instance. Only supports GET command. 


## Setup
**Build and Test**

> First clone the file
```shell

$ git clone https://github.com/andyguwc/go-redis-proxy.git
$ cd go-redis-proxy
```
> One click build using Docker
```
$ docker-compose up -d
```

> Run tests
```
$ make test 
```

> Stop container
```
$ make stop
```

**Config**

The following parameters are configurable in the docker-compose file:
 - REDIS_ADDRESS: Address of the backing Redis (default :6379)
 - GLOBAL_EXPIRY: Cache expiry time (default 200 ms)
 - CAPACITY: Capacity (default 100)
 - PORT: TCP/IP port number the proxy listens on (default :8080)


## Architecture Overview
**Overall Design**

Overall there are two main components: the HTTP Redis Proxy and LRU Cache. 


```
go-redis-cache/                    
   ├── cache/                      
   │   ├── cache.go               # — LRU cache
   ├── proxy/                      
   │   ├── proxy.go               # — HTTP proxy
   ├── vendor/                    # - dependencies via Govendor management
   ├── main.go                   
   ├── Dockerfile                 # - docker image for proxy service
   ├── docker-compose.yml         # - docker compose file linking proxy and redis service
   ├── Makefile                   # - one click build
```

**Proxy**

The Proxy starts an HTTP server which handles GET requests ```/GET/:key```. The Proxy directs HTTP Get command to first look for value in local Cache, and fallback on the backing Redis instance if local cache doesn't have the key. 

**LRU Cache**

The implementation of LRU Cache relies on the <a href="https://github.com/hashicorp/golang-lru" target="_blank">golang-lru package</a>. Both the ```Get(key)``` and ```Add(key, val)``` methods have a time complexity of ```O(1)```. Keys are evicted based on the Expiry time as configured. If a ```(key, val)``` pair was obtained from Redis yet not existing in Cache, then we use the ```Add(key, val)``` method to add the pair to Cache.


## Time Estimate

- Initial Research (~1hr)
- Cache (~2hrs): initially takes a bit on implementation from scratch then found the golang-lru package! 
- HTTP Proxy (~2hrs): fairly straightforward
- Docker / One Click Build / Tests (~8hrs): Docker-compose newbie here. Also learned the assert package to write cleaner Go tests.
- Documentation: (~1hr)


## Omitted Requirements
- Ideally would like to have more test coverage
- Redis instance could be configured as a pool and supports concurrency
- Could have implemented the LRU algorithm from scratch but opted for more efficiency here :) 
