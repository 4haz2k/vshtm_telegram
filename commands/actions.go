package commands

import (
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"io"
	"log"
	"net/http"
	"strings"

	"vshtm_telegram/database"
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

func Subscribe(chatId int64) (string, tgbotapi.ReplyKeyboardMarkup) {
	database.SubscribeUser(chatId)
	return resources.Lang("subscribedMessage"), subscribeKeyboard
}

func Unsubscribe(chatId int64) (string, tgbotapi.ReplyKeyboardMarkup) {
	database.UnsubscribeUser(chatId)
	return resources.Lang("unsubscribedMessage"), unsubscribeKeyboard
}

func Schedule() string {
	list := database.GetSchedule()

	message := resources.Lang("ScheduleForWeekMessage")
	for _, item := range list {
		if item.Building.Valid {
			message += fmt.Sprintf(resources.Lang("lanScheduleMessage"),
				item.Subject, item.Theme, item.CreatedAt.Time.Format("02.01.2006 15:04"),
				item.Teacher, item.Building, item.Link)
		} else {
			message += fmt.Sprintf(resources.Lang("onlineScheduleMessage"),
				item.Subject, item.Theme, item.CreatedAt.Time.Format("02.01.2006 15:04"),
				item.Teacher, item.Link)
		}
	}

	return message
}

func UnknownMessage() string {
	resp, err := http.Get("http://rzhunemogu.ru/RandJSON.aspx?CType=1")
	if err != nil {
		log.Fatal(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	var joke string

	for {
		bs := make([]byte, 1014)
		n, err := resp.Body.Read(bs)

		joke += string(bs[:n])

		if n == 0 || err != nil {
			break
		}
	}

	// converting encoding
	decoder := charmap.Windows1251.NewDecoder()
	joke, _ = decoder.String(joke)

	// replacing json chars
	replacer := strings.NewReplacer("{\"content\":\"", "", "\"}", "")

	return fmt.Sprintf(resources.Lang("unknownMessage"), replacer.Replace(joke))
}
