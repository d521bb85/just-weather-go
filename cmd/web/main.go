package main

import (
	"just-weather-go/internal/config"
	"just-weather-go/internal/httpserver"
)

func main() {
	cfg := config.MustInit()
	httpserver.MustStart(cfg)
}
