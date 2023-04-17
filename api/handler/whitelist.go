package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/takez0o/honestwork-api/api/controller"
)

func HandleWhitelist() fiber.Handler {
	return func(c *fiber.Ctx) error {
		wl_controller := controller.NewWhitelistController()
		return c.JSON(wl_controller.GetWhitelist())
	}
}
