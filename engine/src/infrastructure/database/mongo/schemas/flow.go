package schemas

import (
	"loquigo/engine/src/core/modules/template/pool"
)

func NewFlowSchema(flow pool.Flow) (FlowSchema, error) {
	return FlowSchema{
		ID:    flow.ID,
		BotId: flow.BotId,
		Name:  flow.Name,
	}, nil
}

type FlowSchema struct {
	ID      string `bson:"id"`
	BotId   string `bson:"bot_id"`
	BeginId string `bson:"begin_id"`
	Name    string `bson:"name"`
}

func (f FlowSchema) ToDomain() pool.Flow {
	return pool.Flow{ID: f.ID, BotId: f.BotId, BeginId: f.BeginId, Name: f.Name}
}
