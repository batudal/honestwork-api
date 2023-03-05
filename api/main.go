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
	"github.com/takez0o/honestwork-api/api/handler"
	"github.com/takez0o/honestwork-api/api/middleware"
	"github.com/takez0o/honestwork-api/api/worker"

	"github.com/takez0o/honestwork-api/utils/client"
	"github.com/takez0o/honestwork-api/utils/config"
)

func main() {
	app := fiber.New()

	app.Static("/", "../static")

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

	redis := client.NewClient()
	redis_job_index := client.NewSearchClient("jobIndex")
	redis_user_index := client.NewSearchClient("userIndex")

	go worker.WatchRevenues()
	go worker.WatchRatings(redis_job_index, redis_user_index, redis)

	app.Use(cors.New())
	app.Use(recover.New())

	public_api := app.Group("/api/v1")
	member_api := app.Group("/api/v1", func(c *fiber.Ctx) error {
		middleware.AuthorizeMember(c.Params("address"), c.Params("signature"))
		return c.Next()
	},
	)
	guest_api := app.Group("/api/v1", func(c *fiber.Ctx) error {
		middleware.AuthorizeGuest(c.Params("address"), c.Params("signature"))
		return c.Next()
	},
	)
	unknown_api := app.Group("/api/v1", func(c *fiber.Ctx) error {
		middleware.AuthorizeUnknown(c.Params("address"), c.Params("signature"))
		return c.Next()
	})

	//-----------------//
	//  users          //
	//-----------------//
	guest_api.Post("/users/:address/:signature", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleSignup(c.Params("address"), c.Params("signature")))
	})
	public_api.Get("/users/:address", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleGetUser(c.Params("address")))
	})
	member_api.Patch("/users/:address/:signature", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleUserUpdate(c.Params("address"), c.Params("signature"), c.Body()))
	})

	//-----------------//
	//  skills         //
	//-----------------//

	public_api.Get("/api/v1/skills/total", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleGetSkillsTotal())
	})
	public_api.Get("/api/v1/skills/limit/:offset/:size", func(c *fiber.Ctx) error {
		offset, _ := strconv.Atoi(c.Params("offset"))
		size, _ := strconv.Atoi(c.Params("size"))
		return c.JSON(handler.HandleGetSkillsLimit(offset, size))
	})
	public_api.Get("/api/v1/skills/:sort/:order", func(c *fiber.Ctx) error {
		asc, _ := strconv.ParseBool(c.Params("order"))
		return c.JSON(handler.HandleGetAllSkills(c.Params("sort"), asc))
	})
	public_api.Get("/api/v1/skills/:address", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleGetSkills(c.Params("address")))
	})
	public_api.Get("/api/v1/skill/:address/:slot", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleGetSkill(c.Params("address"), c.Params("slot")))
	})
	member_api.Post("/api/v1/skills/:address/:signature", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleAddSkill(c.Params("address"), c.Params("signature"), c.Body()))
	})
	member_api.Patch("/api/v1/skills/:address/:signature/:slot", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleUpdateSkill(c.Params("address"), c.Params("signature"), c.Params("slot"), c.Body()))
	})

	//-----------------//
	//  jobs           //
	//-----------------//

	public_api.Get("/api/v1/jobs/total", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleGetJobsTotal())
	})
	public_api.Get("/api/v1/jobs/limit/:offset/:size", func(c *fiber.Ctx) error {
		offset, _ := strconv.Atoi(c.Params("offset"))
		size, _ := strconv.Atoi(c.Params("size"))
		return c.JSON(handler.HandleGetJobsLimit(offset, size))
	})
	public_api.Get("/api/v1/jobs/:sort/:order", func(c *fiber.Ctx) error {
		asc, _ := strconv.ParseBool(c.Params("order"))
		return c.JSON(handler.HandleGetAllJobs(c.Params("sort"), asc))
	})
	public_api.Get("/api/v1/job/:address/:slot", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleGetJob(c.Params("address"), c.Params("slot")))
	})
	public_api.Get("/api/v1/jobs/:address", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleGetJobs(c.Params("address")))
	})
	public_api.Get("/api/v1/jobs/feed", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleGetJobsFeed())
	})
	guest_api.Post("/api/v1/jobs/:address/:signature", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleAddJob(c.Params("address"), c.Params("signature"), c.Body()))
	})
	guest_api.Patch("/api/v1/jobs/:address/:signature", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleUpdateJob(c.Params("address"), c.Params("signature"), c.Body()))
	})
	member_api.Post("/api/v1/jobs/apply/:address/:signature/:recruiter_address/:slot/", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleApplyJob(c.Params("address"), c.Params("signature"), c.Params("recruiter_address"), c.Params("slot"), c.Body()))
	})

	//-----------------//
	//  watchlist      //
	//-----------------//

	member_api.Post("/api/v1/watchlist/:address/:signature", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleAddWatchlist(c.Params("address"), c.Params("signature"), c.Body()))
	})
	member_api.Delete("/api/v1/watchlist/:address/:signature", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleRemoveWatchlist(c.Params("address"), c.Params("signature"), c.Body()))
	})
	public_api.Get("/api/v1/watchlist/:address", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleGetWatchlist(c.Params("address")))
	})

	//-----------------//
	//  favorites      //
	//-----------------//

	member_api.Post("/api/v1/favorites/:address/:signature", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleAddFavorite(c.Params("address"), c.Params("signature"), c.Body()))
	})
	member_api.Delete("/api/v1/favorites/:address/:signature", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleRemoveFavorite(c.Params("address"), c.Params("signature"), c.Body()))
	})
	public_api.Get("/api/v1/favorites/:address", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleGetFavorites(c.Params("address")))
	})

	//-----------------//
	//  conversations  //
	//-----------------//

	public_api.Get("/api/v1/conversations/:address", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleGetConversations(c.Params("address")))
	})
	unknown_api.Post("/api/v1/conversations/:address/:signature", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleAddConversation(c.Params("address"), c.Params("signature"), c.Body()))
	})

	//-----------------//
	//  deals          //
	//-----------------//

	public_api.Get("/api/v1/deals/:recruiter/:creator", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleGetDeals(c.Params("recruiter"), c.Params("creator")))
	})
	// todo: move to unknown_api (needs frontend adjustments)
	guest_api.Post("/api/v1/deals/:recruiter/:creator/:signature", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleAddDeal(c.Params("recruiter"), c.Params("creator"), c.Params("signature"), c.Body()))
	})
	member_api.Patch("/api/v1/deals/:recruiter/:creator/:signature", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleSignDeal(c.Params("recruiter"), c.Params("creator"), c.Params("signature"), c.Body()))
	})
	// todo: remove record
	guest_api.Delete("/api/v1/deals/:recruiter/:creator/:signature", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleExecuteDeal(c.Params("recruiter"), c.Params("creator"), c.Params("signature"), c.Body()))
	})

	//-----------------//
	//  utils          //
	//-----------------//

	// todo: protection?
	app.Post("api/v1/salt/:address", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleAddSalt(c.Params("address")))
	})
	member_api.Get("api/v1/verify/:address/:signature", func(c *fiber.Ctx) error {
		return c.JSON("success")
	})
	public_api.Get("api/v1/tags", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleGetTags())
	})
	// todo: how will tags be added?
	member_api.Post("api/v1/tags/:address/:signature/:tag", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleAddTag(c.Params("address"), c.Params("signature"), c.Params("tag")))
	})
	public_api.Get("/api/v1/config", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleConfig())
	})
	public_api.Get("/api/v1/rating/:address", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleGetRating(c.Params("address")))
	})

	app.Listen(":" + conf.API.Port)
}
