package handlers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

/*
Callback-ключи — это НЕ текст кнопок.
Это внутренние идентификаторы (роуты), которые приходят при нажатии inline-кнопки.
*/
const (
	cbHelp    = "HELP"
	cbInfo    = "INFO"
	cbWeather = "WEATHER"
	cbBack    = "BACK"
	cbSex     = "SEX"
	cbHulk    = "HULK"
)

func (h *Handler) handleCallback(cq *tgbotapi.CallbackQuery) {
	// Telegram “ждёт”, что мы подтвердим нажатие (иначе будет крутиться индикатор загрузки).
	ack := tgbotapi.NewCallback(cq.ID, "")
	if _, err := h.tgBot.Request(ack); err != nil {
		log.Printf("callback ack error: %v", err)
	}

	chatID := cq.Message.Chat.ID
	messageID := cq.Message.MessageID

	// По умолчанию вернём главное меню
	text := h.cfg.Texts.StartText
	markup := h.mainInlineMenu()

	switch cq.Data {
	case cbHelp:
		text = h.cfg.Texts.Screens.Help
		markup = h.backInlineMenu()

	case cbInfo:
		text = h.cfg.Texts.Screens.Info
		markup = h.backInlineMenu()
	case cbSex:
		text = h.cfg.Texts.Screens.Sex
		markup = h.backInlineMenu()
	case cbHulk:
		text = h.cfg.Texts.Screens.Hulk
		markup = h.backInlineMenu()
	case cbBack:
		text = h.cfg.Texts.StartText
		markup = h.mainInlineMenu()

	default:
		text = h.cfg.Texts.Replies.Unknown
		markup = h.mainInlineMenu()
	}

	/*
		Важно:
		Мы редактируем текущее сообщение, а не отправляем новое.
		Так интерфейс не засоряется лишними сообщениями.
	*/
	edit := tgbotapi.NewEditMessageText(chatID, messageID, text)
	edit.ReplyMarkup = &markup

	if _, err := h.tgBot.Send(edit); err != nil {
		// Если редактирование недоступно — отправим новое сообщение как запасной вариант
		log.Printf("edit message error: %v", err)
		msg := tgbotapi.NewMessage(chatID, text)
		msg.ReplyMarkup = markup
		_, _ = h.tgBot.Send(msg)
	}
}
