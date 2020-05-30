package handlers

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/odwrtw/transmission"
)

func HandleText(text string, chatid int64, bot *tgbotapi.BotAPI, tc *transmission.Client) {
	_, err := tc.Add(text)
	if err != nil {
		log.Println(err)
		bot.Send(tgbotapi.NewMessage(chatid, fmt.Sprintf("There was an error adding your torrent file: %s", err.Error())))
		return
	}

	bot.Send(tgbotapi.NewMessage(chatid, "Added!"))
}
