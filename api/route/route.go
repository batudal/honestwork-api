package route

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/takez0o/honestwork-api/api/handler"
	"github.com/takez0o/honestwork-api/api/middleware"
	"github.com/takez0o/honestwork-api/utils/config"
)

func SetRoutes(app *fiber.App, conf *config.Config) {
	api_v1 := app.Group("/api/v1", func(c *fiber.Ctx) error {
		return c.Next()
	})

	//-----------------//
	//  workers        //
	//-----------------//

	// workers_api := app.Group("/api/workers/", func(c *fiber.Ctx) error {
	// 	return c.Next()
	// })
	// workers_api.Patch("jobs/:address/:slot", func(c *fiber.Ctx) error {
	// 	return c.JSON(handler.HandleConsumeJob(c.Params("address"), c.Params("slot"), c.Body()))
	// })
	// workers_api.Patch("rating/:address", func(c *fiber.Ctx) error {
	// 	return c.JSON(handler.HandleUpdateRating(c.Params("address"), c.Body()))
	// })

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
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
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
	api_v1.Get("/skills_member/:address/:signature", func(c *fiber.Ctx) error {
		err := middleware.AuthorizeMember(c.Params("address"), c.Params("signature"))
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
		}
		return c.JSON(handler.HandleGetSkills(c.Params("address")))
	})
	api_v1.Get("/skills_published/:address", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleGetPublishedSkills(c.Params("address")))
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
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
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
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
		}
		return c.JSON(handler.HandleAddJob(c.Params("address"), c.Params("signature"), c.Body()))
	})
	api_v1.Patch("/jobs/:address/:signature", func(c *fiber.Ctx) error {
		err := middleware.AuthorizeGuest(c.Params("address"), c.Params("signature"))
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
		}
		return c.JSON(handler.HandleUpdateJob(c.Params("address"), c.Params("signature"), c.Body()))
	})
	api_v1.Post("/jobs/apply/:address/:signature/:recruiter_address/:slot/", func(c *fiber.Ctx) error {
		err := middleware.AuthorizeMember(c.Params("address"), c.Params("signature"))
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
		}
		return c.JSON(handler.HandleApplyJob(c.Params("address"), c.Params("signature"), c.Params("recruiter_address"), c.Params("slot"), c.Body()))
	})

	//-----------------//
	//  watchlist      //
	//-----------------//

	api_v1.Post("/watchlist/:address/:signature", func(c *fiber.Ctx) error {
		err := middleware.AuthorizeMember(c.Params("address"), c.Params("signature"))
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
		}
		return c.JSON(handler.HandleAddWatchlist(c.Params("address"), c.Params("signature"), c.Body()))
	})
	api_v1.Delete("/watchlist/:address/:signature", func(c *fiber.Ctx) error {
		err := middleware.AuthorizeMember(c.Params("address"), c.Params("signature"))
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
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
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
		}
		return c.JSON(handler.HandleAddFavorite(c.Params("address"), c.Params("signature"), c.Body()))
	})
	api_v1.Delete("/favorites/:address/:signature", func(c *fiber.Ctx) error {
		err := middleware.AuthorizeMember(c.Params("address"), c.Params("signature"))
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
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
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
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
		err := middleware.AuthorizeGuest(c.Params("recruiter"), c.Params("signature"))
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
		}
		return c.JSON(handler.HandleAddDeal(c.Params("recruiter"), c.Params("creator"), c.Params("signature"), c.Body()))
	})
	api_v1.Patch("/deals/:recruiter/:creator/:signature", func(c *fiber.Ctx) error {
		err := middleware.AuthorizeMember(c.Params("creator"), c.Params("signature"))
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
		}
		return c.JSON(handler.HandleSignDeal(c.Params("recruiter"), c.Params("creator"), c.Params("signature"), c.Body()))
	})
	api_v1.Delete("/deals/:recruiter/:creator/:signature", func(c *fiber.Ctx) error {
		err := middleware.AuthorizeGuest(c.Params("recruiter"), c.Params("signature"))
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
		}
		return c.JSON(handler.HandleExecuteDeal(c.Params("recruiter"), c.Params("creator"), c.Params("signature"), c.Body()))
	})

	//-----------------//
	//  utils          //
	//-----------------//

	api_v1.Post("/salt/:address", func(c *fiber.Ctx) error {
		return c.JSON(handler.HandleAddSalt(c.Params("address")))
	})
	api_v1.Get("/verify/:address/:signature", func(c *fiber.Ctx) error {
		err := middleware.AuthorizeUnknown(c.Params("address"), c.Params("signature"))
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
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
