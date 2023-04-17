package route

import (
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
	//  users          //
	//-----------------//

	api_v1.Post("/users/:address/:signature", handler.HandleSignup())
	api_v1.Get("/users/:address", handler.HandleGetUser())
	api_v1.Patch("/users/:address/:signature", middleware.AuthorizeMember(), handler.HandleUserUpdate())

	//-----------------//
	//  skills         //
	//-----------------//

	api_v1.Get("/skills/total", handler.HandleGetSkillsTotal())
	api_v1.Get("/skills/limit/:offset/:size", handler.HandleGetSkillsLimit())
	api_v1.Get("/skills/:sort/:order", handler.HandleGetAllSkills())
	api_v1.Get("/skills_member/:address/:signature", middleware.AuthorizeMember(), handler.HandleGetSkills())
	api_v1.Get("/skills_published/:address", handler.HandleGetPublishedSkills())
	api_v1.Get("/skill/:address/:slot", handler.HandleGetSkill())
	api_v1.Post("/skills/:address/:signature", middleware.AuthorizeMember(), handler.HandleAddSkill())
	api_v1.Patch("/skills/:address/:signature/:slot", middleware.AuthorizeMember(), handler.HandleUpdateSkill())

	//-----------------//
	//  jobs           //
	//-----------------//

	api_v1.Get("/jobs/total", handler.HandleGetJobsTotal())
	api_v1.Get("/jobs/limit/:offset/:size", handler.HandleGetJobsLimit())
	api_v1.Get("/jobs/:sort/:order", handler.HandleGetAllJobs())
	api_v1.Get("/job/:address/:slot", handler.HandleGetJob())
	api_v1.Get("/jobs/:address", handler.HandleGetJobs())
	api_v1.Get("/jobs/feed", handler.HandleGetJobsFeed())
	api_v1.Post("/jobs/:address/:signature", middleware.AuthorizeGuest(), handler.HandleAddJob())
	api_v1.Patch("/jobs/:address/:signature", middleware.AuthorizeGuest(), handler.HandleUpdateJob())
	api_v1.Post("/jobs/apply/:address/:signature/:recruiter_address/:slot/", middleware.AuthorizeMember(), handler.HandleApplyJob())

	//-----------------//
	//  watchlist      //
	//-----------------//

	api_v1.Post("/watchlist/:address/:signature", middleware.AuthorizeMember(), handler.HandleAddWatchlist())
	api_v1.Delete("/watchlist/:address/:signature", middleware.AuthorizeMember(), handler.HandleRemoveWatchlist())
	api_v1.Get("/watchlist/:address", handler.HandleGetWatchlist())

	//-----------------//
	//  favorites      //
	//-----------------//

	api_v1.Post("/favorites/:address/:signature", middleware.AuthorizeMember(), handler.HandleAddFavorite())
	api_v1.Delete("/favorites/:address/:signature", middleware.AuthorizeMember(), handler.HandleRemoveFavorite())
	api_v1.Get("/favorites/:address", handler.HandleGetFavorites())

	//-----------------//
	//  conversations  //
	//-----------------//

	api_v1.Get("/conversations/:address", handler.HandleGetConversations())
	api_v1.Post("/conversations/:address/:signature", middleware.AuthorizeUnknown(), handler.HandleAddConversation())

	//-----------------//
	//  deals          //
	//-----------------//

	api_v1.Get("/deals/:recruiter/:creator", handler.HandleGetDeals())
	api_v1.Post("/deals/:recruiter/:creator/:signature", middleware.AuthorizeGuest(), handler.HandleAddDeal())
	api_v1.Patch("/deals/:recruiter/:creator/:signature", middleware.AuthorizeMember(), handler.HandleSignDeal())
	api_v1.Delete("/deals/:recruiter/:creator/:signature", middleware.AuthorizeGuest(), handler.HandleExecuteDeal())

	//-----------------//
	//  utils          //
	//-----------------//

	api_v1.Post("/salt/:address", handler.HandleAddSalt())
	api_v1.Get("/tags", handler.HandleGetTags())
	api_v1.Post("/tags/:address/:signature/:tag", middleware.AuthorizeMember(), handler.HandleAddTag())
	api_v1.Get("/config", handler.HandleConfig())
	api_v1.Get("/rating/:address", handler.HandleGetRating())
	api_v1.Get("/whitelist", handler.HandleWhitelist())
	api_v1.Get("/verify/:address/:signature", middleware.AuthorizeUnknown(), func(c *fiber.Ctx) error {
		return c.JSON("success")
	})

	app.Listen(":" + conf.API.Port)
}
