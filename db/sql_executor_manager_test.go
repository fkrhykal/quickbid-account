package db_test

import (
	"context"
	"testing"

	"github.com/fkrhykal/quickbid-account/config"
	"github.com/fkrhykal/quickbid-account/db"
	"github.com/stretchr/testify/assert"
)

func TestSqlExecutorManager(t *testing.T) {
	t.Run("sql executor", func(t *testing.T) {
		pgDB, err := db.NewPostgresDB(config.PostgresTestConfig)
		assert.NoError(t, err)
		defer pgDB.Close()

		execManager := db.NewSqlExecutorManager(pgDB)

		var result int
		err = execManager.Executor().QueryRow("SELECT 1").Scan(&result)
		assert.NoError(t, err)
		assert.Equal(t, 1, result)
	})

	t.Run("sql tx executor", func(t *testing.T) {
		pgDB, err := db.NewPostgresDB(config.PostgresTestConfig)
		assert.NoError(t, err)
		defer pgDB.Close()

		execManager := db.NewSqlExecutorManager(pgDB)
		ctx := context.Background()

		tx, err := execManager.TxExecutor(ctx)
		assert.NoError(t, err)

		err = tx.Commit()
		assert.NoError(t, err)
	})
}
