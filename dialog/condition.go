package dialog

import (
	"context"
	"loquigo/engine/database"
)

type Condition struct {
	Id           string `bson:"_id"`
	TransitionID string `bson:"transition_id"`
}

func (c *Condition) IsValid(event Event) bool {
	return true
}

func (u *Condition) Save() {
	db := database.GetMongoConnection()
	db.Collection(conditionCollection).InsertOne(context.TODO(), u)
}
