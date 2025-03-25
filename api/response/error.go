package response

import "github.com/gofiber/fiber/v2"

type Error[T any] struct {
	Code  int `json:"code"`
	Error T   `json:"error"`
}

func SendError[T any](ctx *fiber.Ctx, code int, err T) error {
	return ctx.Status(code).JSON(&Error[T]{Code: code, Error: err})
}

func SendFiberError(ctx *fiber.Ctx, err *fiber.Error) error {
	return ctx.Status(err.Code).JSON(&Error[string]{
		Code:  err.Code,
		Error: err.Message,
	})
}
