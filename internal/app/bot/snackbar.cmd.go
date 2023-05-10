package bot

import (
	"github.com/dndrheadof/vk-go-bot/internal/app/api"
)

func snackbarcommand(bot *Bot, mes api.MessageEvent) {
	bot.api.CallMethod("messages.sendMessageEventAnswer", map[string]interface{}{
		"event_id":   mes.EventID,
		"user_id":    mes.UserID,
		"peer_id":    mes.PeerID,
		"event_data": "{\"type\": \"show_snackbar\",\"text\": \"хоба снекбар 😎\"}",
	}, nil)
	// bot.SendMessage(SendMessageInput{
	// 	attachment: "wall-133242721_84449",
	// 	peerID:     mes.PeerID,
	// })
}
