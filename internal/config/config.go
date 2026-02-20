package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

// ButtonDef соответствует одному объекту кнопки в YAML:
// {text: "...", cb: "..."}
type ButtonDef struct {
	Text string `yaml:"text"`
	CB   string `yaml:"cb"`
}

// MenuDef соответствует одному меню в YAML:
// menus:
//   main:
//     text: "start"
//     rows: ...
type MenuDef struct {
	// Text — это ключ в screens, который надо показать сверху меню
	// пример: text: "start" => screens.start
	Text string `yaml:"text"`

	// Rows — это строки кнопок:
	// rows:
	//   - [ {text: "...", cb: "..."}, {text: "...", cb: "..."} ]
	//   - [ {text: "...", cb: "..."} ]
	Rows [][]ButtonDef `yaml:"rows"`
}

// Texts — это весь content/config из texts.yaml
type Texts struct {
	// screens:
	//   start: "..."
	//   info:  "..."
	Screens map[string]string `yaml:"screens"`

	// menus:
	//   main: {text: "...", rows: ...}
	//   sex:  {text: "...", rows: ...}
	Menus map[string]MenuDef `yaml:"menus"`

	// replies:
	//   unknown: "..."
	//   hello: "..."
	//   how_are_you:
	//     - "..."
	Replies struct {
		Unknown   string   `yaml:"unknown"`
		Hello     string   `yaml:"hello"`
		HowAreYou []string `yaml:"how_are_you"`
	} `yaml:"replies"`
}

type Config struct {
	BotToken string
	Texts    Texts
}

func NewConfig() *Config {
	loadEnv()

	return &Config{
		BotToken: mustGetEnv("TOKEN_BOT"),
		Texts:    mustLoadTexts("texts.yaml"),
	}
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

	// Мини-валидация, чтобы ловить пустые файлы/ошибки структуры сразу
	if len(t.Screens) == 0 {
		log.Fatalf("texts.yaml: screens is empty or missing")
	}
	if len(t.Menus) == 0 {
		log.Fatalf("texts.yaml: menus is empty or missing")
	}
	if t.Replies.Unknown == "" {
		log.Fatalf("texts.yaml: replies.unknown is empty or missing")
	}

	return t
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env файл не найден, переменные окружения должны быть установлены вручную")
	}
}

func mustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s is not set", key)
	}
	return value
}