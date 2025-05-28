package webserver

import (
	"fmt"
	"just-weather-go/internal/app/web"
	"just-weather-go/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func MustStart(cfg *config.Config) {
	server := fiber.New()
	server.Use(helmet.New())

	v1 := server.Group("/api/v1")
	web.RegisterRoutes(v1)

	addr := fmt.Sprintf(":%s", cfg.HTTPPort)
	if err := server.Listen(addr); err != nil {
		panic(fmt.Sprintf("failed to start HTTP server: %v", err))
	}
}
