package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/fkrhykal/quickbid-account/db/migration"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	SSLMode  string
	Database string
	MaxCon   int
	MinCon   int
	Logger   *slog.Logger
}

func NewPostgresDB(config *PostgresConfig) (*sql.DB, error) {
	config.Logger.Debug("Initializing PostgreSQL connection")

	stringConfig := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.Database, config.SSLMode)

	db, err := sql.Open("pgx", stringConfig)
	if err != nil {
		config.Logger.Error("Failed to open database connection", slog.Any("error", err))
		return nil, err
	}

	db.SetMaxOpenConns(config.MaxCon)
	db.SetMaxIdleConns(config.MinCon)
	db.SetConnMaxLifetime(90 * time.Minute)

	if err := db.Ping(); err != nil {
		config.Logger.Error("Database ping failed", slog.Any("error", err))
		db.Close()
		return nil, err
	}

	config.Logger.Debug("PostgreSQL connection established successfully")
	return db, nil
}

func MigratePostgresDB(db *sql.DB, log *slog.Logger) error {
	log.Debug("Starting database migration")

	sourceDriver, err := iofs.New(migration.MigrationFs, ".")
	if err != nil {
		log.Error("Failed to initialize migration source", slog.Any("error", err))
		return err
	}
	databaseInstance, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Error("Failed to create database instance for migration", slog.Any("error", err))
		db.Close()
		return err
	}

	m, err := migrate.NewWithInstance("iofs", sourceDriver, "pgx", databaseInstance)
	if err != nil {
		log.Error("Failed to initialize migration instance", slog.Any("error", err))
		databaseInstance.Close()
		return err
	}

	err = m.Up()
	if err == nil {
		log.Debug("Migration applied successfully")
		return nil
	}
	if errors.Is(err, migrate.ErrNoChange) {
		log.Debug("No new migrations to apply")
		return nil
	}

	log.Error("Migration failed", slog.Any("error", err))
	return err
}

func SetupPostgresDB(config *PostgresConfig) (*sql.DB, error) {
	config.Logger.Info("Setting up PostgreSQL database")

	pgDB, err := NewPostgresDB(config)
	if err != nil {
		config.Logger.Error("Failed to initialize database", slog.Any("error", err))
		return nil, err
	}

	err = MigratePostgresDB(pgDB, config.Logger)
	if err != nil {
		config.Logger.Error("Database migration failed", slog.Any("error", err))
		return nil, err
	}

	config.Logger.Info("PostgreSQL setup completed successfully")
	return pgDB, nil
}
