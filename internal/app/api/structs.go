package api

type (
	LPSetupResponse struct {
		Response struct {
			Key    string `json:"key"`
			Server string `json:"server"`
			Ts     string `json:"ts"`
		} `json:"response"`
	}

	Message struct {
		Date   int `json:"date"`
		FromID int `json:"from_id"`
		ID     int `json:"id"`
		Out    int `json:"out"`
		PeerID int `json:"peer_id"`
		// Это очень печально...
		Payload string `json:"payload"`
		// Payload               struct {
		// Command string `json:"command"`
		// } `json:"payload"`
		Text                  string `json:"text"`
		ConversationMessageID int    `json:"conversation_message_id"`
		RandomID              int    `json:"random_id"`
	}

	MessageEvent struct {
		ConversationMessageID int    `json:"conversation_message_id"`
		UserID                int    `json:"user_id"`
		PeerID                int    `json:"peer_id"`
		EventID               string `json:"event_id"`
		Payload               struct {
			Command string `json:"command"`
		} `json:"payload"`
	}

	LongpollMessage struct {
		Type    string                 `json:"type"`
		Object  map[string]interface{} `json:"object"`
		GroupID int                    `json:"group_id"`
	}

	LongpollResponse struct {
		Ts      string            `json:"ts"`
		Updates []LongpollMessage `json:"updates"`
	}
)
