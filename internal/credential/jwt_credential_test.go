package credential_test

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/fkrhykal/quickbid-account/internal/credential"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var jwtTestConfig = &credential.JwtCredentialManagerConfig{
	Logger: slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})),
	SecretKey: []byte("secret"),
}

func TestJwtCredentialManager(t *testing.T) {
	credentialManager := credential.NewJwtCredentialManager(jwtTestConfig)

	ctx := context.Background()
	id := uuid.New()

	t.Run("token valid", func(t *testing.T) {
		cred := &credential.UserCredential{
			ID: id,
		}

		token, err := credentialManager.CreateCredentialToken(ctx, cred)
		assert.NoError(t, err)

		userCredential, err := credentialManager.RetrieveUserCredential(ctx, token)
		assert.NoError(t, err)

		assert.EqualValues(t, cred, userCredential)
	})

	t.Run("token invalid", func(t *testing.T) {
		userCredential, err := credentialManager.RetrieveUserCredential(ctx, "nefifrifrifrfrifr")
		assert.ErrorIs(t, err, credential.ErrCredentialInvalid)

		assert.Nil(t, userCredential)
	})
}
