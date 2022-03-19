package domain

import "fmt"

type Bot struct {
	ID      string `json:"id"`
	BeginId string `json:"beginId"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

func (b Bot) Reference() string {
	return fmt.Sprintf("%s:%s", b.ID, b.Version)
}
