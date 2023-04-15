package repository

import (
	"context"
	"log"

	"github.com/takez0o/honestwork-api/utils/client"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Get(collection string, key string, value string) *mongo.Cursor {
	ctx := context.TODO()
	mongo := client.NewMongoClient()
	defer func() {
		if err := mongo.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	coll := mongo.Database("db-0").Collection(collection)
	cursor, err := coll.Find(ctx, bson.D{{key, value}})
	if err != nil {
		log.Fatal(err)
	}
	return cursor
}
