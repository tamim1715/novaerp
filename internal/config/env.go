// Package config manages application configuration loaded from environment variables.
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	AppEnv  string
	Port    string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	GinMode            string
	JWTSecret          string
	JWTPrivateKeyPEM   string
	JWTPublicKeyPEM    string
	CORSAllowedOrigins string
}

var AppConfig Config

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return fallback
}

// LoadEnv loads configuration from .env file or system environment variables with sensible defaults.
func LoadEnv() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using system environment variables")
	}

	AppConfig = Config{
		AppName:            getEnv("APP_NAME", "NovaERP"),
		AppEnv:             getEnv("APP_ENV", "development"),
		Port:               getEnv("PORT", "8080"),
		DBHost:             getEnv("DB_HOST", "localhost"),
		DBPort:             getEnv("DB_PORT", "5432"),
		DBUser:             getEnv("DB_USER", "postgres"),
		DBPassword:         getEnv("DB_PASSWORD", "postgres"),
		DBName:             getEnv("DB_NAME", "novaerp"),
		GinMode:            getEnv("GIN_MODE", "debug"),
		JWTSecret:          getEnv("JWT_SECRET", "novaerp-super-secret-key-change-in-production"),
		JWTPrivateKeyPEM:   getEnv("JWT_PRIVATE_KEY", ""),
		JWTPublicKeyPEM:    getEnv("JWT_PUBLIC_KEY", ""),
		CORSAllowedOrigins: getEnv("CORS_ALLOWED_ORIGINS", "*"),
	}

	return &AppConfig
}
