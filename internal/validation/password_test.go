package validation_test

import (
	"testing"

	"github.com/fkrhykal/quickbid-account/internal/validation"
	"github.com/stretchr/testify/assert"
)

func TestValidatePassword(t *testing.T) {
	t.Run("password valid", func(t *testing.T) {
		err := validation.ValidatePassword("Aa1!Bb2@Cc3#")
		assert.NoError(t, err)
	})

	t.Run("password too short", func(t *testing.T) {
		err := validation.ValidatePassword("Aa1!Bb2@")
		assert.Equal(t, validation.ErrPasswordLength, err)
	})

	t.Run("password too long", func(t *testing.T) {
		err := validation.ValidatePassword("Aa1!Bb2@Cc3#Dd4$Ee5%Ff6&Gg7*sdhewuhfewihfioenfiefeifirneyityteifmhertciurrn7534957348534n5hffruinfer)))44n")
		assert.Equal(t, validation.ErrPasswordLength, err)
	})

	t.Run("password contain whitespace", func(t *testing.T) {
		err := validation.ValidatePassword("Aa1! Bb2@Cc3#")
		assert.Equal(t, validation.ErrPasswordContainWhiteSpace, err)
	})

	t.Run("password contain non-ASCII", func(t *testing.T) {
		err := validation.ValidatePassword("Aa1!世界efwefw")
		assert.Equal(t, validation.ErrPasswordNonASCII, err)
	})

	t.Run("password has no uppercase", func(t *testing.T) {
		err := validation.ValidatePassword("aa1!bb2@cc3#")
		assert.Equal(t, validation.ErrPasswordInvalid, err)
	})

	t.Run("password has no lowercase", func(t *testing.T) {
		err := validation.ValidatePassword("AA1!BB2@CC3#")
		assert.Equal(t, validation.ErrPasswordInvalid, err)
	})

	t.Run("password has no digit", func(t *testing.T) {
		err := validation.ValidatePassword("Aa!Bb@Cc#edm")
		assert.Equal(t, validation.ErrPasswordInvalid, err)
	})

	t.Run("password has no special character", func(t *testing.T) {
		err := validation.ValidatePassword("Aa1Bb2Cc3Dd4")
		assert.Equal(t, validation.ErrPasswordInvalid, err)
	})
}
