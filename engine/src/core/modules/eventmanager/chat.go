package eventmanager

import (
	"loquigo/engine/src/core/domain"
	"loquigo/engine/src/core/modules/dialogmanager"
)

type ChatService struct {
	dm     dialogmanager.RunDialogService
	sender SendMessageService
}

func NewChatService(s SendMessageService, d dialogmanager.RunDialogService) ChatService {
	return ChatService{sender: s, dm: d}
}

func (c ChatService) Run(e domain.Event) ([]domain.Message, error) {
	_, err := verifyBot(e)
	if err != nil {
		return nil, err
	}
	_, err = verifyUser(e)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	messages, _ := c.dm.Run(e)
	c.sender.Send(messages)

	return messages, nil
}

func verifyBot(e domain.Event) (bool, error) {
	return true, nil
}

func verifyUser(e domain.Event) (bool, error) {
	return true, nil
}
