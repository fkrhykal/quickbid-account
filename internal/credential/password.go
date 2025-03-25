package credential

import "errors"

var ErrPasswordMismatch = errors.New("password mismatch")

type PasswordHasher interface {
	Hash(password string) (string, error)
}
type PasswordVerifier interface {
	Verify(hashedPassword, password string) error
}

type PasswordManager interface {
	PasswordHasher
	PasswordVerifier
}
