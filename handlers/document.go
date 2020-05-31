package handlers

import (
	"fmt"
	"log"
	"strings"

	"github.com/emaele/transmissionbot/utility"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	transmission "github.com/hekmon/transmissionrpc"
)

// HandleDocument handles incoming files
func HandleDocument(document tgbotapi.Document, chatid int64, bot *tgbotapi.BotAPI, tc *transmission.Client) {

	if !strings.HasSuffix(document.FileName, ".torrent") {
		bot.Send(tgbotapi.NewMessage(chatid, "Please send me a valid .torrent file"))
		return
	}

	fileURL, err := bot.GetFileDirectURL(document.FileID)
	if err != nil {
		log.Println(err)
		bot.Send(tgbotapi.NewMessage(chatid, fmt.Sprintf("There was an error adding your torrent file: %s", err.Error())))
		return
	}

	path, err := utility.DownloadFile(document.FileName, fileURL)
	if err != nil {
		log.Println(err)
		bot.Send(tgbotapi.NewMessage(chatid, fmt.Sprintf("There was an error adding your torrent file: %s", err.Error())))
		return
	}

	_, err = tc.TorrentAddFile(path)
	if err != nil {
		log.Println(err)
		bot.Send(tgbotapi.NewMessage(chatid, fmt.Sprintf("There was an error adding your torrent file: %s", err.Error())))
		return
	}

	bot.Send(tgbotapi.NewMessage(chatid, "Added!"))
}
