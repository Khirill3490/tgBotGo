package config

import (
	"errors"
	"os"
)

type Config struct {
	Token string
}

func NewConfig() (*Config, error) {
	token := os.Getenv("TOKEN_BOT")

	if token == "" {
		return nil, errors.New("TOKEN_BOT переменная окружения не установлена")
	}
	
	return &Config{
		Token: token,
	}, nil
}