package client

import (
	"log"
	"os"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/go-redis/redis/v8"
	redigo "github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
	"github.com/takez0o/honestwork-api/utils/config"
)

func NewRedisClient() *redis.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conf, err := config.ParseConfig()
	if err != nil {
		log.Fatal("Error loading config file")
	}

	password := os.Getenv("REDIS_PASSWORD")
	host := os.Getenv("REDIS_HOST")

	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       conf.DB.ID,
	})
	return client
}

func NewRedisSearchClient(index_name string) *redisearch.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	password := os.Getenv("REDIS_PASSWORD")
	host := os.Getenv("REDIS_HOST")
	pool := &redigo.Pool{Dial: func() (redigo.Conn, error) {
		return redigo.Dial("tcp", host, redigo.DialPassword(password))
	}}
	client := redisearch.NewClientFromPool(pool, index_name)
	return client
}
