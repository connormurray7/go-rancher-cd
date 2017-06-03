package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-redis/redis"
)

type server struct {
	client *redis.Client
}

func newServer() *server {
	var s server
	s.client = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	return &s
}

func (server *server) handler(w http.ResponseWriter, r *http.Request) {
	res, _ := server.client.Ping().Result()
	io.WriteString(w, fmt.Sprintf("%s", res))
}

func main() {

}
