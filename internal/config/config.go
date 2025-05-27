package config

import (
	"log"
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

func MustInit() *Config {
	once.Do(func() {
		if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
			log.Fatalf("failed to read .env file: %v", err)
		}

		config = &Config{}
		if err := env.Parse(config); err != nil {
			log.Fatalf("failed to parse environment variables: %v", err)
		}
	})
	return config
}

func MustGet() *Config {
	if config == nil {
		log.Fatal("config is not initialized, call Init() first")
	}
	return config
}
