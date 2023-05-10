package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

type LongPollSessionData struct {
	Key    string `json:"key"`
	Server string `json:"server"`
	TS     int    `json:"ts"`
}

type Longpoll struct {
	GroupID         int
	Api             *Api
	Params          url.Values
	Server          string
	NewMessage      chan Message
	NewMessageEvent chan MessageEvent
}

const (
	LPHistoryOld = 1
	LPKeyExpired = 2
	LPInfoLost   = 3
)

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
		GroupID:         groupID,
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

	if res.FailedCode > 0 {
		log.Printf("Пришла ошибка от LP: %d\n", res.FailedCode)
		switch res.FailedCode {
		case LPKeyExpired, LPInfoLost:
			upd := &LPSetupResponse{}
			log.Println("LP ключ устарел, получаем новый")

			err := lp.Api.CallMethod("groups.getLongPollServer", map[string]interface{}{
				"group_id": lp.GroupID,
			}, &upd)
			if err != nil {
				return nil, fmt.Errorf("произошла ошибка при рефреше ключа: %+v", err)
			}
			lp.Params.Set("key", upd.Response.Key)
		case LPHistoryOld:
			log.Println("История устарела")
			lp.Params.Set("ts", fmt.Sprint(res.Ts))
		}
	}

	log.Printf("[LPResponse] %+v", res)

	lp.Params.Set("ts", fmt.Sprint(res.Ts))

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
