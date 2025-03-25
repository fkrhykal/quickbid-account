package service

import (
	"context"
	"log/slog"

	"github.com/fkrhykal/quickbid-account/internal/credential"
	"github.com/fkrhykal/quickbid-account/internal/data"
	"github.com/fkrhykal/quickbid-account/internal/dto"
	"github.com/fkrhykal/quickbid-account/internal/entity"
	"github.com/fkrhykal/quickbid-account/internal/repository"
	"github.com/fkrhykal/quickbid-account/internal/usecase"
	"github.com/fkrhykal/quickbid-account/internal/validation"
	"github.com/google/uuid"
)

func SignUpService[T any](
	log *slog.Logger,
	validate validation.Validator[*dto.SignUpRequest],
	execManager data.ExecutorManager[T],
	saveUser repository.SaveUser[T],
	findUserByUsername repository.FindUserByUsername[T],
	passwordHasher credential.PasswordHasher,
) usecase.SignUp {
	return func(ctx context.Context, req *dto.SignUpRequest) (*dto.SignUpResponse, error) {
		log.DebugContext(ctx, "Starting sign-up process")

		if err := validate(ctx, req); err != nil {
			log.DebugContext(ctx, "Validation failed", slog.Any("error", err))
			return nil, err
		}
		log.DebugContext(ctx, "Validation passed")

		executor := execManager.Executor()
		log.DebugContext(ctx, "Database executor acquired")

		existedUser, err := findUserByUsername(ctx, executor, req.Username)
		if err != nil {
			log.DebugContext(ctx, "Error finding user by username", slog.Any("error", err))
			return nil, err
		}
		if existedUser != nil {
			log.DebugContext(ctx, "Username already exists")
			return nil, usecase.ErrUsernameAlreadyUsed
		}
		log.DebugContext(ctx, "Username is available, proceeding with user creation")

		hashedPassword, err := passwordHasher.Hash(req.Password)
		if err != nil {
			log.DebugContext(ctx, "Password hashing failed", slog.Any("error", err))
			return nil, err
		}
		log.DebugContext(ctx, "password hashing success")

		user := &entity.User{
			ID:       uuid.New(),
			Username: req.Username,
			Password: hashedPassword,
		}

		if err := saveUser(ctx, executor, user); err != nil {
			log.DebugContext(ctx, "Failed to save user", slog.Any("error", err))
			return nil, err
		}
		log.DebugContext(ctx, "User successfully saved")

		log.DebugContext(ctx, "Sign-up process completed")
		return &dto.SignUpResponse{ID: user.ID}, nil
	}
}
