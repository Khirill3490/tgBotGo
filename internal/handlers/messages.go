package handlers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handler) EditMessage(
	chatID int64,
	messageID int,
	text string,
	keyboard tgbotapi.InlineKeyboardMarkup,
) {
	msg := tgbotapi.NewEditMessageText(chatID, messageID, text)

	msg.ReplyMarkup = &keyboard

	_, err := h.bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func (h *Handler) SendMessage(
	chatID int64,
	text string,
	keyboard tgbotapi.InlineKeyboardMarkup,
) {
	msg := tgbotapi.NewMessage(chatID, text)

	msg.ReplyMarkup = keyboard

	_, err := h.bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}
