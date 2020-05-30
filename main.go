package main

import (
	"fmt"
	"log"

	"github.com/emaele/transmissionbot/utility"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	loadEnv()

	//configure polling
	updates := tBot.GetUpdatesChan(tgbotapi.UpdateConfig{Timeout: 60})

	for update := range updates {
		if update.Message != nil {
			go mainBot(update.Message)
		}
	}

}
func mainBot(message *tgbotapi.Message) {
	if message.Text != "" {
		// Handle text
	}

	if message.Document != nil {
		fileURL, err := tBot.GetFileDirectURL(message.Document.FileID)
		if err != nil {
			log.Println(err)
		}

		path, err := utility.DownloadFile(message.Document.FileName, fileURL)
		if err != nil {
			log.Println(err)
		}

		_, err = tc.Add(path)
		if err != nil {
			log.Println(err)
			tBot.Send(tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("There was an error adding your torrent file: %s", err.Error())))
		}

		tBot.Send(tgbotapi.NewMessage(message.Chat.ID, "Added!"))
	}
}
