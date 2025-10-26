package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	APIBaseURL  string
	CORSOrigins []string
	LogLevel    string
	Environment string
}

func Load() *Config {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development"
	}

	envFile := ".env." + env
	if err := godotenv.Load(envFile); err != nil {
		if err := godotenv.Load(); err != nil {
			log.Printf("Warning: Could not load .env file: %v", err)
		}
	}

	config := &Config{
		Port:        getEnv("PORT", "8080"),
		APIBaseURL:  getEnv("API_BASE_URL", getDefaultAPIURL(env)),
		LogLevel:    getEnv("LOG_LEVEL", getDefaultLogLevel(env)),
		Environment: env,
	}

	corsOrigins := getEnv("CORS_ORIGINS", getDefaultCORSOrigins(env))
	config.CORSOrigins = strings.Split(corsOrigins, ",")
	for i, origin := range config.CORSOrigins {
		config.CORSOrigins[i] = strings.TrimSpace(origin)
	}

	if env == "prod" {
		if config.APIBaseURL == "" {
			log.Fatal("API_BASE_URL is required in production")
		}
		if len(config.CORSOrigins) == 0 || config.CORSOrigins[0] == "" {
			log.Fatal("CORS_ORIGINS is required in production")
		}
	}

	return config
}

func getDefaultAPIURL(env string) string {
	switch env {
	case "prod":
		return ""
	default:
		return "http://localhost:8080"
	}
}

func getDefaultCORSOrigins(env string) string {
	switch env {
	case "prod":
		return ""
	default:
		return "http://localhost:5173"
	}
}

func getDefaultLogLevel(env string) string {
	switch env {
	case "prod":
		return "info"
	default:
		return "debug"
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
