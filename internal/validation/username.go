package validation

import (
	"errors"
)

var ErrUsernameLength = errors.New("username must be between 6 and 18 characters")
var ErrUsernamePrefix = errors.New("username must started with letter")
var ErrUsernameSuffix = errors.New("username must ended with letter or number")
var ErrUsernameInvalid = errors.New("username should only contain alphanumeric and underscore")
var ErrUsernameUnderscore = errors.New("username shouldn't have consecutive underscore")

func ValidateUsername(username string) error {
	if !HasLength(username, 6, 18) {
		return ErrUsernameLength
	}
	if !HasPrefixChar(username, IsLowercase, IsUppercase) {
		return ErrUsernamePrefix
	}
	if !HasSuffixChar(username, IsLowercase, IsUppercase, IsDigit) {
		return ErrUsernameSuffix
	}
	var tracker rune
	for _, char := range username {
		if tracker == '_' && char == '_' {
			return ErrUsernameUnderscore
		}
		tracker = char
		if IsAlphanumeric(char) {
			continue
		}
		if char == '_' {
			continue
		}
		return ErrUsernameInvalid
	}
	return nil
}
