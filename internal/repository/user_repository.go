package repository

import (
	"context"

	"github.com/fkrhykal/quickbid-account/internal/entity"
)

type FindUserByUsername[T any] func(ctx context.Context, executor T, username string) (*entity.User, error)
type SaveUser[T any] func(ctx context.Context, executor T, user *entity.User) error
