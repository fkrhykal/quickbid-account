package data

import "context"

type ExecutorManager[T any] interface {
	Executor() T
	TxExecutor(ctx context.Context) (TxExecutor[T], error)
}

type TxExecutor[T any] interface {
	Executor() T
	Rollback() error
	Commit() error
}
