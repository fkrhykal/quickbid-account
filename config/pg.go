package config

import (
	"log/slog"
	"os"

	"github.com/fkrhykal/quickbid-account/db"
)

var PostgresTestConfig = &db.PostgresConfig{
	Host:     "localhost",
	Port:     5431,
	User:     "quickbid",
	Password: "secret",
	Database: "account_db",
	SSLMode:  "disable",
	MaxCon:   100,
	MinCon:   10,
	Logger: slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})),
}

var PostgresDevConfig = &db.PostgresConfig{
	Host:     "localhost",
	Port:     5430,
	User:     "quickbid",
	Password: "secret",
	Database: "account_db",
	SSLMode:  "disable",
	MaxCon:   100,
	MinCon:   10,
	Logger: slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})),
}

func PostgresProdConfig(log *slog.Logger) *db.PostgresConfig {
	return &db.PostgresConfig{
		Host:     MustEnvString("POSTGRES_DB"),
		Port:     MustEnvInt("POSTGRES_PORT"),
		User:     MustEnvString("POSTGRES_USER"),
		Password: MustEnvString("POSTGRES_PASSWORD"),
		SSLMode:  "disable",
		Database: MustEnvString("POSTGRES_DATABASE"),
		MaxCon:   100,
		MinCon:   10,
		Logger:   log,
	}
}
