package contextmanager

import (
	"loquigo/engine/src/core/domain"
)

type UserContextRepository interface {
	SaveMemory(context domain.UserContext) error
	GetMemoryByUserId(userId string) (domain.UserContext, error)
}
