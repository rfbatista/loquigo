package domain

import "fmt"

type Bot struct {
	ID             string `json:"id"`
	BeginId        string `json:"beginId"`
	Name           string `json:"name"`
	CurrentVersion string `json:"version"`
}

func (b Bot) Reference() string {
	return fmt.Sprintf("%s:%s", b.ID, b.CurrentVersion)
}

type BotVersion struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	BotId string `json:"bot_id"`
	Yaml  string `json:"yaml"`
}
