package main

import (
	"context"
	"strings"
	"sync"

	"github.com/takez0o/honestwork-api/utils/client"
)

func MigrateTxs(wg *sync.WaitGroup) {
	mongo_client := client.NewMongoClient()
	collection := mongo_client.Database("honestwork-cluster").Collection("txs")
	collection.Drop(context.Background())
	txs, err := getTxs()
	if err != nil {
		panic(err)
	}
	writeTxsToMongo(txs)
	wg.Done()
}

func getTxs() (map[string]string, error) {
	redis := client.NewRedisClient()
	results, err := redis.Do(context.Background(), "keys", "tx:*").Result()
	if err != nil {
		return map[string]string{}, err
	}
	txs := make(map[string]string)
	for _, key := range results.([]interface{}) {
		tx, err := redis.Do(context.Background(), "GET", key).Result()
		if err != nil {
			return map[string]string{}, err
		}
		txs[key.(string)] = tx.(string)
	}
	return txs, nil
}

func writeTxsToMongo(txs map[string]string) {
	mongo_client := client.NewMongoClient()
	collection := mongo_client.Database("honestwork-cluster").Collection("txs")
	for key := range txs {
		key_trimmed := strings.Split(key, ":")[1]
		_, err := collection.InsertOne(context.TODO(), map[string]string{
			"key": key_trimmed,
		})
		if err != nil {
			panic(err)
		}
	}
}
