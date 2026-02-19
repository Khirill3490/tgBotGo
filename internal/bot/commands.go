package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

/*
RegisterCommands — регистрирует список команд бота в Telegram.

Эти команды отображаются:
- когда пользователь нажимает "/" в поле ввода
- в системном меню Telegram

Важно:
Telegram хранит этот список У СЕБЯ.
Мы просто отправляем его один раз при старте приложения.
*/
func RegisterCommands(api *tgbotapi.BotAPI) {

	// Описываем команды, которые увидит пользователь
	commands := tgbotapi.NewSetMyCommands(
		tgbotapi.BotCommand{
			Command:     "start",
			Description: "Запустить бота",
		},
		tgbotapi.BotCommand{
			Command:     "help",
			Description: "Помощь по использованию",
		},
		tgbotapi.BotCommand{
			Command:     "info",
			Description: "Информация о боте",
		},
		tgbotapi.BotCommand{
			Command:     "hulk",
			Description: "Показать Халка",
		},
		tgbotapi.BotCommand{
			Command:     "sex",
			Description: "Показать секс возможности",
		},
	)

	/*
		Request — это прямой вызов Bot API.
		Мы говорим Telegram:
		"Сохрани вот этот список команд для моего бота".
	*/
	if _, err := api.Request(commands); err != nil {
		log.Fatalf("Не удалось зарегистрировать команды бота: %v", err)
	}
}
