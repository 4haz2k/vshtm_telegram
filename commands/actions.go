package commands

import (
	"vshtm_telegram/resources"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var startKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(resources.Lang("subscribeLabel")),
		tgbotapi.NewKeyboardButton(resources.Lang("unsubscribeLabel")),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(resources.Lang("scheduleLabel")),
	),
)

var subscribeKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(resources.Lang("unsubscribeLabel")),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(resources.Lang("scheduleLabel")),
	),
)

var unsubscribeKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(resources.Lang("subscribeLabel")),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(resources.Lang("scheduleLabel")),
	),
)

func Start() (string, tgbotapi.ReplyKeyboardMarkup) {
	return resources.Lang("startMessage"), startKeyboard
}

func Subscribe() (string, tgbotapi.ReplyKeyboardMarkup) {
	return resources.Lang("startMessage"), subscribeKeyboard
}

func Unsubscribe() (string, tgbotapi.ReplyKeyboardMarkup) {
	return resources.Lang("unsubscribedMessage"), unsubscribeKeyboard
}
