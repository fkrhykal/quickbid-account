package handler

import (
	"log/slog"

	"github.com/fkrhykal/quickbid-account/api/response"
	"github.com/fkrhykal/quickbid-account/internal/dto"
	"github.com/fkrhykal/quickbid-account/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

func SignInHandler(log *slog.Logger, signIn usecase.SignIn) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userCtx := ctx.UserContext()
		log.DebugContext(userCtx, "Handling user sign in")

		req := new(dto.SignInRequest)
		if err := ctx.BodyParser(req); err != nil {
			log.DebugContext(userCtx, "Failed to parse request body", slog.Any("error", err))
			return err
		}

		log.DebugContext(userCtx, "Successfully parse request body")

		res, err := signIn(userCtx, req)
		if err != nil {
			log.DebugContext(userCtx, "Sign in process failed", slog.Any("error", err))
			return err
		}

		log.DebugContext(userCtx, "User signed in successfully")
		return response.SendData(ctx, fiber.StatusOK, res)
	}
}
