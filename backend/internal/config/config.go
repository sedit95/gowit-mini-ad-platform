package config

import (
	"errors"
	"os"
)

// Config holds the configuration for the application.
type Config struct {
	DatabaseURL string
	Port        string
}

// Load reads configuration from the environment variables.
// It returns an error if the required DATABASE_URL is missing.
func Load() (Config, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return Config{}, errors.New("DATABASE_URL environment variable is required")
	}

	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "8080"
	}

	return Config{
		DatabaseURL: databaseURL,
		Port:        port,
	}, nil
}
