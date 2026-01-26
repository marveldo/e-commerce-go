package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host       string
	Port       string
	DBname     string
	DBUser     string
	DBPassword string
	JWTSecret  string
}

var config *Config

func LoadConfig() *Config {
	if config == nil {
		err := godotenv.Load()
		if err != nil {
			panic(fmt.Sprintf("Error loading .env file: %v", err))
		}
		config = &Config{
			Host:       os.Getenv("HOST"),
			Port:       os.Getenv("DB_PORT"),
			DBname:     os.Getenv("DB_NAME"),
			DBUser:     os.Getenv("DB_USER"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			JWTSecret:  os.Getenv("JWTSECRET"),
		}
	}

	return config

}
