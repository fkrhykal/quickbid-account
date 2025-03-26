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
	"github.com/fkrhykal/quickbid-account/internal/credential"
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

	var postgresConfig *db.PostgresConfig
	var credentialConfig *credential.JwtCredentialManagerConfig

	if env := os.Getenv("ENV"); env == "dev" {
		postgresConfig = config.PostgresDevConfig
		credentialConfig = config.JwtDevConfig
	} else {
		postgresConfig = config.PostgresProdConfig(logger)
		credentialConfig = config.JwtProdConfig(logger)
	}

	pgDB, err := db.SetupPostgresDB(postgresConfig)
	if err != nil {
		log.Fatal(err)
	}

	app.Bootstrap(&app.BootstrapConfig{
		Fiber:      fiber.Group("/api/v1"),
		DB:         pgDB,
		Logger:     logger,
		Credential: credentialConfig,
	})

	address := fmt.Sprintf(":%d", config.EnvInt("APP_PORT", 8000))

	if err := fiber.Listen(address); err != nil {
		log.Fatal(err)
	}
}
