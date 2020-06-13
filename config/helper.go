package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func DotEnvVariable(key string) string {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
