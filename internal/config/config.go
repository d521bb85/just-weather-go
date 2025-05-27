package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	HTTPPort string `env:"HTTP_PORT"`
}

var (
	once   sync.Once
	config *Config
)

func MustInit() {
	once.Do(func() {
		if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
			panic(fmt.Sprintf("failed to read .env file: %v", err))
		}

		config = &Config{}
		if err := env.Parse(config); err != nil {
			panic(fmt.Sprintf("failed to parse environment variables: %v", err))
		}
	})
}

func MustGet() *Config {
	if config == nil {
		panic("config is not initialized, call Init() first")
	}
	return config
}
