package main

import (
	"context"
  "sync"

	"github.com/takez0o/honestwork-api/api/controller"
	"github.com/takez0o/honestwork-api/utils/client"
	"github.com/takez0o/honestwork-api/utils/schema"
)

func MigrateSkills(wg *sync.WaitGroup) {
  skill_indexer := controller.NewSkillIndexer("skill")
  skills, err := skill_indexer.GetAllSkills()
  if err != nil {
    panic(err)
  }
  writeSkillsToMongo(skills)
  wg.Done()
}

func writeSkillsToMongo(skills []schema.Skill) {
	mongo_client := client.NewMongoClient()
	collection := mongo_client.Database("honestwork-cluster").Collection("skills")
	to_write := make([]interface{}, len(skills))
	for i, v := range skills {
		to_write[i] = v
	}
	_, err := collection.InsertMany(context.TODO(), to_write)
	if err != nil {
		panic(err)
	}
}
