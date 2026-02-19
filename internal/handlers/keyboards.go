package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

/*
mainInlineMenu — главное меню кнопок под сообщением.
Тексты кнопок берём из YAML, а callback-данные — из констант cb*.
*/
func (h *Handler) mainInlineMenu() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(h.cfg.Texts.Buttons.Weather, cbWeather),
			tgbotapi.NewInlineKeyboardButtonData(h.cfg.Texts.Buttons.Help, cbHelp),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(h.cfg.Texts.Buttons.Info, cbInfo),
			tgbotapi.NewInlineKeyboardButtonData(h.cfg.Texts.Buttons.Sex, cbSex),
			tgbotapi.NewInlineKeyboardButtonData(h.cfg.Texts.Buttons.Hulk, cbHulk),
		),
	)
}

/*
backInlineMenu — кнопка "Назад" для внутренних экранов.
*/
func (h *Handler) backInlineMenu() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(h.cfg.Texts.Buttons.Back, cbBack),
		),
	)
}
