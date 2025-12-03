package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
	DBMode string

	ServerPort string
	JWTSecret  string

	SMTPEmail string
	SMTPPass  string
	SMTPHost  string
	SMTPPort  string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_NAME"),
		DBMode: os.Getenv("DB_SSLMODE"),

		ServerPort: os.Getenv("PORT"),
		JWTSecret:  os.Getenv("SECRET_KEY"),

		SMTPEmail: os.Getenv("SMTP_EMAIL"),
		SMTPPass:  os.Getenv("SMTP_PASSWORD"),
		SMTPHost:  os.Getenv("SMTP_HOST"),
		SMTPPort:  os.Getenv("SMTP_PORT"),
	}
}
