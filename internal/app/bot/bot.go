package bot

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dndrheadof/vk-go-bot/config"
	"github.com/dndrheadof/vk-go-bot/internal/app/api"
)

type Bot struct {
	api     *api.Api
	config  *config.Config
	groupID int
}

type SendMessageInput struct {
	peerID     int
	message    string
	attachment string
	keyboard   *api.Keyboard
}

func NewBot(config *config.Config) error {
	vkApi := api.NewApi(config.Token)

	bot := &Bot{
		api:     vkApi,
		groupID: config.GroupID,
		config:  config,
	}

	lp, err := api.NewLongpoll(vkApi, config.GroupID)
	if err != nil {
		log.Printf("ОШИБКА: %v\n", err)

		return err
	}

	go lp.ListenForEvents()

	for {
		select {
		case mes := <-lp.NewMessage:
			go bot.HandleMessage(mes)
		case mes := <-lp.NewMessageEvent:
			go bot.HandleMessageEvent(mes)
		}
	}
}

func (bot *Bot) HandleMessage(message api.Message) {
	var payload ButtonPayload
	err := json.Unmarshal([]byte(message.Payload), &payload)
	if err != nil {
		log.Printf("ОШИБКА: %v\n", err)

		Commands["mainmenu"].Exec(bot, message)
	}

	if payload.Command != "" {
		cmd, exists := Commands[payload.Command]
		if !exists {

			bot.SendMessage(SendMessageInput{
				peerID:  message.PeerID,
				message: "Команда не найдена :c",
			})

			return
		}
		cmd.Exec(bot, message)
	}
	fmt.Printf("%+v\n", message)
}

func (bot *Bot) HandleMessageEvent(event api.MessageEvent) {
	if event.Payload.Command != "" {
		cmd, exists := Commands[event.Payload.Command]
		if !exists {
			return
		}
		cmd.ExecEvent(bot, event)
	}
	fmt.Printf("%+v\n", event)
}

func (bot *Bot) SendMessage(input SendMessageInput) error {
	params := map[string]interface{}{
		"peer_id":    input.peerID,
		"random_id":  0,
		"message":    input.message,
		"attachment": input.attachment,
	}

	if input.keyboard != nil {
		var b []byte
		b, err := json.Marshal(input.keyboard)
		if err != nil {
			log.Printf("ОШИБКА: %v\n", err)

			return err
		}
		params["keyboard"] = string(b)
	}
	return bot.api.CallMethod("messages.send", params, nil)
}
