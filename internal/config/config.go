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
	Google_Client_Id string
	Google_Client_Secret string
	Redis_Host string
	Redis_Port string
	Email string
	EmailPassword string
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
			Google_Client_Id: os.Getenv("GOOGLE_CLIENT_ID"),
			Google_Client_Secret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			Redis_Host : os.Getenv("REDIS_HOST"),
			Redis_Port: os.Getenv("REDIS_PORT"),
			Email: os.Getenv("GMAIL_SEND_EMAIL"),
			EmailPassword: os.Getenv("GMAIL_APP_PASSWORD"),
        }
	}

	return config

}
