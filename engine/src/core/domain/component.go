package domain

type Stop struct {
	FlowId string
	StepId string
}
type GoTo struct {
	FlowId string
	StepId string
}

type Component interface {
	Run(userMess Message, u UserContext, botMessages []Message) ([]Message, *Stop, *GoTo)
}
