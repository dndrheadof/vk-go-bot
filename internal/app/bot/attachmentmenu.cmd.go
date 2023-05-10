package bot

import (
	"github.com/dndrheadof/vk-go-bot/internal/app/api"
)

func attachmentmenucommand(bot *Bot, mes api.Message) {
	bot.SendMessage(SendMessageInput{
		message:  "Меню с кнопкой, которая отправляет вложение",
		keyboard: AttachmentMenuKeyboard(),
		peerID:   mes.PeerID,
	})
}
