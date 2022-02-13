package domain

type Stop struct {
	FlowId string
	StepId string
}

type Component interface {
	Run(m Message, u UserContext) ([]Message, *Stop, error)
}
