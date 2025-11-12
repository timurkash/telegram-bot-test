package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6089659023:AAHz5Oct6l0Dw9Yr6fj4gCEgm-e6I-IHrGc")
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil && update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите кнопку:")
				keyboard := tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData("Кнопка 1", "btn1"),
						tgbotapi.NewInlineKeyboardButtonData("Кнопка 2", "btn2"),
					),
				)
				msg.ReplyMarkup = keyboard
				_, err := bot.Send(msg)
				if err != nil {
					return
				}
			}
		} else if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
			_, err := bot.Request(callback)
			if err != nil {
				return
			}

			var responseText string
			switch update.CallbackQuery.Data {
			case "btn1":
				responseText = "Вы нажали кнопку 1"
			case "btn2":
				responseText = "Вы нажали кнопку 2"
			default:
				responseText = "Неизвестная кнопка"
			}
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, responseText)
			_, err = bot.Send(msg)
			if err != nil {
				return
			}

		}
	}
}
