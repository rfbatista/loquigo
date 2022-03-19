package runner

import (
	"loquigo/engine/src/core/domain"
)

type UserStateRepo interface {
	FindByUserId(userId string) (domain.UserState, error)
	Update(userState domain.UserState) error
	Create(userId string) (domain.UserState, error)
}
