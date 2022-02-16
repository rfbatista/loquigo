package infrastructure

import (
	"bytes"
	"encoding/json"
	ports "loquigo/engine/src/core"
	"net/http"
	"time"
)

func NewHttpClient() ports.HttpClient {
	return HttpClient{}
}

type HttpClient struct{}

func (h HttpClient) Client(baseurl string) ports.WebserviceClient {
	return WebserviceClient{baseUrl: baseurl, client: &http.Client{
		Timeout: time.Second * 10,
	},
	}
}

type WebserviceClient struct {
	baseUrl string
	client  *http.Client
}

func (w WebserviceClient) Post(data interface{}) (interface{}, error) {
	jsonValue, _ := json.Marshal(data)
	resp, err := w.client.Post(w.baseUrl, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w WebserviceClient) Get(data interface{}) (interface{}, error) {
	return nil, nil
}
