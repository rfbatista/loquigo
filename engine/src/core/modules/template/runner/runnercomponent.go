package runner

import (
	"loquigo/engine/src/core/domain"
	"loquigo/engine/src/core/modules/template/pool"
)

type IComponent interface {
	Run(userMess domain.Message, u domain.UserContext, botMessages []domain.Message) ([]domain.Message, *pool.Stop, *pool.GoTo)
}
