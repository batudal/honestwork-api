package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/takez0o/honestwork-api/api/controller"
	"github.com/takez0o/honestwork-api/utils/crypto"
)

func AuthorizeMember() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user_controller := controller.NewUserController(c.Params("address"))
		user, err := user_controller.GetUser()
		if err != nil {
			return err
		}
		result := crypto.VerifySignature(user.Salt, c.Params("address"), c.Params("signature"))
		if !result {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
		}
		return c.Next()
	}
}

func AuthorizeGuest() fiber.Handler {
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
		} else if c.Route().Path == "/api/v1/jobs/:address/:signature" && c.Method() == "PATCH" {
			message = fmt.Sprintf("HonestWork: Update Job Post\n%s\n\nFor more info: https://docs.honestwork.app", salt)
		} else if c.Route().Path == "/api/v1/users/:address/:signature" && c.Method() == "POST" {
			message = fmt.Sprintf("HonestWork: Login\n%s\n\nFor more info: https://docs.honestwork.app", salt)
		} else if c.Route().Path == "/api/v1/deals/:recruiter/:creator/:signature" && c.Method() == "POST" {
			message = fmt.Sprintf("HonestWork: New Agreement\n%s\n\nFor more info: https://docs.honestwork.app", salt)
		} else if c.Route().Path == "/api/v1/deals/:recruiter/:creator/:signature" && c.Method() == "DELETE" {
			message = fmt.Sprintf("HonestWork: Execute Agreement\n%s\n\nFor more info: https://docs.honestwork.app", salt)
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

func AuthorizeUnknown() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user_controller := controller.NewUserController(c.Params("address"))
		_, err := user_controller.GetUser()
		if err == nil {
			return AuthorizeMember()(c)
		}
		return AuthorizeGuest()(c)
	}
}
