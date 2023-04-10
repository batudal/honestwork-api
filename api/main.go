package main

import (
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/contrib/fibersentry"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joho/godotenv"

	"github.com/takez0o/honestwork-api/utils/config"
)

func main() {

	//-----------------//
	//  load           //
	//-----------------//

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

	//-----------------//
	//  middleware     //
	//-----------------//

	client_key := os.Getenv("CLIENT_KEY")
	client_password := os.Getenv("CLIENT_PASSWORD")
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			client_key: client_password,
		},
	}))
	app.Use(requestid.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("CLIENT_DOMAIN"),
	}))
	app.Use(recover.New())
	app.Use(limiter.New(limiter.Config{
		Max:               50,
		Expiration:        1 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))
	_ = sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_DSN"),
		Debug:            true,
		AttachStacktrace: true,
	})
	app.Use(fibersentry.New(fibersentry.Config{
		Repanic:         true,
		WaitForDelivery: false,
	}))

	startWorkers()
	setRoutes(app, conf)
}
