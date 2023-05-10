package bot

import (
	"github.com/dndrheadof/vk-go-bot/internal/app/api"
)

func textcommand(bot *Bot, mes api.Message) {
	bot.SendMessage(SendMessageInput{
		message: "Супер крутой текст",
		peerID:  mes.PeerID,
	})
}
