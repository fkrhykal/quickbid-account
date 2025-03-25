package validation

import (
	"context"

	"github.com/fkrhykal/quickbid-account/internal/dto"
)

func ValidateSignUpRequest(ctx context.Context, req *dto.SignUpRequest) error {
	detail := make(Detail)

	if err := ValidateUsername(req.Username); err != nil {
		detail.Add("username", err)
	}

	if err := ValidatePassword(req.Password); err != nil {
		detail.Add("password", err)
	}

	if len(detail) > 0 {
		return &ValidationError{
			Detail: detail,
		}
	}
	return nil
}
