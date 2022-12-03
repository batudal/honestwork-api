package main

import (
	"github.com/takez0o/honestwork-api/utils/config"
)

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
