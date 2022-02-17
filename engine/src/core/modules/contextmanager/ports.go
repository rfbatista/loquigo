package contextmanager

import (
	"loquigo/engine/src/core/domain"
)

type UserContextRepository interface {
	SaveMemory(context domain.UserContext) error
	FindByUserId(userId string) (domain.UserContext, error)
}
