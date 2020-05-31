package main

import (
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	transmission "github.com/hekmon/transmissionrpc"
)

var (
	trAddress string
	trUser    string
	trPass    string
	trPort    uint16
	trHTTPS   bool

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

	trPortt, ok := os.LookupEnv("TRANSMISSION_PORT")
	if !ok {
		trPort = 9091
	} else {
		port, err := strconv.Atoi(trPortt)
		if err != nil {
			log.Fatal(err)
		}
		trPort = uint16(port)
	}

	trHTTPSt, ok := os.LookupEnv("TRANSMISSION_HTTPS")
	if !ok {
		trHTTPS = false
	} else {
		trHTTPS = boolFromString(trHTTPSt)
	}

	telegramToken, ok = os.LookupEnv("TELEGRAM_TOKEN")
	if !ok {
		log.Fatal("Can't load TELEGRAM_TOKEN var")
	}
}

func boolFromString(boo string) bool {
	if boo == "true" {
		return true
	}

	return false
}
