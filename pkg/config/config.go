package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	URL       string
	JWTSecret string
}

func Load() Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := os.Getenv("URL")
	port := os.Getenv("PORT")
	jwt := os.Getenv("JWT_SECRET")

	config := Config{
		Port:      port,
		URL:       url,
		JWTSecret: jwt,
	}

	return config
}
