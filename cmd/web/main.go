package main

import (
	"just-weather-go/internal/config"
	"just-weather-go/internal/httpserver"
)

func main() {
	config.MustInit()
	httpserver.MustInit()
}
