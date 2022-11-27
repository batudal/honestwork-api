package main

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	redis := client.NewClient(conf.DB.Users.ID)

	app.Use(cors.New())

	app.Get("/users/:id", func(c *fiber.Ctx) error {
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

	app.Post("/users/new/:address/:signature", func(c *fiber.Ctx) error {
		// change salt to a branded msg
		result := crypto.VerifySignature("post", c.Params("address"), (c.Params("signature")))
		if !result {
			return c.SendString("Wrong signature.")
		}

		var user schema.User
		var user_ schema.User

		err = json.Unmarshal(c.Body(), &user)
		if err != nil {
			fmt.Println("Error:", err)
		}

		data, err := redis.Do(redis.Context(), "JSON.GET", c.Params("address")).Result()
		if err != nil {
			fmt.Println("Error:", err)
		}

		err = json.Unmarshal([]byte(fmt.Sprint(data)), &user_)
		if err != nil {
			fmt.Println("Error:", err)
		}

		if user.Posts != user_.Posts {
			return c.SendString("You can't edit post count manually.")
		}

		user_db := client.NewClient(conf.DB.Users.ID)
		user_db.Do(redis.Context(), "JSON.SET", c.Params("address"), "$", c.Body())
		if err != nil {
			fmt.Println("Error:", err)
		}
		err = json.Unmarshal([]byte(fmt.Sprint(data)), &user)
		if err != nil {
			fmt.Println("Error:", err)
		}

		return c.JSON(user)
	})

	app.Listen(":" + conf.API.Users.Port)
}
