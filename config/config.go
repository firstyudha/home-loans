package config

import (
	"log"
	"os"

	"github.com/subosito/gotenv"
)

type Config struct {
	AppName   string
	AppPort   int
	JWTSecret string
	DBDSN     string
}

func Init() *Config {
	defaultEnv := ".env"

	if err := gotenv.Load(defaultEnv); err != nil {
		log.Fatal("failed load .env")
	}

	log.SetOutput(os.Stdout)

	appConfig := &Config{
		AppName:   GetString("APP_NAME"),
		AppPort:   GetInt("APP_PORT"),
		JWTSecret: GetString("JWT_SECRET_KEY"),
		DBDSN:     GetString("DBDSN"),
	}

	return appConfig
}
