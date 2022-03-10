package schemas

import (
	"loquigo/engine/src/core/modules/template/pool"
)

func NewStepSchma(step pool.Step) (StepSchema, error) {
	return StepSchema{
		ID:     step.ID,
		BotId:  step.BotId,
		FlowId: step.FlowId,
		Name:   step.Name,
	}, nil
}

type StepSchema struct {
	ID     string `bson:"id"`
	BotId  string `bson:"bot_id"`
	FlowId string `bson:"flow_id"`
	Name   string `bson:"name"`
}

func (s StepSchema) ToDomain() pool.Step {
	return pool.Step{ID: s.ID, BotId: s.BotId, FlowId: s.FlowId, Name: s.Name}
}
