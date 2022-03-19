package domain

type UserState struct {
	ID      string `bson:"_id"`
	UserId  string `bson:"user_id"`
	GroupId string `bson:"group_id"`
	NodeId  string `bson:"node_id"`
}

func NewUserState(userId, GroupId string, NodeId string) UserState {
	return UserState{UserId: userId, GroupId: GroupId, NodeId: NodeId}
}

func NewState(groupId string, nodeId string) State {
	return State{GroupId: groupId, NodeId: nodeId}
}

type State struct {
	GroupId string `bson:"grup_id"`
	NodeId  string `bson:"node_id"`
}
