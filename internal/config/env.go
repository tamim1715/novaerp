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

	GinMode string
}

var AppConfig Config

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Println(".env file not found, using system environment variables")
	}

	AppConfig = Config{
		AppName:    os.Getenv("APP_NAME"),
		AppEnv:     os.Getenv("APP_ENV"),
		Port:       os.Getenv("PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		GinMode:    os.Getenv("GIN_MODE"),
	}
}
