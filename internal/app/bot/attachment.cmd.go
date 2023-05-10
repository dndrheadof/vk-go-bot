package bot

import (
	"github.com/dndrheadof/vk-go-bot/internal/app/api"
)

func attachmentcommand(bot *Bot, mes api.Message) {
	bot.SendMessage(SendMessageInput{
		attachment: "photo-187263168_457711405",
		peerID:     mes.PeerID,
	})
}
