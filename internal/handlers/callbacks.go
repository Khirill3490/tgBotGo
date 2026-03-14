package handlers

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func getScreen(data string) (string, tgbotapi.InlineKeyboardMarkup) {
	screenPrefix := "screen:"
	menuPrefix := "menu:"
	testPrefix := "test:"
	var text string
	var keyboard tgbotapi.InlineKeyboardMarkup

	var command string

	if strings.HasPrefix(data, screenPrefix) {
		command = strings.TrimPrefix(data, screenPrefix)

		switch command {
		case "weather":
			text, keyboard = getWeatherMenu()
		case "info":
			text, keyboard = getInfoMenu()
		default:
			text = "Вы нажали неизвестную кнопку"
			keyboard = getBackButton()
		}

	} else if strings.HasPrefix(data, menuPrefix) {
		command = strings.TrimPrefix(data, menuPrefix)

		switch command {
		case "main":
			text, keyboard = getMainMenu()
		case "test_intro":
			text, keyboard = getTestIntroMenu()
		default:
			text = "Вы нажали неизвестную кнопку"
			keyboard = getBackButton()
		}
	} else if strings.HasPrefix(data, testPrefix) {
		command = strings.TrimPrefix(data, testPrefix)

		switch command {
		case "test_start":
			text, keyboard = getTestQuestion1Menu()
		case "test_afraid":
			text, keyboard = getTestAfraidScreen()
		case "test_stupid":
			text, keyboard = getTestStupidScreen()
		default:
			text = "Вы нажали неизвестную кнопку"
			keyboard = getBackButton()
		}
	} else {
		text = "Вы нажали неизвестную кнопку"
		keyboard = getBackButton()
	}

	return text, keyboard
}


func getMainMenu() (string, tgbotapi.InlineKeyboardMarkup) {
	text := "Главное меню. Выберите нужный раздел:"
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("ℹ️ О боте", "screen:info"),
			tgbotapi.NewInlineKeyboardButtonData("🌦 Погода", "screen:weather"),
			tgbotapi.NewInlineKeyboardButtonData("🧪 Тестовое интро", "menu:test_intro"),
		),
	)
	return text, keyboard
}

func getTestIntroMenu() (string, tgbotapi.InlineKeyboardMarkup) {
	text := "Это тестовое интро меню. Здесь будет информация о том, как пользоваться ботом."
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("ℹ️ Да", "test:test_start"),
			tgbotapi.NewInlineKeyboardButtonData("🌦 Нет, но да", "test:test_start"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("ℹ️ Я босюь", "test:test_afraid"),
			tgbotapi.NewInlineKeyboardButtonData("🌦 Я тупой", "test:test_stupid"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬅️ Назад", "menu:main"),
		),
	)
	return text, keyboard
}

func getTestQuestion1Menu() (string, tgbotapi.InlineKeyboardMarkup) {
	text := "Первый вопрос: Кто пернул?"
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Я", "test:test_q1_yes"),
			tgbotapi.NewInlineKeyboardButtonData("Не я", "test:test_q1_no"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Выйти из теста", "menu:main"),
		),
	)
	return text, keyboard
}

func getTestAfraidScreen() (string, tgbotapi.InlineKeyboardMarkup) {
	text := "Вы выбрали 'Я боюсь'. Не волнуйтесь, бот очень дружелюбный!"
	keyboard := getBackButton()
	return text, keyboard
}

func getTestStupidScreen() (string, tgbotapi.InlineKeyboardMarkup) {
	text := "Вы выбрали 'Я тупой'. Не переживайте, бот прост в использовании!"
	keyboard := getBackButton()
	return text, keyboard
}

func getInfoMenu() (string, tgbotapi.InlineKeyboardMarkup) {
	text := `Этот бот предоставляет информацию о погоде.
	Вы можете узнать текущую погоду, прогноз на неделю и многое другое!`
	keyboard := getBackButton()
	return text, keyboard
}

func getWeatherMenu() (string, tgbotapi.InlineKeyboardMarkup) {
	text := "В разработке... Скоро будет доступна информация о погоде!"
	keyboard := getBackButton()
	return text, keyboard
}

func getBackButton() tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬅️ Назад", "menu:main"),
		),
	)
	return keyboard
}