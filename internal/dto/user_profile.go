package dto

import "github.com/google/uuid"

type UserProfile struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Avatar   string    `json:"avatar"`
}
