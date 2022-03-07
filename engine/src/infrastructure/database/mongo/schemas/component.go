package schemas

import (
	"loquigo/engine/src/core/modules/templatepool"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewComponentSchema(component templatepool.Component) (ComponentSchema, error) {
	flow_id, _ := primitive.ObjectIDFromHex(component.FlowId)
	step_id, _ := primitive.ObjectIDFromHex(component.StepId)
	ID, IDError := primitive.ObjectIDFromHex(component.ID)
	if IDError != nil {
		ID = primitive.NewObjectID()
	}
	return ComponentSchema{
		ID:       ID,
		Flow_id:  flow_id,
		Step_id:  step_id,
		Type:     component.Type,
		Data:     component.Data,
		Sequence: component.Sequence,
	}, nil
}

type ComponentSchema struct {
	ID       primitive.ObjectID         `bson:"_id" json:"id,omitempty"`
	Flow_id  primitive.ObjectID         `bson:"flow_id"`
	Step_id  primitive.ObjectID         `bson:"step_id"`
	Type     string                     `bson:"type"`
	Data     templatepool.ComponentData `bson:"data"`
	Sequence int                        `bson:"sequence"`
}

func (c ComponentSchema) ToDomain() templatepool.Component {
	return templatepool.Component{
		ID:       c.ID.Hex(),
		FlowId:   c.Flow_id.Hex(),
		StepId:   c.Step_id.Hex(),
		Type:     c.Type,
		Data:     c.Data,
		Sequence: c.Sequence,
	}
}
