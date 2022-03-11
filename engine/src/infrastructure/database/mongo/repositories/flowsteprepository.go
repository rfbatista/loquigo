package repositories

import (
	"context"
	"loquigo/engine/src/core/modules/components"
	"loquigo/engine/src/core/modules/template/pool"
	"loquigo/engine/src/core/modules/template/runner"
	database "loquigo/engine/src/infrastructure/database/mongo"
	"loquigo/engine/src/infrastructure/database/mongo/schemas"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewStepRepository(mongodb database.MongoDB) StepRepository {
	usersCollection := mongodb.Collection("flow_step")
	return StepRepository{collection: usersCollection}
}

type StepRepository struct {
	collection mongo.Collection
}

func (s StepRepository) FindByFlowId(id string) ([]pool.Step, error) {
	filter := bson.D{
		primitive.E{Key: "flow_id", Value: id},
	}
	projection := bson.D{}
	opts := options.Find().SetProjection(projection)
	var schemas []schemas.StepSchema
	cursor, err := s.collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return []pool.Step{}, err
	}
	defer cursor.Close(context.TODO())
	if err := cursor.All(context.TODO(), &schemas); err != nil {
		return []pool.Step{}, err
	}
	var components = []pool.Step{}
	for _, schema := range schemas {
		components = append(components, schema.ToDomain())
	}
	return components, nil
}

func (s StepRepository) FindById(id string) (pool.Step, error) {
	ID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{
		primitive.E{Key: "_id", Value: ID},
	}
	projection := bson.D{}
	opts := options.FindOne().SetProjection(projection)
	var schemas schemas.StepSchema
	err := s.collection.FindOne(context.TODO(), filter, opts).Decode(&schemas)
	if err != nil {
		return pool.Step{}, err
	}
	return schemas.ToDomain(), nil
}

func (s StepRepository) Create(step pool.Step) (pool.Step, error) {
	schema, _ := schemas.NewStepSchma(step)
	_, err := s.collection.InsertOne(context.TODO(), schema)
	if err != nil {
		return pool.Step{}, err
	}
	return schema.ToDomain(), nil
}

func (s StepRepository) Update(step pool.Step) (pool.Step, error) {
	schema, _ := schemas.NewStepSchma(step)
	opts := options.Update().SetUpsert(false)
	filter := bson.D{primitive.E{Key: "_id", Value: schema.ID}}
	_, err := s.collection.UpdateOne(context.TODO(), filter, schema, opts)
	if err != nil {
		return pool.Step{}, err
	}
	return schema.ToDomain(), nil
}

func (s StepRepository) Delete(step pool.Step) (pool.Step, error) {
	schema, _ := schemas.NewStepSchma(step)
	opts := options.Delete()
	filter := bson.D{primitive.E{Key: "_id", Value: schema.ID}}
	_, err := s.collection.DeleteOne(context.TODO(), filter, opts)
	if err != nil {
		return pool.Step{}, err
	}
	return schema.ToDomain(), nil
}

func (s StepRepository) DeleteByBotID(botId string) error {
	opts := options.Delete()
	filter := bson.M{"bot_id": bson.M{"$eq": botId}}
	result, err := s.collection.DeleteMany(context.TODO(), filter, opts)
	if result.DeletedCount == 0 {

	}
	if err != nil {
		return err
	}
	return nil
}

func (c StepRepository) FindByIdAndFlowId(flowId string, stepId string) (pool.Step, error) {
	filter := bson.D{
		primitive.E{Key: "flow_id", Value: flowId},
		primitive.E{Key: "id", Value: stepId},
	}
	projection := bson.D{}
	opts := options.FindOne().SetProjection(projection)
	var schema schemas.StepSchema
	err := c.collection.FindOne(context.TODO(), filter, opts).Decode(&schema)
	if err != nil {
		return pool.Step{}, err
	}
	return schema.ToDomain(), nil
}

func (c StepRepository) FindByFlowIdAndStepId(flowId string, stepId string) (runner.RunnerStep, error) {
	filter := bson.D{
		primitive.E{Key: "flow_id", Value: flowId},
		primitive.E{Key: "id", Value: stepId},
	}
	projection := bson.D{}
	opts := options.FindOne().SetProjection(projection)
	var schema []schemas.StepSchema
	err := c.collection.FindOne(context.TODO(), filter, opts).Decode(&schema)
	if err != nil {
		return components.Step{}, err
	}
	return components.Step{}, nil
}
