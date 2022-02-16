package eventmanager

import (
	port "loquigo/engine/src/core"
	"loquigo/engine/src/core/domain"
)

type SendMessageService struct {
	client port.WebserviceClient
}

func NewSendMessageService(client port.HttpClient) SendMessageService {
	c := SendMessageService{}
	c.SetClient(client.Client("http:localhost"))
	return c
}

func (s *SendMessageService) SetClient(c port.WebserviceClient) {
	s.client = c
}

func (s *SendMessageService) Send(messages []domain.Message) {
	s.client.Post(messages)
}
