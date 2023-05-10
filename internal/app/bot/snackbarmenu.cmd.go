package bot

import (
	"github.com/dndrheadof/vk-go-bot/internal/app/api"
)

func snackbarmenucommand(bot *Bot, mes api.Message) {
	bot.SendMessage(SendMessageInput{
		message:  "Меню с кнопкой, которая показывает snackbar",
		keyboard: SnackbarMenuKeyboard(),
		peerID:   mes.PeerID,
	})
}
