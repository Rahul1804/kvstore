package config

import (
	"os"
)

// Config holds the application configuration.
type Config struct {
	Port string
}

// LoadConfig loads the configuration from environment variables.
func LoadConfig() *Config {
	return &Config{
		Port: getEnv("PORT", "8080"),
	}
}

// getEnv gets an environment variable or returns a default value.
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
