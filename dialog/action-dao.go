package dialog

import (
	"context"
	"loquigo/engine/database"

	"go.mongodb.org/mongo-driver/bson"
)

type ActionDAO struct {
	Id   string
	Type string
}

func (u *ActionDAO) Save() {
	db := database.GetMongoConnection()
	db.Collection(collection).InsertOne(context.TODO(), u)
}

func actionToModel(a ActionDAO) *Action {
	return &Action{}
}

func FindById(id string) (*Action, error) {
	db := database.GetMongoConnection()
	filter := bson.M{"_id": id}
	var result ActionDAO
	err := db.Collection(collection).FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return actionToModel(result), nil
}
