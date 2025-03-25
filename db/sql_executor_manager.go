package db

import (
	"context"
	"database/sql"

	"github.com/fkrhykal/quickbid-account/internal/data"
)

type SqlExecutorManager struct {
	db *sql.DB
}

func (m *SqlExecutorManager) Executor() SqlExecutor {
	return m.db
}

func (m *SqlExecutorManager) TxExecutor(ctx context.Context) (data.TxExecutor[SqlExecutor], error) {
	tx, err := m.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	return &SqlTxExecutor{executor: tx}, nil
}

func NewSqlExecutorManager(db *sql.DB) data.ExecutorManager[SqlExecutor] {
	return &SqlExecutorManager{
		db: db,
	}
}

type SqlTxExecutor struct {
	executor *sql.Tx
}

func (t *SqlTxExecutor) Executor() SqlExecutor {
	return t.executor
}

func (t *SqlTxExecutor) Commit() error {
	return t.executor.Commit()
}

func (t *SqlTxExecutor) Rollback() error {
	return t.executor.Rollback()
}
