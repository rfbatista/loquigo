package dialog

import (
	"context"
	"loquigo/engine/database"

	"go.mongodb.org/mongo-driver/bson"
)

var transitionCollection = "transition"

func FindTransitionById(id string) (*Transition, error) {
	db := database.GetMongoConnection()
	filter := bson.M{"_id": id}
	var result Transition
	err := db.Collection(transitionCollection).FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
