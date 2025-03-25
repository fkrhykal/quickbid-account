package service_test

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"testing"

	"github.com/fkrhykal/quickbid-account/internal/data"
	"github.com/fkrhykal/quickbid-account/internal/dto"
	"github.com/fkrhykal/quickbid-account/internal/entity"
	"github.com/fkrhykal/quickbid-account/internal/service"
	"github.com/fkrhykal/quickbid-account/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))

	defaultValidate := func(ctx context.Context, v *dto.SignUpRequest) error {
		return nil
	}
	defaultSave := func(ctx context.Context, executor any, user *entity.User) error {
		return nil
	}
	defaultFindByUsername := func(ctx context.Context, executor any, username string) (*entity.User, error) {
		return nil, nil
	}

	t.Run("case sign up success", func(t *testing.T) {
		ctx := context.Background()

		execManager := data.NewMockExecutorManager[any](t)
		signUp := service.SignUpService(log, defaultValidate, execManager, defaultSave, defaultFindByUsername)

		execManager.EXPECT().Executor().Return(nil)

		_, err := signUp(ctx, &dto.SignUpRequest{})
		assert.NoError(t, err)
	})

	t.Run("case validation failed", func(t *testing.T) {
		ctx := context.Background()
		validationError := errors.New("request invalid")

		validationFailed := func(ctx context.Context, req *dto.SignUpRequest) error {
			return validationError
		}

		execManager := data.NewMockExecutorManager[any](t)
		signUp := service.SignUpService(log, validationFailed, execManager, defaultSave, defaultFindByUsername)

		_, err := signUp(ctx, &dto.SignUpRequest{})
		assert.ErrorIs(t, validationError, err)
	})

	t.Run("case username used", func(t *testing.T) {
		ctx := context.Background()

		usernameUsed := func(ctx context.Context, exec any, username string) (*entity.User, error) {
			return new(entity.User), nil
		}

		execManager := data.NewMockExecutorManager[any](t)
		signUp := service.SignUpService(log, defaultValidate, execManager, defaultSave, usernameUsed)

		execManager.EXPECT().Executor().Return(nil)

		_, err := signUp(ctx, &dto.SignUpRequest{})
		assert.ErrorIs(t, usecase.ErrUsernameAlreadyUsed, err)

	})

	t.Run("case saving user failed", func(t *testing.T) {
		ctx := context.Background()
		saveError := errors.New("failed to save user")

		saveFailed := func(ctx context.Context, executor any, user *entity.User) error {
			return saveError
		}

		execManager := data.NewMockExecutorManager[any](t)
		signUp := service.SignUpService(log, defaultValidate, execManager, saveFailed, defaultFindByUsername)

		execManager.EXPECT().Executor().Return(nil)

		_, err := signUp(ctx, &dto.SignUpRequest{})
		assert.ErrorIs(t, saveError, err)
	})
}
