package dialogmanager

import "loquigo/engine/src/core/domain"

type DialogRepository interface{}
type UserContextRepository interface {
	SaveMemory(ctx domain.UserContext) error
	FindByUserId(userId string) (domain.UserContext, error)
	Create(userContext domain.UserContext)
}
