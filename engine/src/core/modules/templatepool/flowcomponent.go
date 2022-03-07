package templatepool

import "loquigo/engine/src/core/domain"

type Stop struct {
	FlowId string
	StepId string
}
type GoTo struct {
	FlowId string
	StepId string
}

type IComponent interface {
	Run(userMess domain.Message, u domain.UserContext, botMessages []domain.Message) ([]domain.Message, *Stop, *GoTo)
}
