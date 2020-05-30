package main

import (
	"github.com/emaele/transmissionbot/handlers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	//configure polling
	updates := tBot.GetUpdatesChan(tgbotapi.UpdateConfig{Timeout: 60})

	for update := range updates {
		if update.Message != nil {
			go mainBot(update.Message)
		}
	}

}
func mainBot(message *tgbotapi.Message) {

	if message.IsCommand() {
		handlers.HandleCommand(message.Command(), message.Chat.ID, tBot, tc)
		return
	}

	if message.Text != "" {
		handlers.HandleText(message.Text, message.Chat.ID, tBot, tc)
	}

	if message.Document != nil {
		handlers.HandleDocument(*message.Document, message.Chat.ID, tBot, tc)
	}
}
