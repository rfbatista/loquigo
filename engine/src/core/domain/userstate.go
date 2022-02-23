package domain

type UserState struct {
	ID     string `bson:"_id"`
	UserId string `bson:"user_id"`
	FlowId string `bson:"flow_id"`
	StepId string `bson:"step_id"`
}

func NewUserState(userId, flowId string, stepID string) UserState {
	return UserState{UserId: userId, FlowId: flowId, StepId: stepID}
}

func NewState(flowId string, stepID string) State {
	return State{FlowId: flowId, StepId: stepID}
}

type State struct {
	FlowId string `bson:"flow_id"`
	StepId string `bson:"step_id"`
}
