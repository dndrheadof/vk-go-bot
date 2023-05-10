package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"time"

	"github.com/valyala/fasthttp"
)

const (
	APIEndpoint = "https://api.vk.com/method"
	APIVersion  = "5.131"
)

type Api struct {
	Token   string
	Url     string
	Version string
	LPData  LongPollSessionData
	Client  *fasthttp.Client
}

func NewApi(token string) *Api {
	return &Api{
		Token:   token,
		Url:     APIEndpoint,
		Version: APIVersion,
		Client: &fasthttp.Client{
			ReadTimeout:              time.Second * 5,
			WriteTimeout:             time.Second * 5,
			MaxIdleConnDuration:      time.Minute,
			NoDefaultUserAgentHeader: true,
		},
	}
}

func (api *Api) CallMethod(method string, params map[string]interface{}, response interface{}) error {
	params["access_token"] = api.Token
	params["v"] = api.Version

	log.Printf("Вызываю метод %s\n", method)

	values := url.Values{}
	for key, value := range params {
		var str string
		switch t := value.(type) {
		case string:
			str = value.(string)
		case int:
			str = strconv.Itoa(value.(int))
		default:
			bytes, err := json.Marshal(value)
			if err != nil {
				fmt.Printf("неизвестный тип: %v", t)
				log.Printf("ОШИБКА: %v\n", err)

				return err
			}
			str = string(bytes)
		}
		values.Add(key, str)
	}

	encodedParams := values.Encode()
	req := []byte(encodedParams)
	url := fmt.Sprintf("%s/%s?%s", api.Url, method, encodedParams)
	return api.SendRequest(url, req, response)
}

func (api *Api) SendRequest(url string, params []byte, response interface{}) error {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)
	req.SetBody(params)
	req.Header.SetMethod("POST")
	req.SetRequestURI(url)

	api.Client.Do(req, res)

	body := res.Body()
	if len(body) == 0 {
		return nil
	}

	return json.Unmarshal(body, response)
}
