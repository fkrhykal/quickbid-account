package route

import "github.com/gofiber/fiber/v2"

func SignInRoute(router fiber.Router, handler fiber.Handler) {
	router.Post("/sign-in", handler)
}
