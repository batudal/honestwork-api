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

	conf, err := config.ParseConfig()
	if err != nil {
		fmt.Println("Error:", err)
	}

	redis := client.NewClient(conf.DB.ID)
	redis_search := client.NewSearchClient("skillIndex")

	app.Use(cors.New())

	// user routes
	app.Post("/api/v1/users/:address/:salt/:signature", func(c *fiber.Ctx) error {
		return c.JSON(HandleSignup(redis, c.Params("address"), c.Params("salt"), c.Params("signature")))
	})
	app.Get("/api/v1/users/:address", func(c *fiber.Ctx) error {
		return c.JSON(HandleGetUser(redis, c.Params("address")))
	})
	app.Patch("/api/v1/users/:address/:salt/:signature", func(c *fiber.Ctx) error {
		return c.JSON(HandleUserUpdate(redis, c.Params("address"), c.Params("salt"), c.Params("signature"), c.Body()))
	})

	// skill routes
	app.Get("/api/v1/skills/total", func(c *fiber.Ctx) error {
		return c.JSON(HandleGetSkillsTotal(redis_search))
	})
	app.Get("/api/v1/skills/limit/:offset/:size", func(c *fiber.Ctx) error {
		offset, _ := strconv.Atoi(c.Params("offset"))
		size, _ := strconv.Atoi(c.Params("size"))
		return c.JSON(HandleGetSkillsLimit(redis_search, offset, size))
	})
	app.Get("/api/v1/skills/:sort/:order", func(c *fiber.Ctx) error {
		asc, _ := strconv.ParseBool(c.Params("order"))
		return c.JSON(HandleGetAllSkills(redis_search, c.Params("sort"), asc))
	})
	app.Get("/api/v1/skills/:address", func(c *fiber.Ctx) error {
		return c.JSON(HandleGetSkills(redis_search, c.Params("address")))
	})
	app.Get("/api/v1/skills/:address/:slot", func(c *fiber.Ctx) error {
		return c.JSON(HandleGetSkill(redis, c.Params("address"), c.Params("slot")))
	})
	app.Post("/api/v1/skills/:address/:salt/:signature", func(c *fiber.Ctx) error {
		return c.JSON(HandleAddSkill(redis, redis_search, c.Params("address"), c.Params("salt"), c.Params("signature"), c.Body()))
	})
	app.Patch("/api/v1/skills/:address/:salt/:signature/:slot", func(c *fiber.Ctx) error {
		return c.JSON(HandleUpdateSkill(redis, c.Params("address"), c.Params("salt"), c.Params("signature"), c.Params("slot"), c.Body()))
	})

	app.Listen(":" + conf.API.Port)
}
