package credential

import (
	"log/slog"

	"golang.org/x/crypto/bcrypt"
)

type BcryptPasswordManager struct {
	log *slog.Logger
}

func (b *BcryptPasswordManager) Verify(hashedPassword, password string) error {
	b.log.Debug("Verifying password")
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		b.log.Debug("Password verification failed", slog.Any("error", err))
		return ErrPasswordMismatch
	}
	b.log.Debug("Password verification succeeded")
	return nil
}

func (b *BcryptPasswordManager) Hash(password string) (string, error) {
	b.log.Debug("Hashing password")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		b.log.Debug("Password hashing failed", slog.Any("error", err))
		return "", err
	}
	b.log.Debug("Password hashing succeeded")
	return string(hashedPassword), nil
}

func NewBcryptPasswordManager(log *slog.Logger) PasswordManager {
	log.Debug("Initializing BcryptPasswordManager")
	return &BcryptPasswordManager{
		log: log,
	}
}
