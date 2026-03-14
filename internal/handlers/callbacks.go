package handlers

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func getScreen(data string) (string, tgbotapi.InlineKeyboardMarkup) {
	screenPrefix := "screen:"
	menuPrefix := "menu:"
	testPrefix := "test:"
	answerPrefix := "answer:"
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
		case "start":
			text, keyboard = getTestQuestion1Menu()
		case "q1":
			text, keyboard = getTestQuestion1Menu()
		case "q2":
			text, keyboard = getTestQuestion2Menu()
		case "afraid":
			text, keyboard = getTestAfraidScreen()
		case "stupid":
			text, keyboard = getTestStupidScreen()
		default:
			text = "Вы нажали неизвестную кнопку"
			keyboard = getBackButton()
		}
	} else if strings.HasPrefix(data, answerPrefix) {
		command = strings.TrimPrefix(data, answerPrefix)

		switch command {
		case "q1:a1":
			feedbackText := "Вы ответили 'Я'. Это правильный ответ!"
			text, keyboard = getTestQuestion2Menu()
			text = feedbackText + "\n\n" + text

		case "q1:a2":
			feedbackText := "Вы ответили 'Не я'. Это неправильный ответ."
			text, keyboard = getTestQuestion2Menu()
			text = feedbackText + "\n\n" + text
		case "q2:a1":
			feedbackText := "Вы ответили 'Да'. Это правильный ответ!"
			text, keyboard = getTestResultMenu()
			text = feedbackText + "\n\n" + text

		case "q2:a2":
			feedbackText := "Вы ответили 'Нет'. Это неправильный ответ."
			text, keyboard = getTestResultMenu()
			text = feedbackText + "\n\n" + text
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
	text := "Вы готовы начать тест?."
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("ℹ️ Начать тест", "test:start"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬅️ Назад", "menu:main"),
		),
	)

	return text, keyboard

	// keyboard := tgbotapi.NewInlineKeyboardMarkup(
	// 	tgbotapi.NewInlineKeyboardRow(
	// 		tgbotapi.NewInlineKeyboardButtonData("ℹ️ Да", "test:test_start"),
	// 		tgbotapi.NewInlineKeyboardButtonData("🌦 Нет, но да", "test:test_start"),
	// 	),
	// 	tgbotapi.NewInlineKeyboardRow(
	// 		tgbotapi.NewInlineKeyboardButtonData("ℹ️ Я босюь", "test:test_afraid"),
	// 		tgbotapi.NewInlineKeyboardButtonData("🌦 Я тупой", "test:test_stupid"),
	// 	),
	// 	tgbotapi.NewInlineKeyboardRow(
	// 		tgbotapi.NewInlineKeyboardButtonData("⬅️ Назад", "menu:main"),
	// 	),
	// )
}

func getTestQuestion1Menu() (string, tgbotapi.InlineKeyboardMarkup) {
	text := "Первый вопрос: Кто пернул?"
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Я", "answer:q1:a1"),
			tgbotapi.NewInlineKeyboardButtonData("Не я", "answer:q1:a2"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Выйти из теста", "menu:main"),
		),
	)
	return text, keyboard
}

func getTestQuestion2Menu() (string, tgbotapi.InlineKeyboardMarkup) {
	text := "Второй вопрос"
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Да", "answer:q2:a1"),
			tgbotapi.NewInlineKeyboardButtonData("Нет", "answer:q2:a2"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Выйти из теста", "menu:main"),
		),
	)
	return text, keyboard
}

func getTestResultMenu() (string, tgbotapi.InlineKeyboardMarkup) {
	text := "Результаты теста: Вы молодец!"
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
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
