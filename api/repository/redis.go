package repository

import (
	"fmt"
	"time"

	"github.com/takez0o/honestwork-api/utils/client"
	"github.com/takez0o/honestwork-api/utils/loggersentry"
)

func JSONRead(record_id string) (interface{}, error) {
	redis := client.NewRedisClient()
	defer redis.Close()
	data, err := redis.Do(redis.Context(), "JSON.GET", record_id).Result()
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "JSONRead")
		return nil, err
	}
	return data, nil
}

func JSONWrite(record_id string, data []byte, ttl time.Duration) error {
	redis := client.NewRedisClient()
	defer redis.Close()
	err := redis.Do(redis.Context(), "JSON.SET", record_id, ".", string(data)).Err()
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "JSONWrite")
		return err
	}
	return nil
}

func JSONDelete(record_id string) error {
	redis := client.NewRedisClient()
	defer redis.Close()
	err := redis.Do(redis.Context(), "JSON.DEL", record_id).Err()
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "JSONDelete")
		return err
	}
	return nil
}

func StringRead(record_id string) (string, error) {
	redis := client.NewRedisClient()
	defer redis.Close()
	data, err := redis.Get(redis.Context(), record_id).Result()
	fmt.Println("Data:", data)
	if err != nil {
		return "", err
	}
	return data, nil
}

func StringWrite(record_id string, data string, ttl time.Duration) error {
	redis := client.NewRedisClient()
	defer redis.Close()
	fmt.Println("Writing string:", record_id, data, ttl)
	err := redis.Set(redis.Context(), record_id, data, ttl).Err()
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "StringWrite")
		return err
	}
	return nil
}

func StringDelete(record_id string) error {
	redis := client.NewRedisClient()
	defer redis.Close()
	err := redis.Del(redis.Context(), record_id).Err()
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "StringDelete")
		return err
	}
	return nil
}
