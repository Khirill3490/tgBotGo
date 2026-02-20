package handlers

import (
	"log"
	"math/rand"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// если у тебя уже где-то есть seed — можно убрать init()
func init() {
	rand.Seed(time.Now().UnixNano())
}

func (h *Handler) handleMessage(msg *tgbotapi.Message) {
	chatID := msg.Chat.ID

	// 1) Команды вида /start, /help
	if msg.IsCommand() {
		h.handleCommand(chatID, msg.Command())
		return
	}

	// 2) Обычный текст
	text := strings.TrimSpace(strings.ToLower(msg.Text))
	if text == "" {
		return
	}

	replyText, replyMarkup := h.buildTextReply(text)

	out := tgbotapi.NewMessage(chatID, replyText)
	if replyMarkup != nil {
		out.ReplyMarkup = replyMarkup
	}

	if _, err := h.tgBot.Send(out); err != nil {
		log.Printf("send message error: %v", err)
	}
}

// msg.Command() возвращает команду БЕЗ слэша: "start", "help", ...
func (h *Handler) handleCommand(chatID int64, cmd string) {
	var text string
	var markup tgbotapi.InlineKeyboardMarkup
	var ok bool

	switch strings.ToLower(cmd) {
	case "start", "старт":
		text, markup = h.renderMenu("main")

	case "sex", "секс":
		text, markup = h.renderMenu("sex")

	case "info", "инфо":
		text, markup = h.renderScreen("info")

	case "help", "помощь":
		text, markup = h.renderScreen("help")

	case "weather", "погода":
		text, markup = h.renderScreen("weather")

	case "hulk", "халк":
		text, markup = h.renderScreen("hulk")

	default:
		text = h.cfg.Texts.Replies.Unknown
		markup, ok = h.buildMenu("main")
		if !ok {
			markup = tgbotapi.NewInlineKeyboardMarkup()
		}
	}

	out := tgbotapi.NewMessage(chatID, text)
	out.ReplyMarkup = markup

	if _, err := h.tgBot.Send(out); err != nil {
		log.Printf("send command message error: %v", err)
	}
}

// buildTextReply — ответы на обычные сообщения без команд
// Возвращает (text, markup). markup может быть nil — значит без клавиатуры.
func (h *Handler) buildTextReply(normalized string) (string, *tgbotapi.InlineKeyboardMarkup) {
	switch normalized {
	case "привет", "hi", "hello":
		return h.cfg.Texts.Replies.Hello, nil

	case "как дела", "как дела?", "как ты", "как ты?":
		answers := h.cfg.Texts.Replies.HowAreYou
		if len(answers) == 0 {
			return "Нормально 🙂", nil
		}
		return answers[rand.Intn(len(answers))], nil

	// Можешь добавить другие "триггеры" по желанию

	default:
		// Можно в unknown добавлять кнопку "открыть меню"
		m, ok := h.buildMenu("main")
		if ok {
			return h.cfg.Texts.Replies.Unknown, &m
		}
		return h.cfg.Texts.Replies.Unknown, nil
	}
}