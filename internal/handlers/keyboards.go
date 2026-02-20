package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// buildMenu строит inline-клавиатуру по имени меню из YAML
func (h *Handler) buildMenu(name string) (tgbotapi.InlineKeyboardMarkup, bool) {
	menu, ok := h.cfg.Texts.Menus[name]
	if !ok {
		return tgbotapi.InlineKeyboardMarkup{}, false
	}

	var rows [][]tgbotapi.InlineKeyboardButton

	for _, yamlRow := range menu.Rows {

		var row []tgbotapi.InlineKeyboardButton

		for _, btn := range yamlRow {
			row = append(row,
				tgbotapi.NewInlineKeyboardButtonData(btn.Text, btn.CB),
			)
		}

		rows = append(rows, row)
	}

	return tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: rows,
	}, true
}