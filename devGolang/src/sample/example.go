package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("1057300654:AAFEuLNx-9wgeLq2XyorkmuKIPmKjVGV57s")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		switch update.Message.Text {
		case "/help":
			update.Message.Text = "알고싶나요?"
		default:
			update.Message.Text = "reply " + update.Message.Text
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}

	// msg := tgbotapi.NewMessage(960527111, "A test message from the test library in telegram-bot-api")
	// msg.ReplyToMessageID = 35
	// _, err = bot.Send(msg)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
