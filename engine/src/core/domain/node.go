package domain

func NewStep(ID string, FlowID string, Name string) Node {
	return Node{
		ID:     ID,
		NodeId: FlowID,
		Name:   Name,
	}
}

type Node struct {
	ID           string
	BotReference string `json:"bot_id"`
	NodeId       string `json:"group_id"`
	Name         string `json:"name"`
	Components   []Component
}
