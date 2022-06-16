package dialog

import (
	"context"
	"loquigo/engine/database"

	"go.mongodb.org/mongo-driver/bson"
)

type Condition struct {
	Id string
}

func (c Condition) IsValid(event Event) bool {
	return true
}

type ConditionDAO struct {
	Id string
}

func (u *ConditionDAO) Save() {
	db := database.GetMongoConnection()
	db.Collection(collection).InsertOne(context.TODO(), u)
}

func conditionToModel(a ConditionDAO) *Condition {
	return &Condition{}
}

func FindConditionById(id string) (*Condition, error) {
	db := database.GetMongoConnection()
	filter := bson.M{"_id": id}
	var result ConditionDAO
	err := db.Collection(collection).FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return conditionToModel(result), nil
}
