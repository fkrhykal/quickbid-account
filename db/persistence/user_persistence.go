package persistence

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/fkrhykal/quickbid-account/db"
	"github.com/fkrhykal/quickbid-account/internal/entity"
	"github.com/fkrhykal/quickbid-account/internal/repository"
	"github.com/google/uuid"
)

func PgSaveUser(log *slog.Logger) repository.SaveUser[db.SqlExecutor] {
	return func(ctx context.Context, executor db.SqlExecutor, user *entity.User) error {
		log.DebugContext(ctx, "Saving new user")
		query := `INSERT INTO users(id, username, password, avatar) VALUES($1, $2, $3, $4)`
		_, err := executor.ExecContext(ctx, query, user.ID, user.Username, user.Password, user.Avatar)
		if err != nil {
			log.DebugContext(ctx, "Failed to save user", slog.Any("error", err))
			return err
		}
		log.DebugContext(ctx, "User saved successfully")
		return nil
	}
}

func PgFindUserByUsername(log *slog.Logger) repository.FindUserByUsername[db.SqlExecutor] {
	return func(ctx context.Context, executor db.SqlExecutor, username string) (*entity.User, error) {
		log.DebugContext(ctx, "Finding user by username")
		user := new(entity.User)
		query := `SELECT id, username, password, avatar FROM users WHERE username = $1`
		err := executor.QueryRowContext(ctx, query, username).Scan(&user.ID, &user.Username, &user.Password, &user.Avatar)
		if err == nil {
			log.DebugContext(ctx, "User found")
			return user, nil
		}
		if errors.Is(err, sql.ErrNoRows) {
			log.DebugContext(ctx, "User not found")
			return nil, nil
		}
		log.DebugContext(ctx, "Error finding user", slog.Any("error", err))
		return nil, err
	}
}

func PgFindUserByID(log *slog.Logger) repository.FindUserByID[db.SqlExecutor] {
	return func(ctx context.Context, executor db.SqlExecutor, ID uuid.UUID) (*entity.User, error) {
		log.DebugContext(ctx, "Finding user by id")
		user := new(entity.User)
		query := `SELECT id, username, password, avatar FROM users WHERE id = $1`
		err := executor.QueryRowContext(ctx, query, ID).Scan(&user.ID, &user.Username, &user.Password, &user.Avatar)
		if err == nil {
			log.DebugContext(ctx, "User found")
			return user, nil
		}
		if errors.Is(err, sql.ErrNoRows) {
			log.DebugContext(ctx, "User not found")
			return nil, nil
		}
		log.DebugContext(ctx, "Error finding user", slog.Any("error", err))
		return nil, err
	}
}
