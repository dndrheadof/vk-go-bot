package bot

import (
	"github.com/dndrheadof/vk-go-bot/internal/app/api"
)

func mainmenucommand(bot *Bot, mes api.Message) {
	bot.SendMessage(SendMessageInput{
		message:  "Добро пожаловать в главное меню",
		keyboard: MainMenuKeyboard(),
		peerID:   mes.PeerID,
	})
}
