package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_HOST     string
	DB_PORT     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	JWT_SECRET  []byte
}

var Envs = InitConfig()

func InitConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Print(err.Error())
	}

	return &Config{
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),
		JWT_SECRET:  []byte(os.Getenv("JWT_SECRET")),
	}
}
