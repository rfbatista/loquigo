package eventmanager

import (
	"loquigo/engine/pkg/core/domain"
)

type ChatServicePort interface{}

type EventRepository interface{}

type UserRepository interface {
	FindUserOrCreate(userExternalId string) (domain.User, error)
}
