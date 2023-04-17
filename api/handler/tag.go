package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/takez0o/honestwork-api/api/controller"
	"github.com/takez0o/honestwork-api/utils/schema"
)

func HandleGetTags() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tag_controller := controller.NewTagController()
		tags, err := tag_controller.GetTags()
		if err != nil {
			return c.JSON(schema.Tags{})
		}
		return c.JSON(tags)
	}
}

func HandleAddTag() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tag_controller := controller.NewTagController()
		tags, _ := tag_controller.GetTags()

		for _, t := range tags.Tags {
			if t == c.Params("tag") {
				return fiber.NewError(fiber.StatusBadRequest, "tag already exists")
			}
		}
		tags.Tags = append(tags.Tags, c.Params("tag"))

		err := tag_controller.SetTags(&tags)
		if err != nil {
			return err
		}
		return c.SendString("success")
	}
}
