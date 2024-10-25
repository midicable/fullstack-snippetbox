package config

import (
	"fullstack-snippetbox-backend/pkg/utils"
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost string
	DBPort string
	DBUser string
	DBPassword string
	DBName string
}

var appConfig Config

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	appConfig = Config{
		DBHost: utils.GetEnv("DB_HOST", ""),
		DBPort: utils.GetEnv("DB_POST", ""),
		DBUser: utils.GetEnv("DB_USER", ""),
		DBPassword: utils.GetEnv("DB_PASSWORD", ""),
		DBName: utils.GetEnv("DB_NAME", ""),
	}
}

func GetAppConfig() Config {
	return appConfig
}