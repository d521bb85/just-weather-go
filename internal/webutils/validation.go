package webutils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ValidateBody[D any]() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body D

		if err := c.BodyParser(&body); err != nil {
			return c.SendStatus(fiber.ErrBadRequest.Code)
		}

		if err := validate.Struct(&body); err != nil {
			return c.SendStatus(fiber.ErrBadRequest.Code)
		}

		c.Locals("body", body)
		return c.Next()
	}
}
