package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	url string
}

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
		return
	}
	log.Println("Successfully loaded .env file")
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		url: os.Getenv("DATABASE_URL"),
	}

}
