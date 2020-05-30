package handlers

import (
	"fmt"
	"log"

	"github.com/emaele/transmissionbot/utility"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/odwrtw/transmission"
)

func HandleDocument(document tgbotapi.Document, chatid int64, bot *tgbotapi.BotAPI, tc *transmission.Client) {
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

	_, err = tc.Add(path)
	if err != nil {
		log.Println(err)
		bot.Send(tgbotapi.NewMessage(chatid, fmt.Sprintf("There was an error adding your torrent file: %s", err.Error())))
		return
	}

	bot.Send(tgbotapi.NewMessage(chatid, "Added!"))
}