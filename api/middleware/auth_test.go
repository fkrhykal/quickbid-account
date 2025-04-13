package middleware_test

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"testing"

	"github.com/fkrhykal/quickbid-account/api/middleware"
	"github.com/fkrhykal/quickbid-account/internal/credential"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBearerMiddleware(t *testing.T) {
	app := fiber.New()

	jwtManager := credential.NewJwtCredentialManager(&credential.JwtCredentialManagerConfig{
		Logger:    slog.Default(),
		SecretKey: []byte("secret"),
	})

	userCredential := &credential.UserCredential{
		ID: uuid.New(),
	}
	bearer, err := jwtManager.CreateCredentialToken(context.Background(), userCredential)
	assert.NoError(t, err)

	app.Get("/test", middleware.BearerMiddleware(jwtManager), func(c *fiber.Ctx) error {
		userCredential, ok := c.Locals("credential").(*credential.UserCredential)
		if !ok {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		return c.JSON(userCredential)
	})

	req, err := http.NewRequest(http.MethodGet, "/test", nil)
	assert.NoError(t, err)

	req.Header.Add("Authorization", "Bearer "+bearer)

	res, err := app.Test(req, 5_000)
	assert.NoError(t, err)

	var body *credential.UserCredential
	err = json.NewDecoder(res.Body).Decode(&body)
	if assert.NoError(t, err) {
		assert.EqualValues(t, userCredential, body)
	}
}
