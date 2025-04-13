package middleware

import (
	"strings"

	"github.com/fkrhykal/quickbid-account/internal/credential"
	"github.com/gofiber/fiber/v2"
)

func BearerMiddleware(credentialRetriever credential.UserCredentialRetriever) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorization, ok := c.GetReqHeaders()["Authorization"]
		if !ok {
			return c.Next()
		}
		if len(authorization) == 0 {
			return c.Next()
		}
		bearer := authorization[0]
		if strings.HasPrefix(bearer, "Bearer ") {
			token := strings.TrimPrefix(bearer, "Bearer ")
			userCredential, err := credentialRetriever.RetrieveUserCredential(c.UserContext(), token)
			if err != nil {
				return err
			}
			c.Locals("credential", userCredential)
		}
		return c.Next()
	}
}
