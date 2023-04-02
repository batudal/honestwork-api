package mongo

import (
	"context"
	"fmt"
	"github.com/takez0o/honestwork-api/utils/client"
	"github.com/takez0o/honestwork-api/utils/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"time"
)

type MongoUser struct {
	ID      primitive.ObjectID `bson:"_id"`
	address string
	schema.User
}

func JSONRead(record_id string) (interface{}, error) {
	mongo := client.NewMongoClient()
	defer func() {
		if err := mongo.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	record_arr := strings.Split(record_id, ":")
	coll := mongo.Database("db-0").Collection(record_arr[0])
	var filter bson.D
	if len(record_arr) == 2 {
		filter = bson.D{primitive.E{Key: "address", Value: record_arr[1]}}
	} else {
		filter = bson.D{primitive.E{Key: "address", Value: record_arr[1]}, primitive.E{Key: "slot", Value: record_arr[2]}}
	}
	var result MongoUser
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println("Error")
	}
	return result, nil
}

func JSONWrite(record_id string, data []byte, ttl time.Duration) error {
	mongo := client.NewMongoClient()
	defer func() {
		if err := mongo.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	record_arr := strings.Split(record_id, ":")
	coll := mongo.Database("db-0").Collection(record_arr[0])
	var filter bson.D
	if len(record_arr) == 2 {
		filter = bson.D{primitive.E{Key: "address", Value: record_arr[1]}}
	} else {
		filter = bson.D{primitive.E{Key: "address", Value: record_arr[1]}, primitive.E{Key: "slot", Value: record_arr[2]}}
	}

	_, err := coll.UpdateOne(context.TODO(), filter, data)
	if err != nil {
		return err
	}
	return nil
}

func JSONDelete(record_id string) error {
	redis := client.NewRedisClient()
	defer redis.Close()
	err := redis.Do(redis.Context(), "JSON.DEL", record_id).Err()
	if err != nil {
		return err
	}
	return nil
}
