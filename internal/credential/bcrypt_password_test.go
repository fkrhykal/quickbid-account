package credential_test

import (
	"log/slog"
	"os"
	"testing"

	"github.com/fkrhykal/quickbid-account/internal/credential"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestBcryptPassword(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))

	passwordManager := credential.NewBcryptPasswordManager(logger)

	password := faker.Password()

	hashedPassword, err := passwordManager.Hash(password)
	assert.NoError(t, err)

	err = passwordManager.Verify(hashedPassword, password)
	assert.NoError(t, err)

	err = passwordManager.Verify(hashedPassword, faker.Password())
	assert.ErrorIs(t, err, credential.ErrPasswordMismatch)
}
