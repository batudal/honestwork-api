package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/takez0o/honestwork-api/utils/client"
	"github.com/takez0o/honestwork-api/utils/config"
)

func main() {
	app := fiber.New()

	conf, err := config.ParseConfig()
	if err != nil {
		fmt.Println("Error:", err)
	}

	redis := client.NewClient(conf.DB.ID)

	app.Use(cors.New())

	app.Post("/api/v1/users/:address/:salt/:signature", func(c *fiber.Ctx) error {
		return c.JSON(HandleSignup(redis, c.Params("address"), c.Params("salt"), c.Params("signature")))
	})

	app.Get("/api/v1/users/:address", func(c *fiber.Ctx) error {
		return c.JSON(HandleGetUser(redis, c.Params("address")))
	})

	app.Patch("/api/v1/users/:address/:signature", func(c *fiber.Ctx) error {
		return c.JSON(HandleUserUpdate(redis, c.Params("address"), c.Params("signature"), c.Body()))
	})

	app.Get("/api/v1/skills/:address", func(c *fiber.Ctx) error {
		return c.JSON(HandleGetSkills(redis, c.Params("address")))
	})

	app.Get("/api/v1/skills/:address/:slot", func(c *fiber.Ctx) error {
		return c.JSON(HandleGetSkill(redis, c.Params("address"), c.Params("slot")))
	})

	app.Post("/api/v1/skills/:address/:signature", func(c *fiber.Ctx) error {
		return c.JSON(HandleAddSkill(redis, c.Params("address"), c.Params("signature"), c.Body()))
	})

	app.Patch("/api/v1/skills/:address/:signature/:slot", func(c *fiber.Ctx) error {
		return c.JSON(HandleUpdateSkill(redis, c.Params("address"), c.Params("signature"), c.Params("slot"), c.Body()))
	})

	app.Listen(":" + conf.API.Port)
}
