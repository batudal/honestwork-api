package main

import (
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"

	"github.com/go-redis/redis/v8"
	"github.com/takez0o/honestwork-api/utils/parser"
	"github.com/takez0o/honestwork-api/utils/schema"
	"github.com/takez0o/honestwork-api/utils/web3"
)

func ValidateUserInput(redis *redis.Client, user *schema.User, address string) bool {
	validate := validator.New()
	err := validate.StructExcept(user, "watchlist", "favorites", "rating")
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println("Error:", err)
		}
		return false
	}
	token_id, _ := strconv.Atoi(user.NFTId)
	if !web3.CheckNFTOwner(address, user.NFTAddress, token_id) {
		fmt.Println("NFT Error")
		return false
	}
	if web3.CheckENSOwner(address, user.EnsName) {
		fmt.Println("ENS Error")
		return false
	}
	bio_length := len(parser.ParseContent(user.Bio))
	if bio_length < 200 || bio_length > 2000 {
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
	description_length := len(parser.ParseContent(skill.Description))
	if description_length < 200 || description_length > 2000 {
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
	description_length := len(parser.ParseContent(job.Description))
	if description_length > 200 || description_length < 2000 {
		return false
	}
	return true
}
