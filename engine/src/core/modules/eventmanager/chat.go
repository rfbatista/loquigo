package eventmanager

import (
	"loquigo/engine/src/core/domain"
	"loquigo/engine/src/core/modules/dialogmanager"
)

func NewChatService(s SendMessageService, d dialogmanager.RunDialogService, userRepo UserRepository) ChatService {
	return ChatService{sender: s, dm: d, userRepo: userRepo}
}

type ChatService struct {
	dm       dialogmanager.RunDialogService
	sender   SendMessageService
	userRepo UserRepository
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

	go c.userRepo.FindUserOrCreate(e.User.ExternalId)

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
