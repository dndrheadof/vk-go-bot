package bot

import "github.com/dndrheadof/vk-go-bot/internal/app/api"

type ButtonPayload struct {
	Command string `json:"command"`
}

func MainMenuKeyboard() *api.Keyboard {
	return &api.Keyboard{
		OneTime: false,
		Buttons: [][]api.Button{
			{
				{
					Action: api.Action{
						Type: "text",
						Payload: ButtonPayload{
							Command: "text_menu",
						},
						Label: "Меню с текстом",
					},
					Color: api.PositiveButtonColor,
				},
				{
					Action: api.Action{
						Type: "text",
						Payload: ButtonPayload{
							Command: "attachment_menu",
						},
						Label: "Меню с вложением",
					},
					Color: api.NegativeButtonColor,
				},
			}, {
				{
					Action: api.Action{
						Type: "text",
						Payload: ButtonPayload{
							Command: "snackbar_menu",
						},
						Label: "Меню с snackbar",
					},
					Color: api.DefaultButtonColor,
				},
				{
					Action: api.Action{
						Type: "text",
						Payload: ButtonPayload{
							Command: "what",
						},
						Label: "А это кто...",
					},
					Color: api.PrimaryButtonColor,
				},
			},
		},
	}
}

func AttachmentMenuKeyboard() *api.Keyboard {
	return &api.Keyboard{
		OneTime: false,
		Buttons: [][]api.Button{
			{
				{
					Action: api.Action{
						Type: "text",
						Payload: ButtonPayload{
							Command: "attachment",
						},
						Label: "Покажи вложение",
					},
					Color: api.PositiveButtonColor,
				},
				{
					Action: api.Action{
						Type: "text",
						Payload: ButtonPayload{
							Command: "mainmenu",
						},
						Label: "Назад",
					},
					Color: api.NegativeButtonColor,
				},
			},
		},
	}
}

func SnackbarMenuKeyboard() *api.Keyboard {
	return &api.Keyboard{
		OneTime: false,
		Buttons: [][]api.Button{
			{
				{
					Action: api.Action{
						Type: "callback",
						Payload: ButtonPayload{
							Command: "snackbar",
						},
						Label: "Покажи snackbar",
					},
					Color: api.PositiveButtonColor,
				},
				{
					Action: api.Action{
						Type: "text",
						Payload: ButtonPayload{
							Command: "mainmenu",
						},
						Label: "Назад",
					},
					Color: api.NegativeButtonColor,
				},
			},
		},
	}
}

func TextCommandKeyboard() *api.Keyboard {
	return &api.Keyboard{
		OneTime: false,
		Buttons: [][]api.Button{
			{
				{
					Action: api.Action{
						Type: "text",
						Payload: ButtonPayload{
							Command: "text",
						},
						Label: "Покажи текст",
					},
					Color: api.PositiveButtonColor,
				},
				{
					Action: api.Action{
						Type: "text",
						Payload: ButtonPayload{
							Command: "mainmenu",
						},
						Label: "Назад",
					},
					Color: api.NegativeButtonColor,
				},
			},
		},
	}
}
