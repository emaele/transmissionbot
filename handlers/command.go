package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/odwrtw/transmission"
)

// HandleCommand handles incoming bot commands
func HandleCommand(text string, chatid int64, bot *tgbotapi.BotAPI, tc *transmission.Client) {
	switch text {
	case "start":
		bot.Send(tgbotapi.NewMessage(chatid, "Hello! Send a me torrent file or a magnet."))
	}
}
