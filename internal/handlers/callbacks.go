package handlers

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	cbMenuPrefix   = "menu:"
	cbScreenPrefix = "screen:"
)

func (h *Handler) handleCallback(cq *tgbotapi.CallbackQuery) {
	// 1) Telegram ждёт подтверждение нажатия (иначе будет "крутилка")
	ack := tgbotapi.NewCallback(cq.ID, "")
	if _, err := h.tgBot.Request(ack); err != nil {
		log.Printf("callback ack error: %v", err)
	}

	chatID := cq.Message.Chat.ID
	messageID := cq.Message.MessageID

	// 2) Разбираем действие из callback_data
	text, markup := h.resolveCallbackAction(cq.Data)

	// 3) Редактируем текущее сообщение (как у тебя и было)
	edit := tgbotapi.NewEditMessageText(chatID, messageID, text)
	edit.ReplyMarkup = &markup

	if _, err := h.tgBot.Send(edit); err != nil {
		// Если редактирование недоступно — отправим новое сообщение
		log.Printf("edit message error: %v", err)
		msg := tgbotapi.NewMessage(chatID, text)
		msg.ReplyMarkup = markup
		_, _ = h.tgBot.Send(msg)
	}
}

// resolveCallbackAction понимает только 2 вида callback:
// - menu:<name>   -> открыть меню из YAML
// - screen:<key>  -> показать экран из YAML
func (h *Handler) resolveCallbackAction(data string) (string, tgbotapi.InlineKeyboardMarkup) {
	switch {
	case strings.HasPrefix(data, cbMenuPrefix):
		menuName := strings.TrimPrefix(data, cbMenuPrefix)
		return h.renderMenu(menuName)

	case strings.HasPrefix(data, cbScreenPrefix):
		screenKey := strings.TrimPrefix(data, cbScreenPrefix)
		return h.renderScreen(screenKey)

	default:
		// неизвестный callback
		text := h.cfg.Texts.Replies.Unknown
		markup, ok := h.buildMenu("main")
		if !ok {
			return text, tgbotapi.NewInlineKeyboardMarkup()
		}
		return text, markup
	}
}

// renderMenu показывает текст меню + клавиатуру меню.
// menu.Text — это ключ в screens.
func (h *Handler) renderMenu(menuName string) (string, tgbotapi.InlineKeyboardMarkup) {
	menu, ok := h.cfg.Texts.Menus[menuName]
	if !ok {
		// нет такого меню в YAML -> неизвестное
		text := h.cfg.Texts.Replies.Unknown
		markup, ok2 := h.buildMenu("main")
		if !ok2 {
			return text, tgbotapi.NewInlineKeyboardMarkup()
		}
		return text, markup
	}

	// Берём текст по ключу menu.Text (пример: "start" -> screens["start"])
	text := h.cfg.Texts.Screens[menu.Text]
	if text == "" {
		text = h.cfg.Texts.Replies.Unknown
	}

	// Строим клавиатуру из YAML
	markup, ok := h.buildMenu(menuName)
	if !ok {
		markup = tgbotapi.NewInlineKeyboardMarkup()
	}

	return text, markup
}

// renderScreen показывает один экран (screens[key]) + кнопку "Назад".
// Пока делаем упрощение: "Назад" всегда ведёт в главное меню.
func (h *Handler) renderScreen(screenKey string) (string, tgbotapi.InlineKeyboardMarkup) {
	text := h.cfg.Texts.Screens[screenKey]
	if text == "" {
		text = h.cfg.Texts.Replies.Unknown
	}

	back := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬅️ Назад", "menu:main"),
		),
	)

	return text, back
}