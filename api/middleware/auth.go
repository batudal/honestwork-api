package middleware

import (
	"github.com/takez0o/honestwork-api/api/controller"
	"github.com/takez0o/honestwork-api/utils/crypto"
)

func Authorize(address string, signature string) bool {
	user_controller := controller.NewUserController(address)
	user, err := user_controller.GetUser()
	if err != nil {
		return false
	}
	result := crypto.VerifySignature(user.Salt, address, signature)
	if !result {
		return false
	}
	return true
}

func AuthorizeGuest(address string, signature string) (string, error) {
	salt_controller := controller.NewSaltController(address)
	salt, err := salt_controller.GetSalt()
	if err != nil {
		return "", err
	}
	result := crypto.VerifySignature(salt, address, signature)
	if !result {
		return "", err
	}
	return salt, nil
}
