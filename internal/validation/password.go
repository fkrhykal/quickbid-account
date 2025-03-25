package validation

import "errors"

var ErrPasswordLength = errors.New("password must be between 12 and 36 characters")
var ErrPasswordContainWhiteSpace = errors.New("password shouldn't contain white space")
var ErrPasswordNonASCII = errors.New("password should only contain ascii character")
var ErrPasswordInvalid = errors.New("password should contain uppercase, lowercase, number and special character")

func ValidatePassword(password string) error {
	if !HasLength(password, 12, 64) {
		return ErrPasswordLength
	}
	var hasLowercase bool
	var hasUppercase bool
	var hasDigit bool
	var hasSpecialChar bool
	for _, char := range password {
		if char == ' ' {
			return ErrPasswordContainWhiteSpace
		}
		if IsLowercase(char) {
			hasLowercase = true
			continue
		}
		if IsUppercase(char) {
			hasUppercase = true
			continue
		}
		if IsDigit(char) {
			hasDigit = true
			continue
		}
		if IsSpecialChar(char) {
			hasSpecialChar = true
			continue
		}
		return ErrPasswordNonASCII
	}
	if hasLowercase && hasUppercase && hasDigit && hasSpecialChar {
		return nil
	}
	return ErrPasswordInvalid
}
