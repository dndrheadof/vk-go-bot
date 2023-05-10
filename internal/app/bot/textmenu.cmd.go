package bot

import (
	"github.com/dndrheadof/vk-go-bot/internal/app/api"
)

func textmenucommand(bot *Bot, mes api.Message) {
	bot.SendMessage(SendMessageInput{
		message:  "Супер полезный текст, рассказывающий о чудесной кнопке снизу",
		keyboard: TextCommandKeyboard(),
		peerID:   mes.PeerID,
	})
}
