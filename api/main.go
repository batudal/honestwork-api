package main

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/takez0o/honestwork-api/utils/client"
	"github.com/takez0o/honestwork-api/utils/config"
)

func main() {
	app := fiber.New()

	conf, err := config.ParseConfig("../config.yaml")
	if err != nil {
		fmt.Println("Error:", err)
	}

	redis := client.NewClient(conf.DB.Users.ID)

	app.Use(cors.New())

	app.Get("/users/:address", func(c *fiber.Ctx) error {
		return c.JSON(HandleGetAddress(redis, c.Params("address")))
	})

	app.Post("/users/update/:address/:signature", func(c *fiber.Ctx) error {
		return c.JSON(HandleUserUpdate(redis, c.Params("address"), c.Params("signature"), c.Body()))
	})

	app.Get("/skills/:address", func(c *fiber.Ctx) error {
		return c.JSON(HandleGetSkills(redis, c.Params("address")))
	})

	app.Get("/skills/:address/:slot", func(c *fiber.Ctx) error {
		s, _ := strconv.Atoi(c.Params("slot"))
		return c.JSON(HandleGetSkill(redis, c.Params("address"), s))
	})

	app.Post("/skills/add/:address/:signature", func(c *fiber.Ctx) error {
		return c.JSON(HandleAddSkill(redis, c.Params("address"), c.Params("signature"), c.Body()))
	})

	app.Post("/skills/update/:address/:signature/:slot", func(c *fiber.Ctx) error {
		s, _ := strconv.Atoi(c.Params("slot"))
		return c.JSON(HandleUpdateSkill(redis, c.Params("address"), c.Params("signature"), s, c.Body()))
	})

	app.Listen(":" + conf.API.Users.Port)
}
