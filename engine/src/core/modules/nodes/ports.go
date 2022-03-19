package nodes

import "loquigo/engine/src/core/domain"

type UserStateRepo interface {
	FindByUserId(userId string) (domain.UserState, error)
	Update(userState domain.UserState) error
	Create(userId string) (domain.UserState, error)
}

type GroupRepository interface {
	FindByBotId(id string) ([]Group, error)
	Create(group Group) (Group, error)
	Update(group Group) (Group, error)
	Delete(group Group) (Group, error)
	DeleteByBotID(botReference string) error
	FindBeginId(botReference string, groupId string) (string, error)
}

type NodeRepository interface {
	FindByGroupId(id string) ([]Node, error)
	FindById(id string) (Node, error)
	Create(node Node) (Node, error)
	Update(node Node) (Node, error)
	Delete(node Node) (Node, error)
	DeleteByBotID(botReference string) error
	FindByIdAndGroupId(groupId string, nodeId string) (Node, error)
	FindByGroupIdAndNodeId(botReference string, groupId string, nodeId string) (Node, error)
}
