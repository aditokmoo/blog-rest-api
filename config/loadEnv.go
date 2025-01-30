package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	requiredEnvVars := []string{"PORT", "HOST", "DB_URL", "JWT_SECRET", "SMTP_USER", "SMTP_PASS"}
	for _, envKey := range requiredEnvVars {
		if os.Getenv(envKey) == "" {
			log.Fatalf("Required environment variable %s is not set", envKey)
		}
	}
}