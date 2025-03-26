package db_test

import (
	"context"
	"log/slog"
	"testing"

	"github.com/fkrhykal/quickbid-account/config"
	"github.com/fkrhykal/quickbid-account/db"
	"github.com/stretchr/testify/assert"
)

func TestNewPostgresDB(t *testing.T) {
	pgDB, err := db.NewPostgresDB(config.PostgresTestConfig)
	assert.NoError(t, err)
	defer pgDB.Close()

	ctx := context.Background()

	var result int
	err = pgDB.QueryRowContext(ctx, `SELECT 1`).Scan(&result)
	assert.NoError(t, err)
	assert.Equal(t, 1, result)
}

func TestMigrationPostgresDB(t *testing.T) {
	pgDB, err := db.NewPostgresDB(config.PostgresTestConfig)
	assert.NoError(t, err)
	defer pgDB.Close()

	err = db.MigratePostgresDB(pgDB, slog.Default())
	assert.NoError(t, err)
}
