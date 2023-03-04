package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/takez0o/honestwork-api/api/middleware"
	"github.com/takez0o/honestwork-api/api/worker"

	"github.com/takez0o/honestwork-api/utils/client"
	"github.com/takez0o/honestwork-api/utils/config"
)

func main() {
	app := fiber.New()

	app.Static("/", "./static")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client_key := os.Getenv("CLIENT_KEY")
	client_password := os.Getenv("CLIENT_PASSWORD")

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			client_key: client_password,
		},
	}))

	app.Use(
		logger.New(),
	)

	dsn := os.Getenv("SENTRY_DSN")
	err = sentry.Init(sentry.ClientOptions{
		Dsn:              dsn,
		TracesSampleRate: 1.0, // todo: adjust in production (0.1max)
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)

	conf, err := config.ParseConfig()
	if err != nil {
		sentry.CaptureMessage("Error: " + err.Error())
	}

	redis := client.NewClient(conf.DB.ID)
	redis_skill_index := client.NewSearchClient("skillIndex")
	redis_job_index := client.NewSearchClient("jobIndex")
	redis_user_index := client.NewSearchClient("userIndex")

	go worker.WatchRevenues()
	go worker.WatchRatings(redis_job_index, redis_user_index, redis)

	app.Use(cors.New())
	app.Use(recover.New())

	public_api := app.Group("/api/v1")
	auth_api := app.Group("/api/v1", func(c *fiber.Ctx) error {
		middleware.Authorize(c.Params("address"), c.Params("signature"))
		return c.Next()
	},
	)
	guest_api := app.Group("/api/v1", func(c *fiber.Ctx) error {
		middleware.AuthorizeGuest(c.Params("address"), c.Params("signature"))
		return c.Next()
	},
	)

	// user routes
	guest_api.Post("/users/:address/:signature", func(c *fiber.Ctx) error {
		return c.JSON(HandleSignup(redis, c.Params("address"), c.Params("signature")))
	})
	public_api.Get("/users/:address", func(c *fiber.Ctx) error {
		return c.JSON(HandleGetUser(redis, c.Params("address")))
	})
	auth_api.Patch("/users/:address/:signature", func(c *fiber.Ctx) error {
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
		return c.JSON(HandleGetJob(c.Params("address"), c.Params("slot")))
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
	app.Post("/api/v1/conversations/:address/:signature", func(c *fiber.Ctx) error {
		return c.JSON(HandleAddConversation(redis, redis_job_index, c.Params("address"), c.Params("signature"), c.Body()))
	})

	// deals
	app.Get("/api/v1/deals/:recruiter/:creator", func(c *fiber.Ctx) error {
		return c.JSON(HandleGetDeals(redis, c.Params("recruiter"), c.Params("creator")))
	})
	app.Post("/api/v1/deals/:recruiter/:creator/:signature", func(c *fiber.Ctx) error {
		return c.JSON(HandleAddDeal(redis, c.Params("recruiter"), c.Params("creator"), c.Params("signature"), c.Body()))
	})
	app.Patch("/api/v1/deals/:recruiter/:creator/:signature", func(c *fiber.Ctx) error {
		return c.JSON(HandleSignDeal(redis, c.Params("recruiter"), c.Params("creator"), c.Params("signature"), c.Body()))
	})
	// todo: remove record
	app.Delete("/api/v1/deals/:recruiter/:creator/:signature", func(c *fiber.Ctx) error {
		return c.JSON(HandleExecuteDeal(redis, c.Params("recruiter"), c.Params("creator"), c.Params("signature"), c.Body()))
	})

	// config
	app.Get("/api/v1/config", func(c *fiber.Ctx) error {
		return c.JSON(HandleConfig())
	})

	// rating
	app.Get("/api/v1/rating/:address", func(c *fiber.Ctx) error {
		return c.JSON(HandleGetRating(redis, c.Params("address")))
	})

	app.Listen(":" + conf.API.Port)
}
