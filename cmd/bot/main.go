package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/vskltsv/telegram-bot-pocket-golang/pkg/telegram"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
)

func main(){

	bot, err := tgbotapi.NewBotAPI("2039723901:AAFi1XYj0ju6UXc17eoPA-1Y5ZXEWt0S6OI")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	pocketClient, err := pocket.NewClient("99299-de314e03fe8469631c60b853")
	if err != nil{
		log.Fatal(err)
	}

telegramBot := telegram.NewBot(bot, pocketClient, "http://localhost/")
 if err := telegramBot.Start(); err != nil {
	 log.Fatal(err)
 }


}
