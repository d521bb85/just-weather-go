package web

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(router fiber.Router) {
	user := router.Group("/user")
	user.Post("/create", createUser...)
}
