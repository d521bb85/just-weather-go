package httpserver

import (
	"fmt"
	"just-weather-go/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func MustInit() {
	cfg := config.MustGet()

	server := fiber.New()
	server.Use(helmet.New())

	api := server.Group("/api/v1")

	user := api.Group("/users")
	user.Post("/", create)

	server.Listen(fmt.Sprintf(":%s", cfg.HTTPPort))
}
