package repositories

import (
	"context"
	"fmt"
	"log"
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

func NewComponentRepo(mongodb database.MongoDB) ComponentRepository {
	usersCollection := mongodb.Collection("flow_component")
	return ComponentRepository{collection: usersCollection}
}

type ComponentRepository struct {
	collection mongo.Collection
}

func (c ComponentRepository) FindByFlowAndStepId(flowId string, stepId string) ([]pool.Component, error) {
	filter := bson.M{"flow_id": flowId, "step_id": stepId}
	opts := options.Find()
	cur, err := c.collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return []pool.Component{}, err
	}
	defer cur.Close(context.TODO())
	var components = []pool.Component{}
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem schemas.ComponentSchema
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(elem)
		components = append(components, elem.ToDomain())
	}
	return components, nil
}

func (c ComponentRepository) Create(Icomponent pool.Component) (pool.Component, error) {
	schema, _ := schemas.NewComponentSchema(Icomponent)
	schema.ID = primitive.NewObjectID()
	_, err := c.collection.InsertOne(context.TODO(), schema)
	if err != nil {
		return pool.Component{}, err
	}
	return schema.ToDomain(), nil
}

func (c ComponentRepository) Update(Icomponent pool.Component) (pool.Component, error) {
	schema, _ := schemas.NewComponentSchema(Icomponent)

	opts := options.Update().SetUpsert(true)
	filter := bson.M{"_id": schema.ID}
	update := bson.D{{"$set", bson.M{
		"flow_id":  schema.FlowId,
		"step_id":  schema.StepId,
		"type":     schema.Type,
		"data":     schema.Data,
		"sequence": schema.Sequence},
	}}
	_, err := c.collection.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {

		return pool.Component{}, err
	}
	return schema.ToDomain(), nil
}

func (c ComponentRepository) Delete(Icomponent pool.Component) (pool.Component, error) {
	schema, _ := schemas.NewComponentSchema(Icomponent)
	opts := options.Delete()
	filter := bson.M{"_id": schema.ID}
	result, err := c.collection.DeleteOne(context.TODO(), filter, opts)
	if result.DeletedCount == 0 {
		fmt.Println("Error deleting component")
	}
	if err != nil {
		return pool.Component{}, err
	}
	return schema.ToDomain(), nil
}

func (c ComponentRepository) DeleteByBotID(botId string) error {
	opts := options.Delete()
	filter := bson.M{"bot_id": bson.M{"$eq": botId}}
	result, err := c.collection.DeleteMany(context.TODO(), filter, opts)
	if result.DeletedCount == 0 {
		fmt.Println("Error deleting component")
	}
	if err != nil {
		return err
	}
	return nil
}

func (c ComponentRepository) FindByFlowIdAndStepId(flowId string, stepId string) ([]runner.RunnerComponent, error) {
	filter := bson.M{"flow_id": flowId, "step_id": stepId}
	opts := options.Find()
	cur, err := c.collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return []runner.RunnerComponent{}, err
	}
	defer cur.Close(context.TODO())
	var runnerComponents = []runner.RunnerComponent{}
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem schemas.ComponentSchema
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		runnerComponents = append(runnerComponents, components.BuildRunnerComponent(elem.ToDomain()))
	}
	return runnerComponents, nil
}
