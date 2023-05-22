package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/takez0o/honestwork-api/utils/client"
	"github.com/takez0o/honestwork-api/utils/schema"

	"github.com/RediSearch/redisearch-go/redisearch"
)

func MigrateUsers(wg *sync.WaitGroup) {
	mongo_client := client.NewMongoClient()
	collection := mongo_client.Database("honestwork-cluster").Collection("users")
	collection.Drop(context.TODO())
	users, err := getUsers()
	if err != nil {
		panic(err)
	}
	writeUsersToMongo(users)
	wg.Done()
}

type User struct {
	User          schema.User           `json:"user" bson:"user"`
	Address       string                `json:"address" bson:"address"`
	Conversations []schema.Conversation `json:"conversations" bson:"conversations"`
}

func writeUsersToMongo(users []User) {
	mongo_client := client.NewMongoClient()
	collection := mongo_client.Database("honestwork-cluster").Collection("users")
	convss, err := GetConversations()
	if err != nil {
		panic(err)
	}
	to_write := make([]interface{}, len(users))
	for i, user := range users {
		for key, conv := range convss {
			if key == user.Address {
				user.Conversations = conv
			}
		}
		to_write[i] = user
	}
	_, err = collection.InsertMany(context.TODO(), to_write)
	if err != nil {
		panic(err)
	}
}

func getUsers() ([]User, error) {
	redis := client.NewRedisSearchClient("userIndex")
	data, _, err := redis.Search(redisearch.NewQuery("*"))
	if err != nil {
		return []User{}, err
	}
	var users []User
	for _, d := range data {
		translationKeys := make([]string, 0, len(d.Properties))
		for key := range d.Properties {
			translationKeys = append(translationKeys, key)
		}
		var user User
		err = json.Unmarshal([]byte(fmt.Sprint(d.Properties[translationKeys[0]])), &user.User)
		if err != nil {
			return []User{}, err
		}
		addr := strings.Split(d.Id, ":")[1]
		user.Address = addr
		users = append(users, user)
	}
	return users, nil
}
