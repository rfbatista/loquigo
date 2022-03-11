package schemas

import (
	"loquigo/engine/src/core/modules/template/pool"
)

func NewComponentSchema(component pool.Component) (ComponentSchema, error) {
	return ComponentSchema{
		ID:       component.ID,
		BotId:    component.BotId,
		FlowId:   component.FlowId,
		StepId:   component.StepId,
		Type:     component.Type,
		Data:     component.Data,
		Sequence: component.Sequence,
	}, nil
}

type ComponentSchema struct {
	ID       string             `bson:"id" json:"id,omitempty"`
	BotId    string             `bson:"bot_id"`
	FlowId   string             `bson:"flow_id"`
	StepId   string             `bson:"step_id"`
	Type     string             `bson:"type"`
	Data     pool.ComponentData `bson:"data"`
	Sequence int                `bson:"sequence"`
}

func (c ComponentSchema) ToDomain() pool.Component {
	return pool.Component{
		ID:       c.ID,
		FlowId:   c.FlowId,
		StepId:   c.StepId,
		Type:     c.Type,
		Data:     c.Data,
		Sequence: c.Sequence,
	}
}
