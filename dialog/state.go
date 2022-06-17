package dialog

import (
	"context"
	"loquigo/engine/database"
)

type StateType int16

const (
	UndefinedState StateType = iota
	HoldState
	DefaultState
)

func (m StateType) String() string {
	switch m {
	case HoldState:
		return "hold"
	case DefaultState:
		return "default"
	}
	return "unknown"
}

type State struct {
	Id          string `bson:"_id"`
	Name        string `bson:"name"`
	Type        StateType
	Transitions []Transition
}

func (s State) Next(event Event) (State, *string, error) {
	return State{}, nil, nil
}

func (s *State) Save() {
	db := database.GetMongoConnection()
	db.Collection(stateCollection).InsertOne(context.TODO(), s)
}
