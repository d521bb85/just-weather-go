package routes

import "github.com/gofiber/fiber/v2"

func Register(router fiber.Router) {
	router.Post("/", create)
}

func create(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{})
}
