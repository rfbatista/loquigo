package nodes

import (
	"loquigo/engine/src/core/domain"
	"loquigo/engine/src/core/services/bot"
	"loquigo/engine/src/core/services/components"
)

type RunnerNode interface {
	Run(message domain.Message, context domain.UserContext, messages []domain.Message) ([]domain.Message, *domain.Stop, *domain.GoTo)
	AddComponents(components []components.RunnerComponent) (RunnerNode, error)
	IsValid() bool
}

func NewRunnerNodeService(bot bot.BotService, flow GroupRepository, stepRepo NodeRepository, c components.RunnerComponentService) RunnerNodeService {
	return RunnerNodeService{bot: bot, group: flow, stepRepo: stepRepo, componentService: c}
}

type RunnerNodeService struct {
	bot              bot.BotService
	group            GroupRepository
	stepRepo         NodeRepository
	componentService components.RunnerComponentService
}

func (r RunnerNodeService) FindBotBegin(botId string) (string, error) {
	flow, _ := r.bot.FindBotBegin(botId)
	return flow, nil
}

func (r RunnerNodeService) FindGroupBeginId(botId string, flowId string) (string, error) {
	nodeId, _ := r.group.FindBeginId(botId, flowId)
	return nodeId, nil
}

func (r RunnerNodeService) FindGroupBegin(botReference string, groupId string) (RunnerNode, error) {
	nodeId, _ := r.group.FindBeginId(botReference, groupId)
	node, _ := r.stepRepo.FindByGroupIdAndNodeId(botReference, groupId, nodeId)
	nodeRunner := NewNodeRunner(node)
	components, _ := r.componentService.FindByGroupIdAndNodeId(botReference, groupId, nodeId)
	nodeWithComponents, _ := nodeRunner.AddComponents(components)
	if nodeWithComponents.IsValid() == false {
		return nodeWithComponents, InvalidNode{BotReference: botReference, GroupId: groupId, NodeId: nodeId}
	}
	return nodeWithComponents, nil
}

func (r RunnerNodeService) FindByGroupIdAndNodeId(botReference string, groupId string, nodeId string) (RunnerNode, error) {
	node, _ := r.stepRepo.FindByGroupIdAndNodeId(botReference, groupId, nodeId)
	components, _ := r.componentService.FindByGroupIdAndNodeId(botReference, groupId, nodeId)
	nodeRunner := NewNodeRunner(node)
	nodeWithComponents, _ := nodeRunner.AddComponents(components)
	if nodeWithComponents.IsValid() == false {
		return nodeWithComponents, InvalidNode{BotReference: botReference, GroupId: groupId, NodeId: nodeId}
	}
	return nodeWithComponents, nil
}
