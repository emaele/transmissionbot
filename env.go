package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/odwrtw/transmission"
)

var (
	trAddress string
	trUser    string
	trPass    string

	telegramToken string

	tBot *tgbotapi.BotAPI
	tc   *transmission.Client
)

func loadEnv() {
	var ok bool

	trAddress, ok = os.LookupEnv("TRANSMISSION_ADDRESS")
	if !ok {
		log.Fatal("Can't load TRANSMISSION_ADDRESS var")
	}

	trUser, ok = os.LookupEnv("TRANSMISSION_USER")
	if !ok {
		log.Fatal("Can't load TRANSMISSION_USER var")
	}

	trPass, ok = os.LookupEnv("TRANSMISSION_PASS")
	if !ok {
		log.Fatal("Can't load TRANSMISSION_PASS var")
	}

	telegramToken, ok = os.LookupEnv("TELEGRAM_TOKEN")
	if !ok {
		log.Fatal("Can't load TELEGRAM_TOKEN var")
	}
}
