package config

import (
	"fmt"
	"os"
	"strconv"

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
	Paystack_Secret_Key string
	Request_Per_Second int
	Request_Burst int
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
			Paystack_Secret_Key : os.Getenv("PAYSTACK_SECRET_KEY"),
			Request_Per_Second: func () int {
				rps, err := strconv.Atoi(os.Getenv("REQUEST_PER_SECOND"))
				if err != nil {
					panic(fmt.Sprintf("Error converting REQUEST_PER_SECOND to int: %v", err))
				}
				return rps
			}(),
			Request_Burst: func () int {
				rb, err := strconv.Atoi(os.Getenv("REQUEST_BURST"))
				if err != nil {
					panic(fmt.Sprintf("Error converting REQUEST_BURST to int: %v", err))
				}
				return rb
			}(),
        }
	}

	return config

}
