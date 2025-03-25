package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/fkrhykal/quickbid-account/api"
	"github.com/fkrhykal/quickbid-account/app"
	"github.com/fkrhykal/quickbid-account/config"
	"github.com/fkrhykal/quickbid-account/db"
)

func PostgresWithLogger(logger *slog.Logger) func(config *db.PostgresConfig) *db.PostgresConfig {
	return func(config *db.PostgresConfig) *db.PostgresConfig {
		config.Logger = logger
		return config
	}
}

func main() {
	logger := slog.Default()

	fiber := api.NewFiberApp(logger)

	postgresConfig := db.PostgresDevConfig

	if env := os.Getenv("ENV"); env == "production" {
		postgresConfig = db.PostgresProdConfig(logger)
	}

	pgDB, err := db.SetupPostgresDB(config.Configure(
		postgresConfig,
		PostgresWithLogger(logger),
	))

	if err != nil {
		log.Fatal(err)
	}

	app.Bootstrap(&app.BootstrapConfig{
		Fiber:  fiber.Group("/api/v1"),
		DB:     pgDB,
		Logger: logger,
	})

	address := fmt.Sprintf(":%d", config.EnvInt("APP_PORT", 8000))

	if err := fiber.Listen(address); err != nil {
		log.Fatal(err)
	}
}
