package usecase

import (
	"context"
	"errors"

	"github.com/fkrhykal/quickbid-account/internal/dto"
)

var ErrAuthentication = errors.New("authentication error")

type SignIn func(ctx context.Context, req *dto.SignInRequest) (*dto.SignInResponse, error)
