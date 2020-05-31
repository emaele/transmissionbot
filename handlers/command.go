package handlers

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	transmission "github.com/hekmon/transmissionrpc"
)

// HandleCommand handles incoming bot commands
func HandleCommand(text string, chatid int64, bot *tgbotapi.BotAPI, tc *transmission.Client) {
	switch text {
	case "start":
		bot.Send(tgbotapi.NewMessage(chatid, "Hello! Send a me torrent file or a magnet."))
	case "mydownloads":
		text, err := getActiveTorrents(tc)
		if err != nil {
			bot.Send(tgbotapi.NewMessage(chatid, fmt.Sprintf("There was an error with your request: %s", err.Error())))
		}
		bot.Send(tgbotapi.NewMessage(chatid, text))
	}

}

func getActiveTorrents(tc *transmission.Client) (res string, err error) {

	torrents, err := tc.TorrentGetAll()
	if err != nil {
		log.Println(err)
		return "", err
	}

	for _, t := range torrents {
		res += fmt.Sprintf("%s - %f\n", *t.Name, *t.PercentDone)
	}

	return res, nil
}
