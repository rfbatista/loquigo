package dialog

import (
	"context"
	"loquigo/engine/database"

	"go.mongodb.org/mongo-driver/bson"
)

var actionCollection = "action"

func FindById(id string) (*Action, error) {
	db := database.GetMongoConnection()
	filter := bson.M{"_id": id}
	var result Action
	err := db.Collection(actionCollection).FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
