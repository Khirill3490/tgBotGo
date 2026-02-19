package main

import (
	"log"

	"github.com/Khirill3490/weatherBot/internal/bot"
	"github.com/Khirill3490/weatherBot/internal/config"
	"github.com/Khirill3490/weatherBot/internal/handlers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

/*
main.go теперь делает только:
1) загрузка конфига
2) инициализация Telegram API клиента
3) регистрация команд
4) запуск обработчика
*/
func main() {
	cfg := config.NewConfig()

	tgBot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		log.Fatalf("failed to init bot: %v", err)
	}

	tgBot.Debug = false
	log.Printf("Authorized on account %s", tgBot.Self.UserName)

	// Системное меню Telegram: /start /help
	bot.RegisterCommands(tgBot)

	// Запускаем обработчик событий
	h := handlers.New(tgBot, cfg)
	h.Run()
}
