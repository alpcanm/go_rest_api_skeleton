package config_

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
	mong := os.Getenv("MONGOURI")

	return mong
}
