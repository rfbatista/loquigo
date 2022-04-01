package editor

import (
	"loquigo/engine/pkg/adapters"
	"loquigo/engine/pkg/core/domain"
	"loquigo/engine/pkg/core/services/bot"
	"loquigo/engine/pkg/core/services/components"
	"loquigo/engine/pkg/core/services/nodes"

	"gopkg.in/yaml.v2"
)

func NewEditor(f nodes.GroupService, s nodes.NodeService, c components.ComponentService, b bot.BotService, l adapters.Logger) EditorService {
	return EditorService{groupService: f, nodeService: s, componentService: c, botService: b, logger: l}
}

type EditorService struct {
	botService       bot.BotService
	groupService     nodes.GroupService
	nodeService      nodes.NodeService
	componentService components.ComponentService
	logger           adapters.Logger
}

func (e EditorService) UpdateBot(data string, version string, botId string) (adapters.Result, error) {
	var botSchema BotEditor
	err := yaml.Unmarshal([]byte(data), &botSchema)
	if err != nil {
		e.logger.Error("error: %v", err)
		return adapters.Result{}, err
	}
	bot, err := e.botService.UpdateBot(botSchema.ToDomain(version, botId))
	e.botService.CreateVersion(bot, version, data)
	if err != nil {
		e.logger.Error("error: %v", err)
		return adapters.Result{}, err
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
	return adapters.Result{}, nil
}

func (e EditorService) FindBot(botId string) (BotEditor, error) {
	bot, _ := e.botService.FindBotById(botId)
	botEditor := BotDomainToEditor(bot)
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
	botEditor.AddGroups(finalEditorFlows)
	return botEditor, nil
}

func (e EditorService) FindBotVersions(botId string) ([]domain.BotVersion, error) {
	versions, _ := e.botService.FindVersionsByBotId(botId)
	return versions, nil
}

func (e EditorService) FindVersionByIdAndBotId(versionId string, botId string) (domain.BotVersion, error) {
	version, _ := e.botService.FindVersionByIdAndBotId(versionId, botId)
	return version, nil
}
