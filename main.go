package main

import (
	"fmt"
	"log"

	"/home/asmolin/go/bot/botdbconnect"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("7016522405:AAEw9DeZs9bOHg3cl-iLuF1MLcG9UwjP7S0")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		reply := "Не знаю что сказать азазазаз "
		if update.Message == nil {
			continue
		}
		userid := update.Message.From.ID

		switch update.Message.Command() {
		case "start":
			check := botdbconnect.Checkuserifexist(userid)
			fmt.Print(check)
			reply = "Привет. Я телеграм-бот"
		case "hello":
			reply = "world"
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)

		bot.Send(msg)

	}
}
