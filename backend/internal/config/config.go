package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Discord struct {
		ClientID     string
		ClientSecret string
	}
	Server struct {
		Port int
	}
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: .env file not found")
	}

	cfg := &Config{}

	cfg.Discord.ClientID = getEnvOrDefault("CLIENTID", "")
	cfg.Discord.ClientSecret = getEnvOrDefault("CLIENTSECRET", "")

	port := getEnvOrDefault("PORT", "8085")

	v, err := strconv.Atoi(port)

	if err != nil {
		return nil, err
	}

	cfg.Server.Port = v

	if cfg.Discord.ClientID == "" {
		return nil, fmt.Errorf("CLIENTID environment variable is required")
	}

	if cfg.Discord.ClientSecret == "" {
		return nil, fmt.Errorf("SECRETID environment variable is required")
	}

	return cfg, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
