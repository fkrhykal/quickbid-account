package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"

	"github.com/fkrhykal/quickbid-account/api/response"
	"github.com/fkrhykal/quickbid-account/internal/usecase"
	"github.com/fkrhykal/quickbid-account/internal/validation"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(log *slog.Logger) fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		if typeError, ok := err.(*json.UnmarshalTypeError); ok {
			log.WarnContext(ctx.UserContext(), "received request body with mismatch data type", slog.Any("error", typeError))
			return response.SendError(ctx, fiber.StatusUnprocessableEntity, fiber.Map{
				typeError.Field: fmt.Sprintf("%s has wrong type", typeError.Field),
			})
		}

		if valError, ok := err.(*validation.ValidationError); ok {
			log.WarnContext(ctx.UserContext(), "received invalid request body", slog.Any("error", valError))
			return response.SendError(ctx, fiber.StatusBadRequest, valError.Detail)
		}

		if errors.Is(err, usecase.ErrUsernameAlreadyUsed) {
			log.WarnContext(ctx.UserContext(), "application err", slog.Any("error", err))
			return response.SendError(ctx, fiber.StatusConflict, err)
		}

		if errors.Is(err, usecase.ErrAuthentication) {
			log.WarnContext(ctx.UserContext(), "authentication error")
			return response.SendFiberError(ctx, fiber.ErrUnauthorized)
		}

		if fiberError, ok := err.(*fiber.Error); ok {
			log.WarnContext(ctx.UserContext(), "application error", slog.Any("error", fiberError))
			return response.SendFiberError(ctx, fiberError)
		}

		log.ErrorContext(ctx.UserContext(), "unexpected error", slog.Any("error", err))
		return response.SendFiberError(ctx, fiber.ErrInternalServerError)
	}
}
