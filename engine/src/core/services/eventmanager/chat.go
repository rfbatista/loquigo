package eventmanager

import (
	"loquigo/engine/src/core/domain"
	"loquigo/engine/src/core/services/dialogmanager"
)

func NewChatService(d dialogmanager.DialogManagerService) ChatService {
	return ChatService{dm: d}
}

type ChatService struct {
	dm dialogmanager.DialogManagerService
}

func (c ChatService) Run(e domain.Event) ([]domain.Message, error) {
	_, err := verifyBot(e)
	if err != nil {
		return nil, err
	}
	// Maybe block user as a feature?
	// _, err = verifyUser(e)
	// if err != nil {
	// 	return nil, err
	// }

	if err != nil {
		return nil, err
	}

	messages, err := c.dm.Run(e)
	if err != nil {
		return messages, err
	}

	return messages, nil
}

func verifyBot(e domain.Event) (bool, error) {
	return true, nil
}

// func verifyUser(e domain.Event) (bool, error) {
// 	return true, nil
// }
