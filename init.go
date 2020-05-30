package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/odwrtw/transmission"
)

func init() {
	loadEnv()

	var err error
	tBot, err = tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		log.Fatal(err)
	}

	// Fill transmission config
	conf := transmission.Config{
		Address:  trAddress,
		User:     trUser,
		Password: trPass,
	}

	tc, err = transmission.New(conf)
	if err != nil {
		log.Fatal(err)
	}
}
