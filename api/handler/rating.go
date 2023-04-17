package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/takez0o/honestwork-api/api/controller"
)

func HandleGetRating() fiber.Handler {
	return func(c *fiber.Ctx) error {
		rating_controller := controller.NewRatingController(c.Params("address"))
		rating, err := rating_controller.GetRating()
		if err != nil {
			return c.JSON("0")
		}
		return c.JSON(rating)
	}
}
