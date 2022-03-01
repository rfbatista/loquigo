package templatepool

type FlowHash = map[string]StepHash
type StepHash = map[string]IStep

type IFlow interface {
	Name() string
	Steps(string) Step
}

func NewFlow(ID string, BotId string, name string) Flow {
	return Flow{ID: ID, BotId: BotId, Name: name}
}

type Flow struct {
	ID    string `json:"id"`
	BotId string `json:"bot_id"`
	Name  string `json:"name"`
}
