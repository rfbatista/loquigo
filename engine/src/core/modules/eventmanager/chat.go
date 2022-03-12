package eventmanager

import (
	"loquigo/engine/src/core/domain"
	"loquigo/engine/src/core/modules/dialogmanager"
)

func NewChatService(d dialogmanager.RunDialogService, userRepo UserRepository) ChatService {
	return ChatService{dm: d, userRepo: userRepo}
}

type ChatService struct {
	dm dialogmanager.RunDialogService
	// sender   SendMessageService
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

	c.userRepo.FindUserOrCreate(e.User.ExternalId)

	messages, err := c.dm.Run(e)
	if err != nil {
		return messages, err
	}
	// c.sender.Send(messages)

	return messages, nil
}

func verifyBot(e domain.Event) (bool, error) {
	return true, nil
}

func verifyUser(e domain.Event) (bool, error) {
	return true, nil
}
