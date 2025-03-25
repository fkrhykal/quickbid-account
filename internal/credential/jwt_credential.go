package credential

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtCredentialManager struct {
	log *slog.Logger
	key []byte
}

func (j *JwtCredentialManager) CreateCredentialToken(ctx context.Context, userCredential *UserCredential) (string, error) {
	j.log.DebugContext(ctx, "Creating credential token")

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, CredentialClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
		UserCredential: userCredential,
	})

	tokenString, err := jwtToken.SignedString(j.key)
	if err != nil {
		j.log.DebugContext(ctx, "Failed to sign JWT token", slog.Any("error", err))
		return "", err
	}

	j.log.DebugContext(ctx, "Credential token created successfully")
	return tokenString, nil
}

func (j *JwtCredentialManager) RetrieveUserCredential(ctx context.Context, token string) (*UserCredential, error) {
	j.log.DebugContext(ctx, "Retrieving user credential from token")

	jwtToken, err := jwt.ParseWithClaims(token, &CredentialClaims{}, func(t *jwt.Token) (any, error) {
		return j.key, nil
	})
	if err == nil {
		credentialClaims, ok := jwtToken.Claims.(*CredentialClaims)
		if !ok {
			j.log.DebugContext(ctx, "Invalid credential claims")
			return nil, ErrCredentialInvalid
		}
		j.log.DebugContext(ctx, "User credential retrieved successfully")
		return credentialClaims.UserCredential, nil
	}

	if errors.Is(err, jwt.ErrTokenExpired) {
		j.log.DebugContext(ctx, "Token has expired")
		return nil, ErrCredentialExpired
	}

	j.log.DebugContext(ctx, "Invalid credential token", slog.Any("error", err))
	return nil, ErrCredentialInvalid
}

func NewJwtCredentialManager(config *JwtCredentialManagerConfig) CredentialManager {
	config.Logger.Debug("Initializing JWT Credential Manager")
	return &JwtCredentialManager{
		log: config.Logger,
		key: config.SecretKey,
	}
}

type JwtCredentialManagerConfig struct {
	Logger    *slog.Logger
	SecretKey []byte
}

type CredentialClaims struct {
	jwt.RegisteredClaims
	UserCredential *UserCredential `json:"userCredential"`
}
