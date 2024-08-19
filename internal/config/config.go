package config

import (
	"os"
	"strconv"
)

type Config struct {
	DatabasePath string
	ServerPort   int
	LogLevel     string
}

func LoadConfig() *Config {
	return &Config{
		DatabasePath: getEnv("DATABASE_PATH", "portfolio.db"),
		ServerPort:   getEnvAsInt("SERVER_PORT", 8080),
		LogLevel:     getEnv("LOG_LEVEL", "info"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
