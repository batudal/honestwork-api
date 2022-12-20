package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/go-redis/redis/v8"
	"github.com/takez0o/honestwork-api/utils/config"
	"github.com/takez0o/honestwork-api/utils/crypto"
	"github.com/takez0o/honestwork-api/utils/schema"
	"github.com/takez0o/honestwork-api/utils/web3"
)

// 0-no tokens, 1-not soulbound, 2-soulbound(tier 1), 3-soulbound(tier 2), 4-soulbound(tier-3)
// todo: fix error handling
// todo: move validation to middleware
func getUser(redis *redis.Client, address string) schema.User {
	record_id := "user:" + address
	var user schema.User
	data, err := redis.Do(redis.Context(), "JSON.GET", record_id).Result()
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = json.Unmarshal([]byte(fmt.Sprint(data)), &user)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return user
}

func getSkill(redis *redis.Client, slot int, address string) schema.Skill {
	s := strconv.Itoa(slot)
	record_id := "skill:" + s + ":" + address
	var skill schema.Skill
	data, err := redis.Do(redis.Context(), "JSON.GET", record_id).Result()
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = json.Unmarshal([]byte(fmt.Sprint(data)), &skill)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return skill
}

func getSkills(redis *redisearch.Client, address string) []schema.Skill {
	// index search with redisearch-go
	data, _, err := redis.Search(redisearch.NewQuery("*").AddFilter(redisearch.Filter{
		Field:   "user_address",
		Options: address,
	}))
	// manual search
	// data, err := redis.Do(redis.Context(), "FT.SEARCH", "skillIndex '@user_adress:(0xC370b50eC6101781ed1f1690A00BF91cd27D77c4)'").Result()
	if err != nil {
		fmt.Println("Error:", err)
	}

	var skills []schema.Skill
	for _, d := range data {
		translationKeys := make([]string, 0, len(d.Properties))
		for key := range d.Properties {
			translationKeys = append(translationKeys, key)
		}
		var skill schema.Skill
		err = json.Unmarshal([]byte(fmt.Sprint(d.Properties[translationKeys[0]])), &skill)
		if err != nil {
			fmt.Println("Error:", err)
		}
		skills = append(skills, skill)
	}
	return skills
}

func getAllSkills(redis *redisearch.Client) []schema.Skill {
	data, _, err := redis.Search(redisearch.NewQuery("*"))
	if err != nil {
		fmt.Println("Error:", err)
	}

	var skills []schema.Skill
	for _, d := range data {
		translationKeys := make([]string, 0, len(d.Properties))
		for key := range d.Properties {
			translationKeys = append(translationKeys, key)
		}
		var skill schema.Skill
		err = json.Unmarshal([]byte(fmt.Sprint(d.Properties[translationKeys[0]])), &skill)
		if err != nil {
			fmt.Println("Error:", err)
		}
		skills = append(skills, skill)
	}

	return skills
}

func getSkillsLimit(redis *redisearch.Client, offset int, size int) []schema.Skill {
	data, _, err := redis.Search(redisearch.NewQuery("*").Limit(offset, size).SetSortBy("created_at", false))
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Data:", data))

	var skills []schema.Skill
	for _, d := range data {
		translationKeys := make([]string, 0, len(d.Properties))
		for key := range d.Properties {
			translationKeys = append(translationKeys, key)
		}
		var skill schema.Skill
		err = json.Unmarshal([]byte(fmt.Sprint(d.Properties[translationKeys[0]])), &skill)
		if err != nil {
			fmt.Println("Error:", err)
		}
		skills = append(skills, skill)
	}

	return skills
}

func getTotalSkills(redis *redisearch.Client) int {
	_, total, err := redis.Search(redisearch.NewQuery("*").Limit(0, 0))
	if err != nil {
		fmt.Println("Error:", err)
	}
	return total
}

func HandleGetTotalSkills(redis *redisearch.Client) int {
	data := getTotalSkills(redis)
	return data
}

func HandleGetSkillsTotal(redis *redisearch.Client) int {
	data := getTotalSkills(redis)
	return data
}

func HandleGetSkillsLimit(redis *redisearch.Client, offset int, size int) []schema.Skill {
	data := getSkillsLimit(redis, offset, size)
	return data
}

func getAllowedSkillAmount(tier int) int {
	conf, err := config.ParseConfig()
	if err != nil {
		fmt.Println("Error:", err)
	}
	switch tier {
	case 1:
		return conf.Settings.Skills.Tier_1
	case 2:
		return conf.Settings.Skills.Tier_2
	case 3:
		return conf.Settings.Skills.Tier_3
	default:
		return 0
	}
}

// todo: implement all validators
func validateUserInput(redis *redis.Client, user schema.User) bool {
	if ValidateUsername(user.Username) &&
		ValidateTitle(user.Title) &&
		ValidateBio(user.Bio) {
		return true
	}
	return false
}

func authorize(redis *redis.Client, address string, salt string, signature string) bool {
	result := crypto.VerifySignature(salt, address, signature)
	if result {
		return AuthorizeSignature(redis, address, salt, signature)
	}
	return false
}

func HandleSignup(redis *redis.Client, address string, salt string, signature string) string {
	result := crypto.VerifySignature(salt, address, signature)
	if !result {
		return "Wrong signature."
	}

	// new user
	user := getUser(redis, address)
	state := web3.FetchUserState(address)
	switch state {
	case 0:
		return "User doesn't have NFT."
	}

	user.Salt = salt
	user.Signature = signature
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

func HandleGetUser(redis *redis.Client, address string) schema.User {
	user := getUser(redis, address)
	return user
}

func HandleUserUpdate(redis *redis.Client, address string, salt string, signature string, body []byte) string {
	authorized := authorize(redis, address, salt, signature)
	if !authorized {
		return "Wrong signature."
	}

	state := web3.FetchUserState(address)
	switch state {
	case 0:
		return "User doesn't have NFT."
	case 1:
		return "User didn't bind NFT yet."
	}

	// new user
	var user schema.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if !validateUserInput(redis, user) {
		return "Invalid input."
	}

	// current user in db
	user_db := getUser(redis, address)

	// filter
	user.Salt = user_db.Salt
	user.Signature = user_db.Signature
	if user.ImageUrl == "" {
		user.ImageUrl = user_db.ImageUrl
	}

	new_data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error:", err)
	}

	record_id := "user:" + address

	redis.Do(redis.Context(), "JSON.SET", record_id, "$", new_data)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return "success"
}

func HandleGetSkills(redis *redisearch.Client, address string) []schema.Skill {
	skills := getSkills(redis, address)
	return skills
}

func HandleGetSkill(redis *redis.Client, address string, slot string) schema.Skill {
	s, _ := strconv.Atoi(slot)
	skill := getSkill(redis, s, address)
	return skill
}

func HandleAddSkill(redis *redis.Client, redis_search *redisearch.Client, address string, salt string, signature string, body []byte) string {
	authorized := authorize(redis, address, salt, signature)
	if !authorized {
		return "Wrong signature."
	}

	state := web3.FetchUserState(address)
	var max_allowed int
	switch state {
	case 0:
		return "User doesn't have NFT."
	case 1:
		return "User didn't bind NFT yet."
	case 2:
		max_allowed = getAllowedSkillAmount(1)
	case 3:
		max_allowed = getAllowedSkillAmount(2)
	case 4:
		max_allowed = getAllowedSkillAmount(3)
	}

	all_skills := getSkills(redis_search, address)
	if len(all_skills) == max_allowed {
		return "User reached skill limit."
	}

	var skill schema.Skill
	err := json.Unmarshal(body, &skill)
	if err != nil {
		fmt.Println("Error:", err)
	}

	slot := strconv.Itoa(len(all_skills))
	record_id := "skill:" + slot + ":" + address

	new_data, err := json.Marshal(skill)
	if err != nil {
		fmt.Println("Error:", err)
	}

	redis.Do(redis.Context(), "JSON.SET", record_id, "$", new_data)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return "success"
}

func HandleUpdateSkill(redis *redis.Client, address string, salt string, signature string, slot string, body []byte) string {
	authorized := authorize(redis, address, salt, signature)
	if !authorized {
		return "Wrong signature."
	}

	s, _ := strconv.Atoi(slot)
	current_skill := getSkill(redis, s, address)
	state := web3.FetchUserState(address)
	var max_allowed int
	switch state {
	case 0:
		return "User doesn't have NFT."
	case 1:
		return "User didn't bind NFT yet."
	case 2:
		max_allowed = getAllowedSkillAmount(1)
	case 3:
		max_allowed = getAllowedSkillAmount(2)
	case 4:
		max_allowed = getAllowedSkillAmount(3)
	}

	if s > max_allowed-1 {
		return "User doesn't have that many skill slots."
	}

	var skill schema.Skill
	err := json.Unmarshal(body, &skill)
	if err != nil {
		fmt.Println("Error:", err)
	}

	for index, url := range skill.ImageUrls {
		if url == "" {
			if len(current_skill.ImageUrls) > index {
				skill.ImageUrls[index] = current_skill.ImageUrls[index]
			} else {
				skill.ImageUrls[index] = ""
			}
		}
	}

	new_data, err := json.Marshal(skill)
	if err != nil {
		fmt.Println("Error:", err)
	}
	record_id := "skill:" + slot + ":" + address

	redis.Do(redis.Context(), "JSON.SET", record_id, "$", new_data)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return "success"
}

func HandleGetAllSkills(redis *redisearch.Client) []schema.Skill {
	skills := getAllSkills(redis)
	return skills
}
