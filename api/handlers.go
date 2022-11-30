package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/takez0o/honestwork-api/utils/crypto"
	"github.com/takez0o/honestwork-api/utils/schema"
)

// todo: fix error handling

func HandleGetAddress(redis *redis.Client, address string) schema.User {
	var user schema.User

	data, err := redis.Do(redis.Context(), "JSON.GET", address).Result()
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = json.Unmarshal([]byte(fmt.Sprint(data)), &user)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return user
}

func HandleUserUpdate(redis *redis.Client, address string, signature string, body []byte) string {
	result := crypto.VerifySignature("post", address, signature)
	if !result {
		return "Wrong signature."
	}

	// new user
	var user schema.User
	err := json.Unmarshal(body, &user) // todo: validate input after unmarshal
	if err != nil {
		fmt.Println("Error:", err)
	}

	// current user in db
	var user_db schema.User
	data, err := redis.Do(redis.Context(), "JSON.GET", address).Result()
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = json.Unmarshal([]byte(fmt.Sprint(data)), &user_db)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// set skills in input json
	user.Skills = user_db.Skills

	// marshal back to bytes
	new_data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error:", err)
	}

	redis.Do(redis.Context(), "JSON.SET", address, "$", new_data)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return "success"
}

func HandleAddSkill(redis *redis.Client, address string, signature string, body []byte) string {
	return ""
}

func HandleGetSkill(redis *redis.Client, address string, slot int) schema.Skill {
	var skill schema.Skill
	return skill
}

func HandleGetSkills(redis *redis.Client, address string) []schema.Skill {
	var skills []schema.Skill
	return skills
}

func HandleUpdateSkill(redis *redis.Client, address string, signature string, slot int, body []byte) schema.Skill {
	var skill schema.Skill
	return skill
}
