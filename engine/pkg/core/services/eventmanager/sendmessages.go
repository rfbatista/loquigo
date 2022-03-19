package eventmanager

import (
	port "loquigo/engine/pkg/core"
	"loquigo/engine/pkg/core/domain"
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
