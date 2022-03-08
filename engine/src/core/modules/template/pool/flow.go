package pool

type Flow struct {
	ID    string `json:"id"`
	BotId string `json:"bot_id"`
	Name  string `json:"name"`
}

func NewFlow(ID string, BotId string, name string) Flow {
	return Flow{ID: ID, BotId: BotId, Name: name}
}
