package main

import (
	"context"
	"sync"

	"github.com/takez0o/honestwork-api/api/controller"
	"github.com/takez0o/honestwork-api/utils/client"
	"github.com/takez0o/honestwork-api/utils/schema"
)

func MigrateJobs(wg *sync.WaitGroup) {
	job_indexer := controller.NewJobIndexer("job")
	jobs, err := job_indexer.GetAllJobs()
	if err != nil {
		panic(err)
	}
	writeToJobsMongo(jobs)
	wg.Done()
}

func writeToJobsMongo(jobs []schema.Job) {
	mongo_client := client.NewMongoClient()
	collection := mongo_client.Database("honestwork-cluster").Collection("jobs")
	to_write := make([]interface{}, len(jobs))
	for i, v := range jobs {
		to_write[i] = v
	}
	_, err := collection.InsertMany(context.TODO(), to_write)
	if err != nil {
		panic(err)
	}
}
