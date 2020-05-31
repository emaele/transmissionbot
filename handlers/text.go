package handlers

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	transmission "github.com/hekmon/transmissionrpc"
)

// HandleText manages incoming text messages
func HandleText(text string, chatid int64, bot *tgbotapi.BotAPI, tc *transmission.Client) {
	_, err := tc.TorrentAdd(&transmission.TorrentAddPayload{
		Filename: &text,
	})
	if err != nil {
		log.Println(err)
		bot.Send(tgbotapi.NewMessage(chatid, fmt.Sprintf("There was an error adding your torrent file: %s", err.Error())))
		return
	}

	bot.Send(tgbotapi.NewMessage(chatid, "Added!"))
}
