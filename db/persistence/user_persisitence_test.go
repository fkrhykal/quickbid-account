package persistence_test

import (
	"context"
	"testing"

	"github.com/fkrhykal/quickbid-account/config"
	"github.com/fkrhykal/quickbid-account/db"
	"github.com/fkrhykal/quickbid-account/db/persistence"
	"github.com/fkrhykal/quickbid-account/internal/entity"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSaveUser(t *testing.T) {
	ctx := context.Background()

	pgDB, err := db.SetupPostgresDB(config.PostgresTestConfig)
	assert.NoError(t, err)

	defer func() {
		pgDB.ExecContext(ctx, "TRUNCATE TABLE users")
		pgDB.Close()
	}()

	execManager := db.NewSqlExecutorManager(pgDB)

	saveUser := persistence.PgSaveUser(config.PostgresTestConfig.Logger)

	user := &entity.User{
		ID:       uuid.New(),
		Username: faker.Username(),
		Password: faker.Password(),
	}

	err = saveUser(ctx, execManager.Executor(), user)
	assert.NoError(t, err)
}

func TestFindUserByUsername(t *testing.T) {
	ctx := context.Background()

	pgDB, err := db.SetupPostgresDB(config.PostgresTestConfig)
	assert.NoError(t, err)

	defer func() {
		pgDB.ExecContext(ctx, "TRUNCATE TABLE users")
		pgDB.Close()
	}()

	execManager := db.NewSqlExecutorManager(pgDB)

	saveUser := persistence.PgSaveUser(config.PostgresTestConfig.Logger)

	user := &entity.User{
		ID:       uuid.New(),
		Username: faker.Username(),
		Password: faker.Password(),
	}

	err = saveUser(ctx, execManager.Executor(), user)
	assert.NoError(t, err)

	findByUsername := persistence.PgFindUserByUsername(config.PostgresTestConfig.Logger)

	savedUser, err := findByUsername(ctx, execManager.Executor(), user.Username)
	assert.NoError(t, err)
	assert.EqualValues(t, user, savedUser)
}
