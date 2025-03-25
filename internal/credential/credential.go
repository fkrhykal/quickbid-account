package credential

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var ErrCredentialInvalid = errors.New("credential token invalid")
var ErrCredentialExpired = errors.New("credential token expired")

type UserCredential struct {
	ID uuid.UUID `json:"id"`
}

type CreateCredentialToken func(ctx context.Context, userCredential *UserCredential) (string, error)
type RetrieveUserCredential func(ctx context.Context, token string) (*UserCredential, error)

type CredentialTokenCreator interface {
	CreateCredentialToken(ctx context.Context, userCredential *UserCredential) (string, error)
}

type UserCredentialRetriever interface {
	RetrieveUserCredential(ctx context.Context, token string) (*UserCredential, error)
}

type CredentialManager interface {
	CredentialTokenCreator
	UserCredentialRetriever
}
