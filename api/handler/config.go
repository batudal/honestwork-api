package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/takez0o/honestwork-api/utils/config"
)

func HandleConfig() fiber.Handler {
	return func(c *fiber.Ctx) error {
		conf, err := config.ParseConfig()
		if err != nil {
			return c.JSON(config.Config{})
		}
		return c.JSON(*conf)
	}
}
