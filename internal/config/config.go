package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Texts struct {
	StartText string `yaml:"start_text"`

	Replies struct {
		Unknown   string `yaml:"unknown"`
		Hello     string `yaml:"hello"`
		HowAreYou string `yaml:"how_are_you"`
	} `yaml:"replies"`

	Buttons struct {
		Weather string `yaml:"weather"`
		Help    string `yaml:"help"`
		Info    string `yaml:"info"`
		Back    string `yaml:"back"`
		Sex     string `yaml:"sex"`
		Hulk    string `yaml:"hulk"`
	} `yaml:"buttons"`

	Screens struct {
		Help string `yaml:"help"`
		Info string `yaml:"info"`
		Sex  string `yaml:"sex"`
		Hulk string `yaml:"hulk"`
	} `yaml:"screens"`
}

type Config struct {
	BotToken string
	Texts    Texts
}

func NewConfig() *Config {
	loadEnv()

	cfg := &Config{
		BotToken: getEnv("TOKEN_BOT"),
		Texts:    mustLoadTexts("texts.yaml"),
	}

	return cfg
}

func mustLoadTexts(path string) Texts {
	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read texts file %s: %v", path, err)
	}

	var t Texts
	if err := yaml.Unmarshal(b, &t); err != nil {
		log.Fatalf("failed to parse texts yaml %s: %v", path, err)
	}

	return t
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
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
