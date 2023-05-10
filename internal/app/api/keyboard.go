package api

const (
	DefaultButtonColor  ButtonColor = "default"
	PrimaryButtonColor  ButtonColor = "primary"
	NegativeButtonColor ButtonColor = "negative"
	PositiveButtonColor ButtonColor = "positive"
)

type (
	ButtonColor string

	Keyboard struct {
		Buttons [][]Button `json:"buttons"`
		OneTime bool       `json:"one_time"`
		// Inline  bool            `json:"inline"`
	}

	Button struct {
		Action Action      `json:"action"`
		Color  ButtonColor `json:"color"`
	}

	Action struct {
		Type    string `json:"type"`
		Payload struct {
			Command string `json:"command"`
		} `json:"payload"`
		Label string `json:"label"`
	}
)
