package handler

import (
	"github.com/gofiber/fiber/v2"
)

func CurrentUserProfile(router fiber.Router) {
	router.Get("/users/_current", func(ctx *fiber.Ctx) error {
		return nil
	})
}
