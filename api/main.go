package main

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
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
	redis_skill_index := client.NewSearchClient("skillIndex")
	redis_job_index := client.NewSearchClient("jobIndex")

	app.Use(cors.New())
	app.Use(recover.New())

	// user routes
	app.Post("/api/v1/users/:address/:signature", func(c *fiber.Ctx) error {
		return c.JSON(HandleSignup(redis, c.Params("address"), c.Params("signature")))
	})
	app.Get("/api/v1/users/:address", func(c *fiber.Ctx) error {
		return c.JSON(HandleGetUser(redis, c.Params("address")))
	})
	app.Patch("/api/v1/users/:address/:signature", func(c *fiber.Ctx) error {
		return c.JSON(HandleUserUpdate(redis, c.Params("address"), c.Params("signature"), c.Body()))
	})

	// skill routes
	app.Get("/api/v1/skills/total", func(c *fiber.Ctx) error {
		return c.JSON(HandleGetSkillsTotal(redis_skill_index))
	})
	app.Get("/api/v1/skills/limit/:offset/:size", func(c *fiber.Ctx) error {
		offset, _ := strconv.Atoi(c.Params("offset"))
		size, _ := strconv.Atoi(c.Params("size"))
		return c.JSON(HandleGetSkillsLimit(redis_skill_index, offset, size))
	})
	app.Get("/api/v1/skills/:sort/:order", func(c *fiber.Ctx) error {
		asc, _ := strconv.ParseBool(c.Params("order"))
		return c.JSON(HandleGetAllSkills(redis_skill_index, c.Params("sort"), asc))
	})
	app.Get("/api/v1/skills/:address", func(c *fiber.Ctx) error {
		return c.JSON(HandleGetSkills(redis_skill_index, c.Params("address")))
	})
	app.Get("/api/v1/skill/:address/:slot", func(c *fiber.Ctx) error {
		return c.JSON(HandleGetSkill(redis, c.Params("address"), c.Params("slot")))
	})
	app.Post("/api/v1/skills/:address/:signature", func(c *fiber.Ctx) error {
		return c.JSON(HandleAddSkill(redis, redis_skill_index, c.Params("address"), c.Params("signature"), c.Body()))
	})
	app.Patch("/api/v1/skills/:address/:signature/:slot", func(c *fiber.Ctx) error {
		return c.JSON(HandleUpdateSkill(redis, c.Params("address"), c.Params("signature"), c.Params("slot"), c.Body()))
	})

	// job routes
	app.Get("/api/v1/jobs/total", func(c *fiber.Ctx) error {
		return c.JSON(HandleGetJobsTotal(redis_job_index))
	})
	app.Get("/api/v1/jobs/limit/:offset/:size", func(c *fiber.Ctx) error {
		offset, _ := strconv.Atoi(c.Params("offset"))
		size, _ := strconv.Atoi(c.Params("size"))
		return c.JSON(HandleGetJobsLimit(redis_job_index, offset, size))
	})
	app.Get("/api/v1/jobs/:sort/:order", func(c *fiber.Ctx) error {
		asc, _ := strconv.ParseBool(c.Params("order"))
		return c.JSON(HandleGetAllJobs(redis_job_index, c.Params("sort"), asc))
	})
	app.Get("/api/v1/job/:address/:slot", func(c *fiber.Ctx) error {
		return c.JSON(HandleGetJob(redis, c.Params("address"), c.Params("slot")))
	})
	app.Get("/api/v1/jobs/:address", func(c *fiber.Ctx) error {
		return c.JSON(HandleGetJobs(redis_job_index, c.Params("address")))
	})
	app.Get("/api/v1/jobs/feed", func(c *fiber.Ctx) error {
		return c.JSON(HandleGetJobsFeed(redis_job_index))
	})
	app.Post("/api/v1/jobs/:address/:signature", func(c *fiber.Ctx) error {
		return c.JSON(HandleAddJob(redis, redis_job_index, c.Params("address"), c.Params("signature"), c.Body()))
	})
	app.Patch("/api/v1/jobs/:address/:signature", func(c *fiber.Ctx) error {
		return c.JSON(HandleUpdateJob(redis, c.Params("address"), c.Params("signature"), c.Body()))
	})
	app.Post("/api/v1/jobs/apply/:address/:signature/:recruiter_address/:slot/", func(c *fiber.Ctx) error {
		return c.JSON(HandleApplyJob(redis, c.Params("address"), c.Params("signature"), c.Params("recruiter_address"), c.Params("slot"), c.Body()))
	})

	// watchlist (for jobs listings)
	app.Post("/api/v1/watchlist/:address/:signature", func(c *fiber.Ctx) error {
		return c.JSON(HandleAddWatchlist(redis, c.Params("address"), c.Params("signature"), c.Body()))
	})
	app.Delete("/api/v1/watchlist/:address/:signature", func(c *fiber.Ctx) error {
		return c.JSON(HandleRemoveWatchlist(redis, c.Params("address"), c.Params("signature"), c.Body()))
	})
	app.Get("/api/v1/watchlist/:address", func(c *fiber.Ctx) error {
		return c.JSON(HandleGetWatchlist(redis, c.Params("address")))
	})

	// favorites (for skills)
	app.Post("/api/v1/favorites/:address/:signature", func(c *fiber.Ctx) error {
		return c.JSON(HandleAddFavorite(redis, c.Params("address"), c.Params("signature"), c.Body()))
	})
	app.Delete("/api/v1/favorites/:address/:signature", func(c *fiber.Ctx) error {
		return c.JSON(HandleRemoveFavorite(redis, c.Params("address"), c.Params("signature"), c.Body()))
	})
	app.Get("/api/v1/favorites/:address", func(c *fiber.Ctx) error {
		return c.JSON(HandleGetFavorites(redis, c.Params("address")))
	})

	app.Get("api/v1/salt/:address", func(c *fiber.Ctx) error {
		return c.JSON(HandleGetSalt(redis, c.Params("address")))
	})

	app.Get("api/v1/verify/:address/:signature", func(c *fiber.Ctx) error {
		return c.JSON(HandleVerify(redis, c.Params("address"), c.Params("signature")))
	})

	app.Get("api/v1/tags", func(c *fiber.Ctx) error {
		return c.JSON(HandleGetTags(redis))
	})
	app.Post("api/v1/tags/:address/:signature/:tag", func(c *fiber.Ctx) error {
		return c.JSON(HandleAddTag(redis, c.Params("address"), c.Params("signature"), c.Params("tag")))
	})

	// conversations
	app.Get("/api/v1/conversations/:address", func(c *fiber.Ctx) error {
		return c.JSON(HandleGetConversations(redis, c.Params("address")))
	})
	// app.Post("/api/v1/conversations/:address/:signature", func(c *fiber.Ctx) error {
	// 	return c.JSON(HandleAddConversation(redis, c.Params("address"), c.Params("signature"), c.Body()))
	// })

	app.Listen(":" + conf.API.Port)
}
