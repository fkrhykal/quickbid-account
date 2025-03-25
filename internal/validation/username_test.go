package validation_test

import (
	"testing"

	"github.com/fkrhykal/quickbid-account/internal/validation"
	"github.com/stretchr/testify/assert"
)

type usernameTC struct {
	Name     string
	Error    error
	Username string
}

func TestUsername(t *testing.T) {
	t.Run("username valid", func(t *testing.T) {
		err := validation.ValidateUsername("f_uee3h")
		assert.NoError(t, err)
	})
	t.Run("username too short", func(t *testing.T) {
		err := validation.ValidateUsername("sdsd")
		assert.ErrorIs(t, err, validation.ErrUsernameLength)
	})
	t.Run("username too long", func(t *testing.T) {
		err := validation.ValidateUsername("qwertyuiopasdfghjklzxcvbbn")
		assert.ErrorIs(t, err, validation.ErrUsernameLength)
	})
	t.Run("username started with .", func(t *testing.T) {
		err := validation.ValidateUsername(".dfsdfd")
		assert.ErrorIs(t, err, validation.ErrUsernamePrefix)
	})
	t.Run("username ended with .", func(t *testing.T) {
		err := validation.ValidateUsername("ssfsdssd.")
		assert.ErrorIs(t, err, validation.ErrUsernameSuffix)
	})
	t.Run("username contain invalid chararter", func(t *testing.T) {
		err := validation.ValidateUsername("sdsdsknde&nf")
		assert.ErrorIs(t, err, validation.ErrUsernameInvalid)
	})
	t.Run("username has consecutive underscore", func(t *testing.T) {
		err := validation.ValidateUsername("dsdijeie__423")
		assert.ErrorIs(t, err, validation.ErrUsernameUnderscore)
	})
}
