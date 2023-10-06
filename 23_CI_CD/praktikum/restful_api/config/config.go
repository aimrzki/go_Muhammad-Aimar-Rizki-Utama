package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func LoadConfig() Config {
	// Load variabel lingkungan dari file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return Config{
		DB_Username: os.Getenv("DB_Username"),
		DB_Password: os.Getenv("DB_Password"),
		DB_Port:     os.Getenv("DB_Port"),
		DB_Host:     os.Getenv("DB_Host"),
		DB_Name:     os.Getenv("DB_Name"),
	}
}
