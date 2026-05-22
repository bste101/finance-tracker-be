package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort       string
	DatabaseURL   string
	JWTSecret     string
	JWTExpiration string
}

func Load() *Config {
	_ = godotenv.Load()

	appPort := envOrDefault("APP_PORT", "8080")
	jwtSecret := envOrDefault("JWT_SECRET", "your_jwt_secret")
	jwtExpiration := envOrDefault("JWT_EXPIRATION", "24h")

	return &Config{
		AppPort:       appPort,
		DatabaseURL:   os.Getenv("DATABASE_URL"),
		JWTSecret:     jwtSecret,
		JWTExpiration: jwtExpiration,
	}
}

func envOrDefault(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func durationOrDefault(key string, fallback time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	parsed, err := time.ParseDuration(value)
	if err != nil {
		return fallback
	}
	return parsed
}
