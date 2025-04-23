package handler

import (
	"github.com/fkrhykal/quickbid-account/api/middleware"
	"github.com/fkrhykal/quickbid-account/api/response"
	"github.com/fkrhykal/quickbid-account/internal/credential"
	"github.com/fkrhykal/quickbid-account/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

func CurrentUserProfile(router fiber.Router, credentialRetriever credential.UserCredentialRetriever, getUserProfile usecase.GetUserProfile) {
	router.Get("/users/_current",
		middleware.BearerMiddleware(credentialRetriever),
		middleware.AuthenticationMiddleware(),
		func(c *fiber.Ctx) error {
			userCredential := c.Locals("credential").(*credential.UserCredential)
			res, err := getUserProfile(c.UserContext(), userCredential.ID)
			if err != nil {
				return err
			}
			return response.SendData(c, fiber.StatusOK, res)
		})
}
