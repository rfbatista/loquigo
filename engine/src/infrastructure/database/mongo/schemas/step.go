package schemas

import (
	"loquigo/engine/src/core/modules/templatepool"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewStepSchma(step templatepool.Step) (StepSchema, error) {
	id, _ := primitive.ObjectIDFromHex(step.ID)
	flowId, _ := primitive.ObjectIDFromHex(step.FlowId)
	return StepSchema{
		ID:     id,
		FlowId: flowId,
		Name:   step.Name,
	}, nil
}

type StepSchema struct {
	ID     primitive.ObjectID `bson:"_id"`
	FlowId primitive.ObjectID `bson:"flow_id"`
	Name   string             `bson:"name"`
}

func (s StepSchema) ToDomain() templatepool.Step {
	return templatepool.NewStep(s.ID.Hex(), s.FlowId.String(), s.Name)
}