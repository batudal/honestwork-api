package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/takez0o/honestwork-api/utils/client"
	"github.com/takez0o/honestwork-api/utils/config"
	"github.com/takez0o/honestwork-api/utils/crypto"
	"github.com/takez0o/honestwork-api/utils/schema"
)

// abstract all into api.go + handlers.go
func main() {
	app := fiber.New()
	conf, err := config.ParseConfig("../config.yaml")
	if err != nil {
		fmt.Println("Error:", err)
	}

	redis := client.NewClient(conf.DB.Posts.ID)

	app.Use(cors.New())

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

	app.Post("/posts/new/:address/:signature", func(c *fiber.Ctx) error {
		result := crypto.VerifySignature("post", c.Params("address"), (c.Params("signature")))
		if !result {
			return c.SendString("Wrong signature.")
		}

		var post schema.Post
		if err := c.BodyParser(&post); err != nil {
			return err
		}

		var user schema.User
		resp, err := http.Get("http://localhost:3002/users/" + c.Params("address"))
		if err != nil {
			print(err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			print(err)
		}
		err = json.Unmarshal([]byte(body), &user)
		if err != nil {
			fmt.Println("Error:", err)
		}

		// deal with response body
		http.Post("http://localhost:3002/users/increment_post/"+c.Params("address")+"/"+c.Params("signature"),
			"application/json", bytes.NewBuffer(body))
		if err != nil {
			print(err)
		}

		salt := "post" + strconv.Itoa(user.Posts+1)
		hash := crypto.GenerateID(salt, c.Params("address"))
		redis.Do(redis.Context(), "JSON.SET", hash, "$", c.Body())

		return c.SendString("1")
	})

	app.Listen(":" + conf.API.Posts.Port)
}
