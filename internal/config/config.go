package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BotToken string
}

// New загружает конфигурацию приложения
func NewConfig() *Config {
	loadEnv()

	cfg := &Config{
		BotToken: getEnv("TOKEN_BOT"),
	}

	return cfg
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env файл не найден, переменные окружения должны быть установлены вручную")
	}
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s is not set", key)
	}
	return value
}
