package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/takez0o/honestwork-api/utils/config"
	"github.com/takez0o/honestwork-api/utils/crypto"
	"github.com/takez0o/honestwork-api/utils/schema"
	"github.com/takez0o/honestwork-api/utils/web3"
)

// 0-no tokens, 1-not soulbound, 2-soulbound(tier 1), 3-soulbound(tier 2), 4-soulbound(tier-3)
// todo: fix error handling
// todo: move validation to middleware
func getUserFromAddress(redis *redis.Client, address string) schema.User {
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

// todo: impletement all validators
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
	var user schema.User
	data, err := redis.Do(redis.Context(), "JSON.GET", address).Result()
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = json.Unmarshal([]byte(fmt.Sprint(data)), &user)
	if err != nil {
		fmt.Println("Error:", err)
	}

	state := web3.FetchUserState(address)
	switch state {
	case 0:
		return "User doesn't have NFT."
	// case 1, 2, 3, 4:
	// 	if (user.Signature != "") && (user.Salt != "") {
	// 		return "User already signed up."
	// 	}
	}

	user.Salt = salt
	user.Signature = signature

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

func HandleGetUser(redis *redis.Client, address string) schema.User {
	user := getUserFromAddress(redis, address)
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
	var user_db schema.User
	data, err := redis.Do(redis.Context(), "JSON.GET", address).Result()
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = json.Unmarshal([]byte(fmt.Sprint(data)), &user_db)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// filter
	user.Skills = user_db.Skills
	user.Salt = user_db.Salt
	user.Signature = user_db.Signature
	if user.ImageUrl == "" {
		user.ImageUrl = user_db.ImageUrl
	}

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

func HandleGetSkills(redis *redis.Client, address string) []schema.Skill {
	user := getUserFromAddress(redis, address)
	return user.Skills
}

func HandleGetSkill(redis *redis.Client, address string, slot string) schema.Skill {
	s, _ := strconv.Atoi(slot)
	user := getUserFromAddress(redis, address)
	return user.Skills[s]
}

func HandleAddSkill(redis *redis.Client, address string, salt string, signature string, body []byte) string {
	authorized := authorize(redis, address, salt, signature)
	if !authorized {
		return "Wrong signature."
	}

	user := getUserFromAddress(redis, address)
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
	if len(user.Skills) == max_allowed {
		return "User reached skill limit."
	}

	var skill schema.Skill
	err := json.Unmarshal(body, &skill)
	if err != nil {
		fmt.Println("Error:", err)
	}

	user.Skills = append(user.Skills, skill)
	updated_user, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error:", err)
	}

	redis.Do(redis.Context(), "JSON.SET", address, "$", updated_user)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return "success"
}

func HandleUpdateSkill(redis *redis.Client, address string, salt string, signature string,slot string, body []byte) string {
	authorized := authorize(redis, address, salt, signature)
	if !authorized {
		return "Wrong signature."
	}

	s, _ := strconv.Atoi(slot)
	user := getUserFromAddress(redis, address)
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

	if (s > max_allowed - 1) {
		return "User doesn't have that many skill slots."
	}

	var skill schema.Skill
	err := json.Unmarshal(body, &skill)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(skill)

	for index, url := range skill.ImageUrls {
		if url == "" {
			if len(user.Skills[s].ImageUrls) > index {
				skill.ImageUrls[index] = user.Skills[s].ImageUrls[index]
			} else {
				skill.ImageUrls[index] = ""
			}
		}
	}

	user.Skills[s] = skill
	updated_user, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error:", err)
	}

	redis.Do(redis.Context(), "JSON.SET", address, "$", updated_user)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return "success"
}
