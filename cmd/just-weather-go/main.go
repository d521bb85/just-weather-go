package main

import (
	"just-weather-go/internal/config"
	"just-weather-go/internal/webserver"
)

func main() {
	cfg := config.MustInit()

	webserver.MustStart(cfg)
}
