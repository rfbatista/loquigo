package domain

type Stop struct {
	GroupId string
	NodeId  string
}
type GoTo struct {
	GroupId string
	NodeId  string
}

type Component struct {
	ID           string        `json:"id"`
	BotReference string        `json:"botReference"`
	GroupId      string        `json:"groupId"`
	NodeId       string        `json:"nodeId"`
	Type         string        `json:"type" `
	Data         ComponentData `json:"data"`
	Sequence     int           `json:"sequence"`
}

type ComponentData struct {
	Text    string
	GroupId string
	NodeId  string
}
