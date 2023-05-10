package bot

import (
	"github.com/dndrheadof/vk-go-bot/internal/app/api"
)

type BotCommand struct {
	Exec      func(bot *Bot, mes api.Message)
	ExecEvent func(bot *Bot, mes api.MessageEvent)
}

var Commands = map[string]BotCommand{
	"start": {
		Exec: startcommand,
	},
	"mainmenu": {
		Exec: mainmenucommand,
	},
	"text_menu": {
		Exec: textmenucommand,
	},
	"attachment_menu": {
		Exec: attachmentmenucommand,
	},
	"snackbar_menu": {
		Exec: snackbarmenucommand,
	},
	"text": {
		Exec: textcommand,
	},
	"attachment": {
		Exec: attachmentcommand,
	},
	"snackbar": {
		ExecEvent: snackbarcommand,
	},
}
