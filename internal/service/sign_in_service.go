package service

import (
	"context"
	"log/slog"

	"github.com/fkrhykal/quickbid-account/internal/credential"
	"github.com/fkrhykal/quickbid-account/internal/data"
	"github.com/fkrhykal/quickbid-account/internal/dto"
	"github.com/fkrhykal/quickbid-account/internal/repository"
	"github.com/fkrhykal/quickbid-account/internal/usecase"
	"github.com/fkrhykal/quickbid-account/internal/validation"
)

func SignInService[T any](
	log *slog.Logger,
	validate validation.Validator[*dto.SignInRequest],
	execManager data.ExecutorManager[T],
	findUserByUsername repository.FindUserByUsername[T],
	passwordVerifier credential.PasswordVerifier,
	credentialCreator credential.CredentialTokenCreator,
) usecase.SignIn {
	return func(ctx context.Context, req *dto.SignInRequest) (*dto.SignInResponse, error) {
		log.DebugContext(ctx, "Starting Sign in process")

		if err := validate(ctx, req); err != nil {
			log.DebugContext(ctx, "Validation failed")
			return nil, usecase.ErrAuthentication
		}

		user, err := findUserByUsername(ctx, execManager.Executor(), req.Username)
		if err != nil {
			log.DebugContext(ctx, "Error finding user by username", slog.Any("error", err))
			return nil, err
		}
		if user == nil {
			log.DebugContext(ctx, "User not found")
			return nil, usecase.ErrAuthentication
		}

		if err := passwordVerifier.Verify(user.Password, req.Password); err != nil {
			log.DebugContext(ctx, "Password verification failed")
			return nil, usecase.ErrAuthentication
		}

		log.DebugContext(ctx, "Password verified successfully")

		bearerToken, err := credentialCreator.CreateCredentialToken(ctx, &credential.UserCredential{ID: user.ID})
		if err != nil {
			log.DebugContext(ctx, "Error creating credential token", slog.Any("error", err))
			return nil, err
		}

		log.DebugContext(ctx, "Sign in process success")
		return &dto.SignInResponse{ID: user.ID, BearerToken: bearerToken}, nil
	}
}
