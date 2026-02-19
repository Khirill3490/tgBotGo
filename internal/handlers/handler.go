package handlers

import (
	"log"

	"github.com/Khirill3490/weatherBot/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

/*
Handler — это "центр" обработки событий Telegram.

Мы храним здесь:
- tgBot: клиент Telegram API
- cfg: конфигурация (включая тексты из YAML)

Так удобнее, чем передавать cfg и bot в каждую функцию.
*/
type Handler struct {
	tgBot *tgbotapi.BotAPI
	cfg   *config.Config
}

// New создает обработчик и сохраняет зависимости внутри.
func New(tgBot *tgbotapi.BotAPI, cfg *config.Config) *Handler {
	return &Handler{
		tgBot: tgBot,
		cfg:   cfg,
	}
}

/*
Run запускает основной цикл long polling и маршрутизирует события:
- Message → handleMessage
- CallbackQuery → handleCallback
*/
func (h *Handler) Run() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := h.tgBot.GetUpdatesChan(u)

	for update := range updates {
		if update.CallbackQuery != nil {
			h.handleCallback(update.CallbackQuery)
			continue
		}

		if update.Message != nil {
			h.handleMessage(update.Message)
			continue
		}
	}

	log.Println("Канал обновлений закрыт")
}
