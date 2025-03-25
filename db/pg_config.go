package db

import (
	"log/slog"

	"github.com/fkrhykal/quickbid-account/config"
)

var PostgresTestConfig = &PostgresConfig{
	Host:     "localhost",
	Port:     5431,
	User:     "quickbid",
	Password: "secret",
	Database: "account_db",
	SSLMode:  "disable",
	MaxCon:   100,
	MinCon:   10,
	Logger:   slog.Default(),
}

var PostgresDevConfig = &PostgresConfig{
	Host:     "localhost",
	Port:     5430,
	User:     "quickbid",
	Password: "secret",
	Database: "account_db",
	SSLMode:  "disable",
	MaxCon:   100,
	MinCon:   10,
	Logger:   slog.Default(),
}

var PostgresProdConfig = func(log *slog.Logger) *PostgresConfig {
	return &PostgresConfig{
		Host:     config.MustEnvString("POSTGRES_DB"),
		Port:     config.MustEnvInt("POSTGRES_PORT"),
		User:     config.MustEnvString("POSTGRES_USER"),
		Password: config.MustEnvString("POSTGRES_PASSWORD"),
		SSLMode:  "disable",
		Database: config.MustEnvString("POSTGRES_DATABASE"),
		MaxCon:   100,
		MinCon:   10,
		Logger:   log,
	}
}
