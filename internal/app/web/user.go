package web

import (
	"just-weather-go/internal/app/mutations"
	"just-weather-go/internal/webutils"

	"github.com/gofiber/fiber/v2"
)

type CreateUserRequestDTO struct {
	Email string `json:"email" validate:"required,email"`
}

type CreateUserResponseDTO struct{}

var createUser = []fiber.Handler{
	webutils.ValidateBody[CreateUserRequestDTO](),
	func(c *fiber.Ctx) error {
		body := c.Locals("body").(CreateUserRequestDTO)

		mutations.CreateUser(&mutations.CreateUserParameters{
			Email: body.Email,
		})

		return c.JSON(fiber.Map{
			"email": body.Email,
		})
	},
}
