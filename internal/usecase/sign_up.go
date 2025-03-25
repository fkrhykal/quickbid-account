package usecase

import (
	"context"
	"errors"

	"github.com/fkrhykal/quickbid-account/internal/dto"
)

var ErrUsernameAlreadyUsed = errors.New("username already used")

type SignUp func(ctx context.Context, req *dto.SignUpRequest) (*dto.SignUpResponse, error)
