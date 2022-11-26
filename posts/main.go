package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/takez0o/honestwork-api/utils/client"
	"github.com/takez0o/honestwork-api/utils/config"
	"github.com/takez0o/honestwork-api/utils/crypto"
	"github.com/takez0o/honestwork-api/utils/schema"
)

func main() {
	app := fiber.New()

	conf, err := config.ParseConfig("../config.yaml")
	if err != nil {
		fmt.Println("Error:", err)
	}

	redis := client.NewClient(conf.DB.Posts.ID)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Healthy looking API, doc.")
	})

	app.Get("/posts/:id", func(c *fiber.Ctx) error {
		var user schema.User

		data, err := redis.Do(redis.Context(), "JSON.GET", c.Params("id")).Result()
		if err != nil {
			fmt.Println("Error:", err)
		}

		err = json.Unmarshal([]byte(fmt.Sprint(data)), &user)
		if err != nil {
			fmt.Println("Error:", err)
		}
		return c.JSON(user)
	})

	app.Get("/test", func(c *fiber.Ctx) error {
		res := crypto.VerifySignatureTest()
		return c.SendString(strconv.FormatBool(res))
	})

	app.Post("/posts/new", func(c *fiber.Ctx) error {
		var user schema.User
		if err := c.BodyParser(&user); err != nil {
			return err
		}

		redis.Do(redis.Context(), "JSON.SET", "testJbson", "$", c.Body())
		return c.JSON(user)
	})

	app.Listen(":" + conf.API.Posts.Port)
}
