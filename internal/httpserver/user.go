package httpserver

import "github.com/gofiber/fiber/v2"

func create(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{})
}
