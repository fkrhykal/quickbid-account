package validation_test

import (
	"context"
	"testing"

	"github.com/fkrhykal/quickbid-account/internal/dto"
	"github.com/fkrhykal/quickbid-account/internal/validation"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestSignUpRequestValidator(t *testing.T) {
	ctx := context.Background()
	t.Run("sign up request validation success", func(t *testing.T) {
		err := validation.ValidateSignUpRequest(ctx, &dto.SignUpRequest{
			Username: faker.Username(),
			Password: "dfnerr7343U9340**",
		})
		assert.NoError(t, err)
	})

	t.Run("sign up request validation fail", func(t *testing.T) {
		err := validation.ValidateSignUpRequest(ctx, &dto.SignUpRequest{
			Username: "fe3*3)_",
			Password: "jdfeif fnrfr",
		})
		validationError, ok := err.(*validation.ValidationError)
		assert.True(t, ok)
		assert.Len(t, validationError.Detail, 2)
		assert.Equal(t, validationError.Detail["username"], "username must ended with letter or number")
		assert.Equal(t, validationError.Detail["password"], "password shouldn't contain white space")
	})
}
