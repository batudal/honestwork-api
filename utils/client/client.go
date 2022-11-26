package client

import (
	"github.com/go-redis/redis/v8"
)

func NewClient(id int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       id, // use default DB
	})
	return client
}
