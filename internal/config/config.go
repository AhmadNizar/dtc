package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Username string
	Password string
	Database string
	Host     string
	Port     string
	SSLMode  string
	AppPort  string
}

func LoadConfig() *Config {

	err := godotenv.Load()

	if err != nil {
		log.Println("Warning: No .env file found, skipping...", err.Error())
	}

	return &Config{
		Username: getEnv("DATABASE_USERNAME", "default_username"),
		Password: getEnv("DATABASE_PASSWORD", "default_password"),
		Database: getEnv("DATABASE_NAME", "default_database"),
		Host:     getEnv("DATABASE_HOST", "localhost"),
		Port:     getEnv("DATABASE_PORT", "5432"),
		SSLMode:  getEnv("DATABASE_SSLMODE", "disable"),
		AppPort:  getEnv("APP_PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); len(value) > 0 {
		return value
	}

	return fallback
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		c.Username, c.Password, c.Database, c.Host, c.Port, c.SSLMode)
}
