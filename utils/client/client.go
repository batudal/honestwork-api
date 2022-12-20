package client

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

func NewClient(id int) *redis.Client {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error:", err)
		log.Fatal("Error loading .env file")
	}

	password := os.Getenv("REDIS_PASSWORD")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: password,
		DB:       id,
	})
	return client
}
