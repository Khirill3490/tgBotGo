package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	cfg "github.com/Khirill3490/weatherBot/internal/config"
)

type Handler struct {
	bot *tgbotapi.BotAPI
	cfg *cfg.Config
}

func New(bot *tgbotapi.BotAPI, cfg *cfg.Config) *Handler {
	return &Handler{
		bot: bot,
		cfg: cfg,
	}
}

func (h *Handler) Run() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := h.bot.GetUpdatesChan(u)

	for update := range updates {

		if update.Message != nil {
			if update.Message.Text == "/start" {
				var text string
				var keyboard tgbotapi.InlineKeyboardMarkup
				chatID := update.Message.Chat.ID

				text, keyboard = getScreen("menu:main")

				h.SendMessage(chatID, text, keyboard)

			}
		}

		if update.CallbackQuery != nil {
			chatID := update.CallbackQuery.Message.Chat.ID
			messageID := update.CallbackQuery.Message.MessageID
			var text string
			var keyboard tgbotapi.InlineKeyboardMarkup

			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
			_, _ = h.bot.Request(callback)

			data := update.CallbackQuery.Data

			text, keyboard = getScreen(data)

			h.EditMessage(chatID, messageID, text, keyboard)

		}
	}
}