package pool

type Flow struct {
	ID      string `json:"id"`
	BotId   string `json:"botId"`
	BeginId string `json:"beginId"`
	Name    string `json:"name"`
}

func NewFlow(ID string, BotId string, name string) Flow {
	return Flow{ID: ID, BotId: BotId, Name: name}
}
