package eventmanager

import (
	"loquigo/engine/src/core/domain"
)

type ChatServicePort interface{}

type EventRepository interface{}

type UserRepository interface {
	FindUserOrCreate(userExternalId string) (domain.User, error)
}
