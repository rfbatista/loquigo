package schemas

import (
	"loquigo/engine/src/core/modules/template/pool"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewComponentSchema(component pool.Component) (ComponentSchema, error) {
	ID, IDError := primitive.ObjectIDFromHex(component.ID)
	if IDError != nil {
		ID = primitive.NewObjectID()
	}
	return ComponentSchema{
		ID:       ID,
		BotId:    component.BotId,
		FlowId:   component.FlowId,
		StepId:   component.StepId,
		Type:     component.Type,
		Data:     component.Data,
		Sequence: component.Sequence,
	}, nil
}

type ComponentSchema struct {
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	BotId    string             `bson:"bot_id"`
	FlowId   string             `bson:"flow_id"`
	StepId   string             `bson:"step_id"`
	Type     string             `bson:"type"`
	Data     pool.ComponentData `bson:"data"`
	Sequence int                `bson:"sequence"`
}

func (c ComponentSchema) ToDomain() pool.Component {
	return pool.Component{
		ID:       c.ID.Hex(),
		FlowId:   c.FlowId,
		StepId:   c.StepId,
		Type:     c.Type,
		Data:     c.Data,
		Sequence: c.Sequence,
	}
}
