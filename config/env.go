package config

import (
	"log"
	"os"
	"strconv"

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

func getString(key, defaultString string) string {
	val := os.Getenv(key)
	if val == "" {
		val = defaultString
	}
	return val
}

func getInt(key string, defaultString int) int {
	valStr := os.Getenv(key)

	val, err := strconv.Atoi(valStr)
	if err != nil {
		return defaultString
	}

	return val
}

func getBool(key string, defaultString bool) bool {
	valStr := os.Getenv(key)

	val, err := strconv.ParseBool(valStr)
	if err != nil {
		return defaultString
	}

	return val
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		url: getString("DATABASE_URL", ""),
	}

}
