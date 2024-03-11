package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"vshtm_telegram/commands"
	"vshtm_telegram/database"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Panic("Environment file not found")
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
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
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
		case "start":
		case "/start":
			message, keyboard := commands.Start()
			msg.Text = message
			msg.ReplyMarkup = keyboard
			break
		case "Подписаться":
			message, keyboard := commands.Subscribe(update.Message.Chat.ID)
			msg.Text = message
			msg.ReplyMarkup = keyboard
			break
		case "Отписаться":
			message, keyboard := commands.Unsubscribe(update.Message.Chat.ID)
			msg.Text = message
			msg.ReplyMarkup = keyboard
			break
		case "Расписание на неделю":
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
