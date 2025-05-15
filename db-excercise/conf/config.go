package conf

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_HOST string
	DB_PORT string
	DB_USER string
	DB_PASS string
	DB_NAME string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	return &Config{
		DB_HOST: os.Getenv("DB_HOST"),
		DB_PORT: os.Getenv("DB_PORT"),
		DB_USER: os.Getenv("DB_USER"),
		DB_PASS: os.Getenv("DB_PASSWORD"),
		DB_NAME: os.Getenv("DB_NAME"),
	}, nil
}
