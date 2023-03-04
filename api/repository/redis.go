package repository

import (
	"log"
	"os"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/takez0o/honestwork-api/utils/config"
)

func newClient() *redis.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	password := os.Getenv("REDIS_PASSWORD")
	host := os.Getenv("REDIS_HOST")

  // todo: rename into MustParseConfig
	conf, err := config.ParseConfig()
	if err != nil {
	  log.Fatal(err)
  }

	return redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       conf.DB.ID,
	})
}

func JSONRead(record_id string) (interface{},error)  {
  redis := newClient() 
	data, err := redis.Do(redis.Context(), "JSON.GET", record_id).Result()
	if err != nil {
		return nil, err
	}
	return data, nil
}
