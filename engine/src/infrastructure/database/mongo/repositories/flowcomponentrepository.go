package repositories

import (
	"context"
	"fmt"
	"log"
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
	flowIDHex, _ := primitive.ObjectIDFromHex(flowId)
	stepIDHex, _ := primitive.ObjectIDFromHex(stepId)
	filter := bson.D{
		primitive.E{Key: "flow_id", Value: flowIDHex},
		primitive.E{Key: "step_id", Value: stepIDHex},
	}
	projection := bson.D{{"sequence", 1}}
	opts := options.Find().SetProjection(projection)
	cur, err := c.collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return []templatepool.Component{}, err
	}
	defer cur.Close(context.TODO())
	var components = []templatepool.Component{}
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

	opts := options.Update().SetUpsert(true)
	filter := bson.M{"_id": schema.ID}
	update := bson.D{{"$set", bson.M{
		"flow_id":  schema.Flow_id,
		"step_id":  schema.Step_id,
		"data":     schema.Data,
		"sequence": schema.Sequence},
	}}
	_, err := c.collection.UpdateOne(context.Background(), filter, update, opts)
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
