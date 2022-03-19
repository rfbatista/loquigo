package editorservice

import (
	"log"
	adapterservices "loquigo/engine/src/adapters/services"
	"loquigo/engine/src/core/domain"
	"loquigo/engine/src/core/services/bot"
	"loquigo/engine/src/core/services/components"
	"loquigo/engine/src/core/services/nodes"

	"gopkg.in/yaml.v2"
)

func NewEditor(f nodes.GroupService, s nodes.NodeService, c components.ComponentService, b bot.BotService) EditorService {
	return EditorService{groupService: f, nodeService: s, componentService: c, botService: b}
}

type EditorService struct {
	botService       bot.BotService
	groupService     nodes.GroupService
	nodeService      nodes.NodeService
	componentService components.ComponentService
}

func (e EditorService) UpdateBot(data string) (adapterservices.Result, error) {
	var botSchema BotEditor
	err := yaml.Unmarshal([]byte(data), &botSchema)
	if err != nil {
		log.Fatalf("error: %v", err)
		return adapterservices.Result{}, err
	}
	bot, err := e.botService.CreateBot(botSchema.ToDomain())
	if err != nil {
		return adapterservices.Result{}, err
	}
	botReference := bot.Reference()
	for _, group := range botSchema.Groups {
		steps := group.NodesToDomain(botReference)
		for _, step := range steps {
			e.nodeService.NewNode(step)
		}
		components := group.ComponentsToDomain(botReference)
		for _, component := range components {
			e.componentService.NewComponent(component)
		}
		e.groupService.NewGroup(group.ToDomain(botReference))
	}
	return adapterservices.Result{}, nil
}

func (e EditorService) FindBot(botId string) (BotEditor, error) {
	flows, _ := e.groupService.FindByBotId(botId)

	var editorFlows []GroupEditor
	for _, flow := range flows {
		editorFlow := FlowDomainToEditorFlow(flow)

		editorFlows = append(editorFlows, editorFlow)
	}

	var finalEditorFlows []GroupEditor
	for _, editorFlow := range editorFlows {

		domainSteps, _ := e.nodeService.FindByGroupId(editorFlow.ID)

		for _, domainStep := range domainSteps {
			editorFlow.AddNode(NodeDomainToEditorNode(domainStep))

			finalEditorFlows = append(finalEditorFlows, editorFlow)
		}
	}
	return BotEditor{Groups: finalEditorFlows}, nil
}

func (e EditorService) FindBotVersions(botId string) ([]domain.BotVersion, error) {
	versions, _ := e.botService.FindVersionsByBotId(botId)
	return versions, nil
}

func (e EditorService) FindVersionByIdAndBotId(versionId string, botId string) (domain.BotVersion, error) {
	version, _ := e.botService.FindVersionByIdAndBotId(versionId, botId)
	return version, nil
}
