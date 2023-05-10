package bot

import (
	"github.com/dndrheadof/vk-go-bot/internal/app/api"
)

func startcommand(bot *Bot, mes api.Message) {
	bot.SendMessage(SendMessageInput{
		message:  "Добро пожаловать в этого замечательного бота! Сильно долго развлекать не смогу, но хоть ненадолго задержу внимание :D",
		keyboard: MainMenuKeyboard(),
		peerID:   mes.PeerID,
	})
}
