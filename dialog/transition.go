package dialog

import (
	"context"
	"loquigo/engine/database"
)

type Transition struct {
	Id         string `bson:"_id"`
	From       string `bson:"from"`
	To         string `bson:"to"`
	Conditions []Condition
	Actions    []Action
}

func (t *Transition) IsValid(event Event) bool {
	for _, condition := range t.Conditions {
		if condition.IsValid(event) {
			return true
		}
	}
	return false
}

func (t *Transition) Save() {
	db := database.GetMongoConnection()
	db.Collection(transitionCollection).InsertOne(context.TODO(), t)
}
