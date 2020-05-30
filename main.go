package main

import (
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
			// TODO: handle err
		}

		path, err := utility.DownloadFile(message.Document.FileName, fileURL)
		if err != nil {
			// TODO: handle err
		}

		_, err = tc.Add(path)
		if err != nil {
			// TODO: handle err
		}

		tBot.Send(tgbotapi.NewMessage(message.Chat.ID, "Added!"))
	}
}
