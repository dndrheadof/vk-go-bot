package api

import (
	"encoding/json"
	"log"
	"net/url"
)

type LongPollSessionData struct {
	Key    string `json:"key"`
	Server string `json:"server"`
	TS     int    `json:"ts"`
}

type Longpoll struct {
	Api             *Api
	Params          url.Values
	Server          string
	NewMessage      chan Message
	NewMessageEvent chan MessageEvent
}

func NewLongpoll(api *Api, groupID int) (*Longpoll, error) {
	res := LPSetupResponse{}
	err := api.CallMethod("groups.getLongPollServer", map[string]interface{}{
		"group_id": groupID,
	}, &res)

	if err != nil {
		log.Printf("ОШИБКА: %v\n", err)

		return nil, err
	}
	urlParams := url.Values{}
	urlParams.Add("wait", "15")
	urlParams.Add("act", "a_check")
	urlParams.Add("ts", res.Response.Ts)
	urlParams.Add("key", res.Response.Key)

	return &Longpoll{
		Api:             api,
		Params:          urlParams,
		Server:          res.Response.Server,
		NewMessage:      make(chan Message),
		NewMessageEvent: make(chan MessageEvent),
	}, nil
}

func (lp *Longpoll) Fetch() (*LongpollResponse, error) {
	res := &LongpollResponse{}
	err := lp.Api.SendRequest(lp.Server, []byte(lp.Params.Encode()), &res)

	if err != nil {
		log.Printf("ОШИБКА: %v\n", err)

		return nil, err
	}

	if res.Ts == "" {
		return nil, nil
	}

	log.Printf("[LPResponse] %+v", res)

	lp.Params.Set("ts", res.Ts)

	return res, nil
}

func (lp *Longpoll) ListenForEvents() {
	for {
		event, err := lp.Fetch()
		if err != nil {
			log.Printf("ОШИБКА: %v\n", err)

			log.Fatal(err)
		}
		if event == nil {
			continue
		}

		for _, update := range event.Updates {
			if update.Type == "message_new" {
				mes := Message{}

				jsonStr, err := json.Marshal(update.Object["message"])
				if err != nil {
					log.Printf("ОШИБКА: %v\n", err)

					log.Fatal(err)
				}

				if err := json.Unmarshal(jsonStr, &mes); err != nil {
					log.Printf("ОШИБКА: %v\n", err)

					log.Fatal(err)
				}

				lp.NewMessage <- mes
			}

			if update.Type == "message_event" {
				mes := MessageEvent{}
				jsonStr, err := json.Marshal(update.Object)
				if err != nil {
					log.Printf("ОШИБКА: %v\n", err)

					log.Fatal(err)
				}
				if err := json.Unmarshal(jsonStr, &mes); err != nil {
					log.Printf("ОШИБКА: %v\n", err)

					log.Fatal(err)
				}

				lp.NewMessageEvent <- mes
			}
		}
	}
}
