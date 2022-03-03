package repositories

import (
	"context"
	"loquigo/engine/src/core/modules/templatepool"
	database "loquigo/engine/src/infrastructure/database/mongo"
	"loquigo/engine/src/infrastructure/database/mongo/schemas"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewComponentRepo(mongodb database.MongoDB) ComponentRepository {
	usersCollection := mongodb.Collection("flow_component")
	return ComponentRepository{collection: usersCollection}
}

type ComponentRepository struct {
	collection mongo.Collection
}

func (c ComponentRepository) FindByFlowAndStepId(flowId string, stepId string) ([]templatepool.Component, error) {
	filter := bson.D{
		primitive.E{Key: "flow_id", Value: flowId},
		primitive.E{Key: "step_id", Value: stepId},
	}
	projection := bson.D{
		primitive.E{Key: "sequence", Value: 1},
	}
	opts := options.Find().SetProjection(projection)
	var schemas []schemas.ComponentSchema
	cursor, err := c.collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return []templatepool.Component{}, err
	}
	defer cursor.Close(context.TODO())
	if err := cursor.All(context.TODO(), &schemas); err != nil {
		return []templatepool.Component{}, err
	}
	var components = []templatepool.Component{}
	for _, schema := range schemas {
		components = append(components, schema.ToDomain())
	}
	return components, nil
}

func (c ComponentRepository) Create(component templatepool.Component) (templatepool.Component, error) {
	schema, _ := schemas.NewComponentSchema(component)
	schema.ID = primitive.NewObjectID()
	_, err := c.collection.InsertOne(context.TODO(), schema)
	if err != nil {
		return templatepool.Component{}, err
	}
	return schema.ToDomain(), nil
}

func (c ComponentRepository) Update(component templatepool.Component) (templatepool.Component, error) {
	schema, _ := schemas.NewComponentSchema(component)
	opts := options.Update().SetUpsert(false)
	filter := bson.D{primitive.E{Key: "_id", Value: schema.ID}}
	_, err := c.collection.UpdateOne(context.TODO(), filter, schema, opts)
	if err != nil {
		return templatepool.Component{}, err
	}
	return schema.ToDomain(), nil
}
func (c ComponentRepository) Delete(component templatepool.Component) (templatepool.Component, error) {
	schema, _ := schemas.NewComponentSchema(component)
	opts := options.Delete()
	filter := bson.D{primitive.E{Key: "_id", Value: schema.ID}}
	_, err := c.collection.DeleteOne(context.TODO(), filter, opts)
	if err != nil {
		return templatepool.Component{}, err
	}
	return schema.ToDomain(), nil
}