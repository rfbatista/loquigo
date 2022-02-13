package domain

type HoldComponent struct {
	FlowId string
	StepId string
}

func (h HoldComponent) Run(m Message, u UserContext) ([]Message, *Stop, error) {
	return []Message{}, &Stop{StepId: h.StepId, FlowId: h.FlowId}, nil
}
