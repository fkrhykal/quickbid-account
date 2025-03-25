package route

import "github.com/gofiber/fiber/v2"

func SignUpRoute(router fiber.Router, handler fiber.Handler) {
	router.Post("/sign-up", handler)
}
