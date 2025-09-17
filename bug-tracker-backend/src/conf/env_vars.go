package conf

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVars() {
	err := godotenv.Load()
	if err != nil {
		// In production (like Vercel), .env file may not exist
		// Environment variables are provided by the platform
		log.Printf("Warning: .env file not found, using system environment variables: %v", err)
	}
}
