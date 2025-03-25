package dto

import "github.com/google/uuid"

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInResponse struct {
	ID          uuid.UUID `json:"id"`
	BearerToken string    `json:"bearerToken"`
}
