package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	"vshtm_telegram/commands"
	"vshtm_telegram/database"
	"vshtm_telegram/resources"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func init() {
	if os.Getenv("APP_ENV") != "prod" {
		err := godotenv.Load()
		if err != nil {
			log.Panic("environment file not found")
		}
	}
}

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	wh, _ := tgbotapi.NewWebhook(os.Getenv("TELEGRAM_WEBHOOK_URL"))

	_, err = bot.Request(wh)
	if err != nil {
		log.Fatal(err)
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}

	if info.LastErrorDate != 0 {
		log.Printf(resources.Error("callbackFailed"), info.LastErrorMessage)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	go func() {
		err := http.ListenAndServe("0.0.0.0:8443", nil)
		if err != nil {
			log.Panic(err)
		}
	}()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		database.LogMessage(update.Message.Chat.ID, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Text {
		case resources.Command("startSlash"):
		case resources.Command("start"):
			message, keyboard := commands.Start()
			msg.Text = message
			msg.ReplyMarkup = keyboard
			break
		case resources.Command("subscribe"):
			message, keyboard := commands.Subscribe(update.Message.Chat.ID)
			msg.Text = message
			msg.ReplyMarkup = keyboard
			break
		case resources.Command("unsubscribe"):
			message, keyboard := commands.Unsubscribe(update.Message.Chat.ID)
			msg.Text = message
			msg.ReplyMarkup = keyboard
			break
		case resources.Command("schedule"):
			message := commands.Schedule()
			msg.Text = message
			break
		default:
			message := commands.UnknownMessage()
			msg.Text = message
			break
		}

		msg.ParseMode = "html"

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
