package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func mongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGOURI")
}
func ApiKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("API_KEY")
}
func JwtKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("JWT_KEY")
}
func LocalHost() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("LOCAL_HOST")
}
