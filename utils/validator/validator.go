package validator

import (
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"

	"github.com/takez0o/honestwork-api/utils/loggersentry"
	"github.com/takez0o/honestwork-api/utils/parser"
	"github.com/takez0o/honestwork-api/utils/schema"
	"github.com/takez0o/honestwork-api/utils/web3"
)

func ValidateUserInput(user *schema.User, address string) bool {
	validate := validator.New()
	err := validate.StructExcept(user, "watchlist", "favorites", "rating")
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			loggersentry.InitSentry()
			loggersentry.CaptureErrorMessage(err.Error())
		}
		return false
	}
	token_id, _ := strconv.Atoi(user.NFTId)
	if user.ShowNFT == BoolAddr(true) && !web3.CheckNFTOwner(address, user.NFTAddress, token_id) {
		return false
	}
	if user.ShowEns == BoolAddr(true) && !web3.CheckENSOwner(address, user.EnsName) {
		return false
	}
	bio_length := len(parser.Parse(user.Bio))
	if bio_length < 200 || bio_length > 2000 {
		return false
	}
	return true
}

func ValidateSkillInput(skill *schema.Skill) error {
	validate := validator.New()
	err := validate.StructExcept(skill, "created_at", "user_address")
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			loggersentry.InitSentry()
			loggersentry.CaptureErrorMessage(err.Error())
			return err
		}
	}
	description_length := len(parser.Parse(skill.Description))
	if description_length < 200 || description_length > 2000 {
		return fmt.Errorf("Description length is invalid")
	}
	return nil
}

func ValidateJobInput(job *schema.Job) error {
	validate := validator.New()
	err := validate.StructExcept(job, "created_at", "application", "slot")
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			loggersentry.InitSentry()
			loggersentry.CaptureErrorMessage(err.Error())
			return err
		}
	}
	description_length := len(parser.Parse(job.Description))
	if description_length < 200 || description_length > 2000 {
		return fmt.Errorf("Description length is invalid")
	}
	return nil
}

func BoolAddr(b bool) *bool {
	boolVar := b
	return &boolVar
}
