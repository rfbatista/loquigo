package schemas

import (
	"loquigo/engine/src/core/modules/nodes"
)

func NewGroupSchema(group nodes.Group) (GroupSchema, error) {
	return GroupSchema{
		ID:           group.ID,
		BotReference: group.BotReference,
		BeginId:      group.BeginId,
		Name:         group.Name,
	}, nil
}

type GroupSchema struct {
	ID           string `bson:"id"`
	BotReference string `bson:"bot_reference"`
	BeginId      string `bson:"begin_id"`
	Name         string `bson:"name"`
}

func (f GroupSchema) ToDomain() nodes.Group {
	return nodes.Group{ID: f.ID, BotReference: f.BotReference, BeginId: f.BeginId, Name: f.Name}
}
