package service

import (
	"context"

	"github.com/fkrhykal/quickbid-account/internal/data"
	"github.com/fkrhykal/quickbid-account/internal/dto"
	"github.com/fkrhykal/quickbid-account/internal/repository"
	"github.com/fkrhykal/quickbid-account/internal/usecase"
	"github.com/google/uuid"
)

func UserProfileService[T any](executorManager data.ExecutorManager[T], findUserByID repository.FindUserByID[T]) usecase.GetUserProfile {
	return func(ctx context.Context, userID uuid.UUID) (*dto.UserProfile, error) {
		user, err := findUserByID(ctx, executorManager.Executor(), userID)
		if err != nil {
			return nil, err
		}
		return &dto.UserProfile{
			ID:       user.ID,
			Username: user.Username,
			Avatar:   user.Avatar,
		}, nil
	}
}
