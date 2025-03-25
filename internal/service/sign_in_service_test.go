package service_test

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"testing"

	"github.com/fkrhykal/quickbid-account/internal/credential"
	"github.com/fkrhykal/quickbid-account/internal/data"
	"github.com/fkrhykal/quickbid-account/internal/dto"
	"github.com/fkrhykal/quickbid-account/internal/entity"
	"github.com/fkrhykal/quickbid-account/internal/repository"
	"github.com/fkrhykal/quickbid-account/internal/service"
	"github.com/fkrhykal/quickbid-account/internal/usecase"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSignIn(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))
	validationSuccessCase := func(ctx context.Context, req *dto.SignInRequest) error {
		return nil
	}
	userExistCase := func(user *entity.User) repository.FindUserByUsername[any] {
		return func(ctx context.Context, executor any, username string) (*entity.User, error) {
			return user, nil
		}
	}
	userNotFoundCase := func(ctx context.Context, executor any, username string) (*entity.User, error) {
		return nil, nil
	}

	t.Run("sign in success", func(t *testing.T) {
		ctx := context.Background()

		execManager := data.NewMockExecutorManager[any](t)
		passwordVerifier := credential.NewMockPasswordVerifier(t)
		credentialCreator := credential.NewMockCredentialTokenCreator(t)

		user := &entity.User{
			ID:       uuid.New(),
			Username: faker.Username(),
			Password: faker.Password(),
		}

		signIn := service.SignInService(
			logger,
			validationSuccessCase,
			execManager,
			userExistCase(user),
			passwordVerifier,
			credentialCreator,
		)

		execManager.EXPECT().Executor().Return(nil)
		passwordVerifier.EXPECT().Verify(user.Password, user.Password).Return(nil)
		credentialCreator.EXPECT().
			CreateCredentialToken(ctx, &credential.UserCredential{ID: user.ID}).
			Return(user.ID.String(), nil)

		res, err := signIn(ctx, &dto.SignInRequest{
			Username: user.Username,
			Password: user.Password,
		})
		assert.NoError(t, err)
		assert.EqualValues(t, user.ID, res.ID)
		assert.EqualValues(t, user.ID.String(), res.BearerToken)
	})

	t.Run("validation fail", func(t *testing.T) {
		ctx := context.Background()

		execManager := data.NewMockExecutorManager[any](t)
		passwordVerifier := credential.NewMockPasswordVerifier(t)
		credentialCreator := credential.NewMockCredentialTokenCreator(t)

		user := &entity.User{
			ID:       uuid.New(),
			Username: faker.Username(),
			Password: faker.Password(),
		}

		validationFail := func(ctx context.Context, req *dto.SignInRequest) error {
			return errors.New("validation fail")
		}

		signIn := service.SignInService(
			logger,
			validationFail,
			execManager,
			userNotFoundCase,
			passwordVerifier,
			credentialCreator,
		)

		res, err := signIn(ctx, &dto.SignInRequest{
			Username: user.Username,
			Password: user.Password,
		})
		assert.ErrorIs(t, err, usecase.ErrAuthentication)
		assert.Nil(t, res)
	})

	t.Run("user not found", func(t *testing.T) {
		ctx := context.Background()

		execManager := data.NewMockExecutorManager[any](t)
		passwordVerifier := credential.NewMockPasswordVerifier(t)
		credentialCreator := credential.NewMockCredentialTokenCreator(t)

		signIn := service.SignInService(
			logger,
			validationSuccessCase,
			execManager,
			userNotFoundCase,
			passwordVerifier,
			credentialCreator,
		)

		execManager.EXPECT().Executor().Return(nil)

		res, err := signIn(ctx, &dto.SignInRequest{
			Username: faker.Username(),
			Password: faker.Password(),
		})

		assert.ErrorIs(t, err, usecase.ErrAuthentication)
		assert.Nil(t, res)
	})

	t.Run("password not match", func(t *testing.T) {
		ctx := context.Background()

		execManager := data.NewMockExecutorManager[any](t)
		passwordVerifier := credential.NewMockPasswordVerifier(t)
		credentialCreator := credential.NewMockCredentialTokenCreator(t)

		user := &entity.User{
			ID:       uuid.New(),
			Username: faker.Username(),
			Password: faker.Password(),
		}

		signIn := service.SignInService(
			logger,
			validationSuccessCase,
			execManager,
			userExistCase(user),
			passwordVerifier,
			credentialCreator,
		)

		execManager.EXPECT().Executor().Return(nil)
		passwordVerifier.EXPECT().
			Verify(user.Password, user.Password).
			Return(credential.ErrPasswordMismatch)

		res, err := signIn(ctx, &dto.SignInRequest{
			Username: user.Username,
			Password: user.Password,
		})

		assert.ErrorIs(t, err, usecase.ErrAuthentication)
		assert.Nil(t, res)
	})

}
