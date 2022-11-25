package main

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/takez0o/honestwork-api/utils/client"
)

type ReturnJson struct {
	Value int `json:"value"`
	Age   int `json:"age"`
}

func main() {
	app := fiber.New()

	redis := client.NewClient()

	app.Get("/age", func(c *fiber.Ctx) error {
		var returnJ ReturnJson

		data, err := redis.Do(redis.Context(), "JSON.GET", "testJson").Result()
		if err != nil {
			fmt.Println("Error:", err)
		}

		err = json.Unmarshal([]byte(fmt.Sprint(data)), &returnJ)
		if err != nil {
			fmt.Println("Error:", err)
		}

		// return c.SendString(strconv.Itoa(returnJ.Age))
		return c.JSON(returnJ)
	})

	app.Post("/newuser", func(c *fiber.Ctx) error {
		payload := struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}{}

		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		cmd := redis.Do(redis.Context(), "JSON.SET", "testJason", "$", c.Body())
		fmt.Println(cmd)
		return c.JSON(payload)
	})

	app.Listen(":3000")
}
