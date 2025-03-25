package response

import "github.com/gofiber/fiber/v2"

type Data[T any] struct {
	Code int `json:"code"`
	Data T   `json:"data"`
}

func SendData[T any](ctx *fiber.Ctx, code int, data T) error {
	return ctx.Status(code).JSON(&Data[T]{
		Code: code,
		Data: data,
	})
}
