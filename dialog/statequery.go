package dialog

import (
	"context"
	"loquigo/engine/database"

	"go.mongodb.org/mongo-driver/bson"
)

var stateCollection = "state"

func FindStateById(id string) (*State, error) {
	db := database.GetMongoConnection()
	filter := bson.M{"_id": id}
	var result State
	err := db.Collection(stateCollection).FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
