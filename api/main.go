package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"

	"github.com/takez0o/honestwork-api/api/handler"
	"github.com/takez0o/honestwork-api/api/middleware"
	"github.com/takez0o/honestwork-api/api/worker"
	"github.com/takez0o/honestwork-api/utils/config"
)

func main() {
	// config/env setup
	conf, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()
	app.Static("/", "../static")

	// core middleware
	client_key := os.Getenv("CLIENT_KEY")
	client_password := os.Getenv("CLIENT_PASSWORD")
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			client_key: client_password,
		},
	}))
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(recover.New())

	// start workers
	rating_watcher := worker.NewRatingWatcher()
	go rating_watcher.WatchRatings()
	revenue_watcher := worker.NewRevenueWatcher()
	go revenue_watcher.WatchRevenues()
	// deal_watcher := worker.NewDealWatcher()
	// go deal_watcher.WatchDeals()

	api_v1 := app.Group("/api/v1", func(c *fiber.Ctx) error {
		return c.Next()
	})

	//-----------------//
	//  users          //
	//-----------------//

	api_v1.Post("/users/:address/:signature", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleSignup(c.Params("address"), c.Params("signature")))
	})
	api_v1.Get("/users/:address", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleGetUser(c.Params("address")))
	})
	api_v1.Patch("/users/:address/:signature", func(c *fiber.Ctx) error {
		err := middleware.AuthorizeMember(c.Params("address"), c.Params("signature"))
		if err != nil {
			return c.JSON(err)
		}
		return c.JSON(handler.HandleUserUpdate(c.Params("address"), c.Params("signature"), c.Body()))
	})

	//-----------------//
	//  skills         //
	//-----------------//

	api_v1.Get("/skills/total", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleGetSkillsTotal())
	})
	api_v1.Get("/skills/limit/:offset/:size", func(c *fiber.Ctx) error {
		offset, _ := strconv.Atoi(c.Params("offset"))
		size, _ := strconv.Atoi(c.Params("size"))
		return c.JSON(handler.HandleGetSkillsLimit(offset, size))
	})
	api_v1.Get("/skills/:sort/:order", func(c *fiber.Ctx) error {
		asc, _ := strconv.ParseBool(c.Params("order"))
		return c.JSON(handler.HandleGetAllSkills(c.Params("sort"), asc))
	})
	api_v1.Get("/skills/:address", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleGetSkills(c.Params("address")))
	})
	api_v1.Get("/skill/:address/:slot", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleGetSkill(c.Params("address"), c.Params("slot")))
	})
	api_v1.Post("/skills/:address/:signature", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleAddSkill(c.Params("address"), c.Params("signature"), c.Body()))
	})
	api_v1.Patch("/skills/:address/:signature/:slot", func(c *fiber.Ctx) error {
		err := middleware.AuthorizeMember(c.Params("address"), c.Params("signature"))
		if err != nil {
			return c.JSON(err)
		}
		return c.JSON(handler.HandleUpdateSkill(c.Params("address"), c.Params("signature"), c.Params("slot"), c.Body()))
	})

	//-----------------//
	//  jobs           //
	//-----------------//

	api_v1.Get("/jobs/total", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleGetJobsTotal())
	})
	api_v1.Get("/jobs/limit/:offset/:size", func(c *fiber.Ctx) error {
		offset, _ := strconv.Atoi(c.Params("offset"))
		size, _ := strconv.Atoi(c.Params("size"))
		return c.JSON(handler.HandleGetJobsLimit(offset, size))
	})
	api_v1.Get("/jobs/:sort/:order", func(c *fiber.Ctx) error {
		asc, _ := strconv.ParseBool(c.Params("order"))
		return c.JSON(handler.HandleGetAllJobs(c.Params("sort"), asc))
	})
	api_v1.Get("/job/:address/:slot", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleGetJob(c.Params("address"), c.Params("slot")))
	})
	api_v1.Get("/jobs/:address", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleGetJobs(c.Params("address")))
	})
	api_v1.Get("/jobs/feed", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleGetJobsFeed())
	})
	api_v1.Post("/jobs/:address/:signature", func(c *fiber.Ctx) error {
		err := middleware.AuthorizeGuest(c.Params("address"), c.Params("signature"))
		if err != nil {
			return c.JSON(err)
		}
		return c.JSON(handler.HandleAddJob(c.Params("address"), c.Params("signature"), c.Body()))
	})
	api_v1.Patch("/jobs/:address/:signature", func(c *fiber.Ctx) error {
		err := middleware.AuthorizeGuest(c.Params("address"), c.Params("signature"))
		if err != nil {
			return c.JSON(err)
		}
		return c.JSON(handler.HandleUpdateJob(c.Params("address"), c.Params("signature"), c.Body()))
	})
	api_v1.Post("/jobs/apply/:address/:signature/:recruiter_address/:slot/", func(c *fiber.Ctx) error {
		err := middleware.AuthorizeMember(c.Params("address"), c.Params("signature"))
		if err != nil {
			return c.JSON(err)
		}
		return c.JSON(handler.HandleApplyJob(c.Params("address"), c.Params("signature"), c.Params("recruiter_address"), c.Params("slot"), c.Body()))
	})

	//-----------------//
	//  watchlist      //
	//-----------------//

	api_v1.Post("/watchlist/:address/:signature", func(c *fiber.Ctx) error {
		err := middleware.AuthorizeMember(c.Params("address"), c.Params("signature"))
		if err != nil {
			return c.JSON(err)
		}
		return c.JSON(handler.HandleAddWatchlist(c.Params("address"), c.Params("signature"), c.Body()))
	})
	api_v1.Delete("/watchlist/:address/:signature", func(c *fiber.Ctx) error {
		err := middleware.AuthorizeMember(c.Params("address"), c.Params("signature"))
		if err != nil {
			return c.JSON(err)
		}
		return c.JSON(handler.HandleRemoveWatchlist(c.Params("address"), c.Params("signature"), c.Body()))
	})
	api_v1.Get("/watchlist/:address", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleGetWatchlist(c.Params("address")))
	})

	//-----------------//
	//  favorites      //
	//-----------------//

	api_v1.Post("/favorites/:address/:signature", func(c *fiber.Ctx) error {
		err := middleware.AuthorizeMember(c.Params("address"), c.Params("signature"))
		if err != nil {
			return c.JSON(err)
		}
		return c.JSON(handler.HandleAddFavorite(c.Params("address"), c.Params("signature"), c.Body()))
	})
	api_v1.Delete("/favorites/:address/:signature", func(c *fiber.Ctx) error {
		err := middleware.AuthorizeMember(c.Params("address"), c.Params("signature"))
		if err != nil {
			return c.JSON(err)
		}
		return c.JSON(handler.HandleRemoveFavorite(c.Params("address"), c.Params("signature"), c.Body()))
	})
	api_v1.Get("/favorites/:address", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleGetFavorites(c.Params("address")))
	})

	//-----------------//
	//  conversations  //
	//-----------------//

	api_v1.Get("/conversations/:address", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleGetConversations(c.Params("address")))
	})
	api_v1.Post("/conversations/:address/:signature", func(c *fiber.Ctx) error {
		err := middleware.AuthorizeUnknown(c.Params("address"), c.Params("signature"))
		if err != nil {
			return c.JSON(err)
		}
		return c.JSON(handler.HandleAddConversation(c.Params("address"), c.Params("signature"), c.Body()))
	})

	//-----------------//
	//  deals          //
	//-----------------//

	api_v1.Get("/deals/:recruiter/:creator", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleGetDeals(c.Params("recruiter"), c.Params("creator")))
	})
	api_v1.Post("/deals/:recruiter/:creator/:signature", func(c *fiber.Ctx) error {
		fmt.Println("post deal")
		err := middleware.AuthorizeUnknown(c.Params("address"), c.Params("signature"))
		if err != nil {
			return c.JSON(err)
		}
		fmt.Println("authorized")
		return c.JSON(handler.HandleAddDeal(c.Params("recruiter"), c.Params("creator"), c.Params("signature"), c.Body()))
	})
	api_v1.Patch("/deals/:recruiter/:creator/:signature", func(c *fiber.Ctx) error {
		err := middleware.AuthorizeMember(c.Params("address"), c.Params("signature"))
		if err != nil {
			return c.JSON(err)
		}
		return c.JSON(handler.HandleSignDeal(c.Params("recruiter"), c.Params("creator"), c.Params("signature"), c.Body()))
	})
	// todo: remove record
	api_v1.Delete("/deals/:recruiter/:creator/:signature", func(c *fiber.Ctx) error {
		err := middleware.AuthorizeGuest(c.Params("address"), c.Params("signature"))
		if err != nil {
			return c.JSON(err)
		}
		return c.JSON(handler.HandleExecuteDeal(c.Params("recruiter"), c.Params("creator"), c.Params("signature"), c.Body()))
	})

	//-----------------//
	//  utils          //
	//-----------------//

	// todo: protection?
	api_v1.Post("/salt/:address", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleAddSalt(c.Params("address")))
	})
	api_v1.Get("/verify/:address/:signature", func(c *fiber.Ctx) error {
		err := middleware.AuthorizeUnknown(c.Params("address"), c.Params("signature"))
		if err != nil {
			return c.JSON(err)
		}
		return c.JSON("success")
	})
	api_v1.Get("/tags", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleGetTags())
	})
	// todo: how will tags be added?
	api_v1.Post("/tags/:address/:signature/:tag", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleAddTag(c.Params("address"), c.Params("signature"), c.Params("tag")))
	})
	api_v1.Get("/config", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleConfig())
	})
	api_v1.Get("/rating/:address", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleGetRating(c.Params("address")))
	})
	api_v1.Get("/whitelist", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleWhitelist())
	})

	app.Listen(":" + conf.API.Port)
}
