package pool

func NewStep(ID string, FlowID string, Name string) Step {
	return Step{
		ID:     ID,
		FlowId: FlowID,
		Name:   Name,
	}
}

type Step struct {
	ID         string `json:"id"`
	BotId      string `json:"bot_id"`
	FlowId     string `json:"flow_id"`
	Name       string `json:"name"`
	Components []Component
}
