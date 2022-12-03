package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/takez0o/honestwork-api/utils/config"
	"github.com/takez0o/honestwork-api/utils/schema"
)

func AuthorizeSignature(redis *redis.Client, address string, salt string, signature string) bool {
	var user_db schema.User
	data, err := redis.Do(redis.Context(), "JSON.GET", address).Result()
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = json.Unmarshal([]byte(fmt.Sprint(data)), &user_db)
	if err != nil {
		fmt.Println("Error:", err)
	}
	if (user_db.Salt == salt) && (user_db.Signature == signature) {
		return true
	}
	return false
}

func ValidateUsername(username string) bool {
	conf, err := config.ParseConfig()
	if len(username) > conf.CharLimits.Profile.Username || err != nil {
		return false
	}
	return true
}

func ValidateShowEns(address string) bool {
	return true
}

func ValidateTitle(title string) bool {
	conf, err := config.ParseConfig()
	if len(title) > conf.CharLimits.Profile.Title || err != nil {
		return false
	}
	return true
}

func ValidateImageUrl(url string) bool {
	return true
}

func ValidateNFT(address string, id int) bool {
	return true
}

func ValidateEmail(email string) bool {
	return true
}

func ValidateTimezone(timezone string) bool {
	return true
}

func ValidateBio(bio string) bool {
	conf, err := config.ParseConfig()
	if len(bio) > conf.CharLimits.Profile.Bio || err != nil {
		return false
	}
	return true
}

func ValidateLinks(links []string) bool {
	return true
}
