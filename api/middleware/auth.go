package middleware

import (
	"errors"

	"github.com/takez0o/honestwork-api/api/controller"
	"github.com/takez0o/honestwork-api/utils/crypto"
)

func AuthorizeMember(address string, signature string) error {
	user_controller := controller.NewUserController(address)
	user, err := user_controller.GetUser()
	if err != nil {
		return err
	}
	result := crypto.VerifySignature(user.Salt, address, signature)
	if !result {
		return errors.New("Invalid signature")
	}
	return nil
}

func AuthorizeGuest(address string, signature string) error {
	salt_controller := controller.NewSaltController(address)
	salt, err := salt_controller.GetSalt()
	if err != nil {
		return err
	}
	result := crypto.VerifySignature(salt, address, signature)
	if !result {
		return err
	}
	err = salt_controller.DeleteSalt()
	if err != nil {
		return err
	}
	return nil
}

func AuthorizeUnknown(address string, signature string) error {
	user_controller := controller.NewUserController(address)
	_, err := user_controller.GetUser()
	if err == nil {
		return AuthorizeMember(address, signature)
	}
	return AuthorizeGuest(address, signature)
}
