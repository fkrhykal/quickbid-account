package service

import (
	"context"

	"github.com/fkrhykal/quickbid-account/internal/dto"
	"github.com/fkrhykal/quickbid-account/internal/usecase"
	"github.com/google/uuid"
)

func UserProfileService() usecase.GetUserProfile {
	return func(ctx context.Context, userID uuid.UUID) (*dto.UserProfile, error) {
		return nil, nil
	}
}
