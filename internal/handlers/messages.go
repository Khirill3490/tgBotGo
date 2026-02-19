package handlers

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handler) handleMessage(m *tgbotapi.Message) {
	text := strings.ToLower(strings.TrimSpace(m.Text))
	chatID := m.Chat.ID

	log.Printf("[%s] %s", m.From.UserName, text)

	// Если это команда — обрабатываем отдельно
	if strings.HasPrefix(text, "/") {
		h.handleCommand(chatID, text)
		return
	}

	// Если не команда — обычный текст
	reply := h.buildReply(text)

	msg := tgbotapi.NewMessage(chatID, reply)
	msg.ReplyMarkup = h.mainInlineMenu()
	_, _ = h.tgBot.Send(msg)
}

/*
buildReply — простейшая бизнес-логика.
Пока это switch, позже можно заменить на полноценные handlers/routers.
*/
func (h *Handler) buildReply(text string) string {
	switch text {
	case "привет", "hello", "hi":
		return h.cfg.Texts.Replies.Hello

	case "как дела?":
		return h.cfg.Texts.Replies.HowAreYou

	default:
		return h.cfg.Texts.Replies.Unknown
	}
}


func (h *Handler) handleCommand(chatID int64, cmd string) {

	var text string

	switch cmd {

	case "/start", "/старт":
		text = h.cfg.Texts.StartText

	case "/help", "/помощь":
		text = h.cfg.Texts.Screens.Help

	case "/info", "/инфо":
		text = h.cfg.Texts.Screens.Info

	case "/sex", "/секс":
		text = h.cfg.Texts.Screens.Sex

	case "/hulk", "/халк":
		text = h.cfg.Texts.Screens.Hulk

	default:
		text = h.cfg.Texts.Replies.Unknown
	}

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = h.mainInlineMenu()
	_, _ = h.tgBot.Send(msg)
}
