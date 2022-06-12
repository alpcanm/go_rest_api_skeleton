package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//! Güvenlik problemi yaratacak dataları .env dosyasına atıp gitignore a ekliyoruz ve github a yüklemek zorunda kalmıyoruz.
//! .env dosyası içerisindeki verileri çektiğimiz fonksiyonlar.
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
