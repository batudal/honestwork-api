package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/takez0o/honestwork-api/utils/client"
	"github.com/takez0o/honestwork-api/utils/schema"

	"github.com/RediSearch/redisearch-go/redisearch"
)

func GetConversations() (map[string][]schema.Conversation, error) {
	redis := client.NewRedisSearchClient("convIndex")
	data, _, err := redis.Search(redisearch.NewQuery("*"))
	if err != nil {
		return map[string][]schema.Conversation{}, err
	}
	convss := make(map[string][]schema.Conversation)
	for _, d := range data {
		translationKeys := make([]string, 0, len(d.Properties))
		for key := range d.Properties {
			translationKeys = append(translationKeys, key)
		}
		var convs []schema.Conversation
		err = json.Unmarshal([]byte(fmt.Sprint(d.Properties[translationKeys[0]])), &convs)
		if err != nil {
			return map[string][]schema.Conversation{}, err
		}
		addr := strings.Split(d.Id, ":")[1]
		convss[addr] = convs
	}
	return convss, nil
}
