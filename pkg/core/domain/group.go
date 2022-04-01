package domain

type Group struct {
	ID           string `json:"id"`
	BotReference string `json:"botId"`
	BeginId      string `json:"beginId"`
	Name         string `json:"name"`
}

func NewFlow(ID string, botReference string, name string) Group {
	return Group{ID: ID, BotReference: botReference, Name: name}
}
