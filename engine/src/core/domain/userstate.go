package domain

type UserState struct {
	ID     string `bson:"_id"`
	UserId string `bson:"user_id"`
	FlowId string `bson:"flow_id"`
	StepId string `bson:"step_id"`
}
