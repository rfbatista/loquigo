package pool

type Stop struct {
	FlowId string
	StepId string
}
type GoTo struct {
	FlowId string
	StepId string
}

type Component struct {
	ID       string        `json:"id"`
	FlowId   string        `json:"flowId"`
	StepId   string        `json:"stepId"`
	Type     string        `json:"type" `
	Data     ComponentData `json:"data"`
	Sequence int           `json:"sequence"`
}

type ComponentData struct {
	Text   string
	FlowId string
	StepId string
}
