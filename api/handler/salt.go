package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/takez0o/honestwork-api/api/controller"
	"github.com/takez0o/honestwork-api/utils/crypto"
)

func HandleAddSalt() fiber.Handler {
	return func(c *fiber.Ctx) error {
		salt_controller := controller.NewSaltController(c.Params("address"))
		salt := crypto.GenerateSalt()
		salt_controller.AddSalt(salt)
		return c.JSON(salt)
	}
}
