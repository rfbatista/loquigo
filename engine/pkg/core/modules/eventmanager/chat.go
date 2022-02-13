package eventmanager

import (
	"loquigo/engine/pkg/core/domain"
	"loquigo/engine/pkg/core/modules/dialogmanager"
)

type ChatService struct {
	dm     dialogmanager.RunDialogService
	sender SendMessageService
}

func NewChatService(s SendMessageService, d dialogmanager.RunDialogService) ChatService {
	return ChatService{sender: s, dm: d}
}

func (c ChatService) Run(e domain.Event) (bool, error) {
	_, err := verifyBot(e)
	if err != nil {
		return false, err
	}
	_, err = verifyUser(e)
	if err != nil {
		return false, err
	}

	if err != nil {
		return false, err
	}

	c.dm.Run()

	var messages []domain.Message
	c.sender.Send(messages)

	return true, nil
}

func verifyBot(e domain.Event) (bool, error) {
	return true, nil
}

func verifyUser(e domain.Event) (bool, error) {
	return true, nil
}
