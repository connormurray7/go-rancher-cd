package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-redis/redis"
)

type cache struct {
	client *redis.Client
}

func newCache() *cache {
	var s cache
	s.client = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	return &s
}

func (cache *cache) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, _ := cache.client.Ping().Result()
	io.WriteString(w, fmt.Sprintf("%s", res))
}

func main() {
	cache := newCache()
	http.Handle("/", cache)

	log.Println("Starting server ...")
	http.ListenAndServe(":3000", nil)
}
