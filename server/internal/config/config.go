package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	APIBaseURL  string
	CORSOrigins []string
	LogLevel    string
	Environment string
	LeagueID    int
	Year        int
	ESPNS2      string
	SWID        string
}

func Load() *Config {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development"
	}

	// Try loading .env files from multiple possible locations
	envFiles := []string{
		".env." + env,
		".env",
		"../.env." + env,
		"../.env",
		"../../.env." + env,
		"../../.env",
	}

	envLoaded := false
	for _, envFile := range envFiles {
		if err := godotenv.Load(envFile); err == nil {
			envLoaded = true
			break
		}
	}

	if !envLoaded {
		log.Printf("Warning: Could not load .env file from any location")
	}

	config := &Config{
		Port:        getEnv("PORT", "8080"),
		APIBaseURL:  getEnv("API_BASE_URL", getDefaultAPIURL(env)),
		LogLevel:    getEnv("LOG_LEVEL", getDefaultLogLevel(env)),
		Environment: env,
		LeagueID:    parseInt(getEnv("ESPN_LEAGUE_ID", "0")),
		Year:        parseInt(getEnv("ESPN_SEASON_YEAR", "2026")),
		ESPNS2:      getEnv("ESPN_S2_COOKIE", ""),
		SWID:        getEnv("ESPN_SWID_COOKIE", ""),
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

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Failed to parse integer: %s", s)
	}
	return i
}
