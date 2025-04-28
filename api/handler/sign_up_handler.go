package handler

import (
	"log/slog"

	"github.com/fkrhykal/quickbid-account/api/response"
	"github.com/fkrhykal/quickbid-account/internal/dto"
	"github.com/fkrhykal/quickbid-account/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

func SetupSignUp(log *slog.Logger, router fiber.Router, signUp usecase.SignUp) {
	router.Post("/sign-up", func(ctx *fiber.Ctx) error {
		userCtx := ctx.UserContext()
		log.DebugContext(userCtx, "Handling sign-up request")

		req := new(dto.SignUpRequest)
		if err := ctx.BodyParser(&req); err != nil {
			log.DebugContext(userCtx, "Failed to parse request body", slog.Any("error", err))
			return err
		}

		log.DebugContext(userCtx, "Request body parsed successfully")

		res, err := signUp(userCtx, req)
		if err != nil {
			log.DebugContext(userCtx, "Sign-up process failed", slog.Any("error", err))
			return err
		}

		log.DebugContext(userCtx, "User signed up successfully")
		return response.SendData(ctx, fiber.StatusCreated, res)
	})
}
