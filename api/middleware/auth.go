package middleware

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
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

func AuthorizeGuest2() fiber.Handler {
	return func(c *fiber.Ctx) error {
		salt_controller := controller.NewSaltController(c.Params("address"))
		salt, err := salt_controller.GetSalt()
		if err != nil {
			return err
		}
		var message string
		// todo: add other types signature content
		if c.Route().Path == "/api/v1/jobs/:address/:signature" && c.Method() == "POST" {
			message = fmt.Sprintf("HonestWork: New Job Post\n%s\n\nFor more info: https://docs.honestwork.app", salt)
		} else {
			message = salt
		}
		result := crypto.VerifySignature(message, c.Params("address"), c.Params("signature"))
		if !result {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
		}
		err = salt_controller.DeleteSalt()
		if err != nil {
			return err
		}
		return c.Next()
	}
}
func AuthorizeGuest(address string, signature string) error {
	salt_controller := controller.NewSaltController(address)
	salt, err := salt_controller.GetSalt()
	if err != nil {
		return err
	}

	result := crypto.VerifySignature(salt, address, signature)
	if !result {
		return errors.New("Invalid signature")
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
