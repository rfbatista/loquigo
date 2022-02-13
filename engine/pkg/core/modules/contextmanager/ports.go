package contextmanager

import (
	"loquigo/engine/pkg/core/domain"
)

type UserContextRepository interface {
	SaveMemory(context domain.UserContext) error
	GetMemoryByUserId(userId string) (domain.UserContext, error)
}
