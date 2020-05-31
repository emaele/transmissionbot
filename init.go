package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hekmon/transmissionrpc"
)

func init() {
	loadEnv()

	var err error
	tBot, err = tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		log.Fatal(err)
	}

	adv := &transmissionrpc.AdvancedConfig{
		HTTPS: trHTTPS,
		Port:  trPort,
	}

	tc, err = transmissionrpc.New(trAddress, trUser, trPass, adv)
	if err != nil {
		log.Fatal(err)
	}
}
