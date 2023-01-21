package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"

	"github.com/go-redis/redis/v8"
	"github.com/takez0o/honestwork-api/utils/schema"
)

func ValidateUserInput(redis *redis.Client, user *schema.User) bool {
	validate := validator.New()
	err := validate.StructExcept(user, "watchlist", "favorites", "rating")
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println("Error:", err)
		}
		return false
	}
	return true
}

func ValidateSkillInput(redis *redis.Client, skill *schema.Skill) bool {
	validate := validator.New()
	err := validate.StructExcept(skill, "created_at", "user_address")
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println("Error:", err)
		}
		return false
	}
	return true
}

func ValidateJobInput(redis *redis.Client, job *schema.Job) bool {
	validate := validator.New()
	err := validate.StructExcept(job, "created_at", "application", "slot")
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println("Error:", err)
		}
		return false
	}
	return true
}
