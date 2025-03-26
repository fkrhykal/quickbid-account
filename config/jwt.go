package config

import (
	"log/slog"
	"os"

	"github.com/fkrhykal/quickbid-account/internal/credential"
)

var JwtTestConfig = &credential.JwtCredentialManagerConfig{
	Logger: slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})),
	SecretKey: []byte("test"),
}

var JwtDevConfig = &credential.JwtCredentialManagerConfig{
	Logger: slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})),
	SecretKey: []byte("dev"),
}

func JwtProdConfig(logger *slog.Logger) *credential.JwtCredentialManagerConfig {
	return &credential.JwtCredentialManagerConfig{
		Logger:    logger,
		SecretKey: []byte(MustEnvString("APP_SECRET_KEY")),
	}
}
