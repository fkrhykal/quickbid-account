package api

import (
	"log/slog"

	"github.com/fkrhykal/quickbid-account/api/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
)

func NewFiberApp(log *slog.Logger) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: handler.ErrorHandler(log),
	})

	app.Use(etag.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	app.Use(healthcheck.New())

	return app
}
