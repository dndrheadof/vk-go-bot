package main

import (
	"log"
	"os"

	"github.com/dndrheadof/vk-go-bot/config"
	"github.com/dndrheadof/vk-go-bot/internal/app/bot"
)

func main() {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	log.SetOutput(file)

	bot.NewBot(&cfg)
}
