package nodes

import "loquigo/engine/src/core/domain"

type UserStateRepo interface {
	FindByUserId(userId string) (domain.UserState, error)
	Update(userState domain.UserState) error
	Create(userId string) (domain.UserState, error)
}

type GroupRepository interface {
	FindByBotId(id string) ([]domain.Group, error)
	Create(group domain.Group) (domain.Group, error)
	Update(group domain.Group) (domain.Group, error)
	Delete(group domain.Group) (domain.Group, error)
	DeleteByBotID(botReference string) error
	FindBeginId(botReference string, groupId string) (string, error)
}

type NodeRepository interface {
	FindByGroupId(id string) ([]domain.Node, error)
	FindById(id string) (domain.Node, error)
	Create(node domain.Node) (domain.Node, error)
	Update(node domain.Node) (domain.Node, error)
	Delete(node domain.Node) (domain.Node, error)
	DeleteByBotID(botReference string) error
	FindByIdAndGroupId(groupId string, nodeId string) (domain.Node, error)
	FindByGroupIdAndNodeId(botReference string, groupId string, nodeId string) (domain.Node, error)
}
