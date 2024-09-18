package main

import (
	"fmt"
	"log"
	"strconv"

	"botdbconnect"

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
		// userid := update.Message.From.ID
		// var convertuserid string
		// fmt.Print(convertuserid)
		// convertuserid = strconv.FormatInt(userid, 10)
		switch update.Message.Command() {
		case "start":
			userid := update.Message.From.ID
			var convertuserid string
			convertuserid = strconv.FormatInt(userid, 10)
			check := botdbconnect.Checkuserifexist(convertuserid)
			if check != convertuserid {
				botdbconnect.InsertNewUserID(convertuserid)
				reply = "Привет, игрок ФК Реддис, пройди регистрацию, напиши свое имя:"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
				bot.Send(msg)
				for update := range updates {
					if update.Message != nil {
						username := update.Message.Text
						botdbconnect.InsertNewUserName(username, convertuserid)

					}

					reply = "Напиши Фамилию"
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
					bot.Send(msg)
					for update := range updates {
						if update.Message != nil {
							usersecondname := update.Message.Text
							botdbconnect.InsertNewUserSecondName(usersecondname, convertuserid)
							whataname := botdbconnect.SelectUserName(convertuserid)
							reply = fmt.Sprintln("Добро пожаловать в ФК Реддис", whataname, "\n", "Вот список команд:\n /balance - проверить свой баланс\n /table - посмотреть расписание тренеровок\n /play - записаться на ближайшее событие\n")
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
							bot.Send(msg)
							break
						}
						break
					}
					break
				}
				break
			} else {
				useridtwo := update.Message.From.ID
				var convertuseridtwo string
				convertuseridtwo = strconv.FormatInt(useridtwo, 10)

				whataisyourname := botdbconnect.SelectUserName(convertuseridtwo)
				reply = fmt.Sprintln("Привет,", whataisyourname, "\n", "Вот список команд: /balance - проверить свой баланс\n /table - посмотреть расписание тренеровок\n /play - записаться на ближайшее событие\n /oplata- пополнить баланс\n")
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
				bot.Send(msg)

			}
			break
		case "balance":
			userid := update.Message.From.ID
			var convertuserid string
			convertuserid = strconv.FormatInt(userid, 10)
			whataisyoubalance := botdbconnect.GetBalance(convertuserid)
			reply = fmt.Sprintln("Твой баланс: ", whataisyoubalance, "\n")
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			bot.Send(msg)

		case "oplata":
			userid := update.Message.From.ID
			var convertuserid string
			convertuserid = strconv.FormatInt(userid, 10)
			reply = "Сколько ты хочешь заплатить?"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			bot.Send(msg)
			for update := range updates {
				if update.Message != nil {
					oplata := update.Message.Text
					botdbconnect.Oplata(oplata, convertuserid)
					reply = "Спасибо!"
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
					bot.Send(msg)
				}
			}
		case "play":
			teamoneprint := botdbconnect.GetSheduleForTemaOne()
			teamtwoprint := botdbconnect.GetSheduleForTemaTwo()
			teamtreeprint := botdbconnect.GetSheduleForTemaTree()
			teamfourprin := botdbconnect.GetSheduleForTemaFour()
			teamone := fmt.Sprint("1.", teamoneprint)
			teamtwo := fmt.Sprint("2.", teamtwoprint)
			teamtree := fmt.Sprint("3.", teamtreeprint)
			teamfour := fmt.Sprint("4.", teamfourprin)

			reply = "Выбери событие:"

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			msg2 := tgbotapi.NewMessage(update.Message.Chat.ID, teamone)
			msg3 := tgbotapi.NewMessage(update.Message.Chat.ID, teamtwo)
			msg4 := tgbotapi.NewMessage(update.Message.Chat.ID, teamtree)
			msg5 := tgbotapi.NewMessage(update.Message.Chat.ID, teamfour)
			bot.Send(msg)
			bot.Send(msg2)
			bot.Send(msg3)
			bot.Send(msg4)
			bot.Send(msg5)
			for update := range updates {
				if update.Message != nil {
					selectedteam := update.Message.Text
					if selectedteam == "1...." {
						fmt.Print("one")
					} else if selectedteam == "2" {
						fmt.Print("2....")
					} else if selectedteam == "3" {
						fmt.Print("3....")
					} else if selectedteam == "4" {
						fmt.Print("4....")
					}
				}

			}
		case "table":

			whatisshedule := botdbconnect.GetSheduleTraining()
			reply = fmt.Sprintln("Ближайшее событие: ", whatisshedule, "\n")
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			bot.Send(msg)

		}
	}
}
