package dialog

import (
	"context"
	"loquigo/engine/database"
	"loquigo/engine/message"
)

var collection = "action"

type ActionType int16

const (
	UndefinedType ActionType = iota
	TextMessage
	ImageMessage
)

func (m ActionType) String() string {
	switch m {
	case TextMessage:
		return "text_message"
	case ImageMessage:
		return "image_message"
	}
	return "unknown"
}

type Action struct {
	Id           string     `bson:"_id"`
	TransitionID string     `bson:"transition_id"`
	Type         ActionType `bson:"type"`
	Data         ActionData `bson:"data,inline"`
}

func (a *Action) Run(event Event) *message.Message {
	return &a.Data.ActionMessage
}

func (a *Action) Save() {
	db := database.GetMongoConnection()
	db.Collection(actionCollection).InsertOne(context.TODO(), a)
}
