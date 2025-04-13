package usecase

import (
	"context"

	"github.com/fkrhykal/quickbid-account/internal/dto"
	"github.com/google/uuid"
)

type GetUserProfile func(ctx context.Context, userID uuid.UUID) (*dto.UserProfile, error)
