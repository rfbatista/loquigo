package adapterservices

import (
	"log"
	"loquigo/engine/src/core/modules/template/pool"

	"gopkg.in/yaml.v2"
)

func NewEditor(f pool.FlowService, s pool.StepService, c pool.ComponentService) EditorService {
	return EditorService{flowService: f, stepService: s, componentService: c}
}

type BotEditor struct {
	ID    string       `yaml:"id"`
	Name  string       `yaml:"name"`
	Begin string       `yaml:"begin"`
	Flows []FlowEditor `yaml:"flows"`
}

type FlowEditor struct {
	ID    string       `yaml:"id"`
	Name  string       `yaml:"name"`
	Begin string       `yaml:"begin"`
	Steps []StepEditor `yaml:"steps"`
}

func (f *FlowEditor) ToDomain(botId string) pool.Flow {
	return pool.Flow{ID: f.ID, BotId: botId, Name: f.Name, BeginId: f.Begin}
}

func (f *FlowEditor) StepsDomain(botId string) []pool.Step {
	var steps []pool.Step
	for _, step := range f.Steps {
		steps = append(steps, step.ToDomain(botId, f.ID))
	}
	return steps
}

func (f *FlowEditor) AddStep(steps StepEditor) {
	f.Steps = append(f.Steps, steps)
}

func (f FlowEditor) ComponentsToDomain(botId string) []pool.Component {
	var components []pool.Component
	for _, step := range f.Steps {
		for _, component := range step.Components {
			components = append(components, component.ToDomain(botId, f.ID, step.ID))
		}
	}
	return components
}

func FlowDomainToEditorFlow(flow pool.Flow) FlowEditor {
	return FlowEditor{ID: flow.ID, Name: flow.Name, Begin: flow.BeginId}
}

type StepEditor struct {
	ID         string            `yaml:"id"`
	Name       string            `yaml:"name"`
	Components []ComponentEditor `yamls:"components"`
}

func (f StepEditor) ToDomain(botId string, flowId string) pool.Step {
	return pool.Step{ID: f.ID, BotId: botId, Name: f.Name, FlowId: flowId}
}

func StepDomainToEditorStep(step pool.Step) StepEditor {
	var components []ComponentEditor
	for _, domainComponent := range step.Components {
		components = append(components, DomainComponentToEditorComponent(domainComponent))
	}
	return StepEditor{ID: step.ID, Name: step.Name, Components: components}
}

type ComponentEditor struct {
	ID       string `yaml:"id"`
	Type     string `yaml:"type"`
	Text     string `yaml:"text,omitempty"`
	Flow     string `yaml:"flow,omitempty"`
	Step     string `yaml:"step,omitempty"`
	Sequence int    `yaml:"sequence"`
}

func (f ComponentEditor) ToDomain(botId string, flowId string, stepId string) pool.Component {
	return pool.Component{
		ID:       f.ID,
		BotId:    botId,
		FlowId:   flowId,
		StepId:   stepId,
		Sequence: f.Sequence,
		Type:     f.Type,
		Data: pool.ComponentData{
			Text:   f.Text,
			FlowId: f.Flow,
			StepId: f.Step,
		},
	}
}

func DomainComponentToEditorComponent(component pool.Component) ComponentEditor {
	return ComponentEditor{
		ID:       component.ID,
		Type:     component.Type,
		Sequence: component.Sequence,
		Flow:     component.Data.FlowId,
		Step:     component.Data.StepId,
		Text:     component.Data.Text,
	}
}

type EditorService struct {
	flowService      pool.FlowService
	stepService      pool.StepService
	componentService pool.ComponentService
}

func (e EditorService) UpdateBot(data string) (Result, error) {
	botID := "teste"
	e.flowService.DeleteByBotID(botID)
	e.stepService.DeleteByBotID(botID)
	e.componentService.DeleteByBotId(botID)
	var botSchema BotEditor
	err := yaml.Unmarshal([]byte(data), &botSchema)
	if err != nil {
		log.Fatalf("error: %v", err)
		return Result{}, err
	}
	for _, flow := range botSchema.Flows {
		steps := flow.StepsDomain(botID)
		for _, step := range steps {

			e.stepService.NewStep(step)
		}
		components := flow.ComponentsToDomain(botID)
		for _, component := range components {
			e.componentService.NewComponent(component)
		}
		e.flowService.NewFlow(flow.ToDomain(botID))
	}
	return Result{}, nil
}

func (e EditorService) FindBot(botId string) (BotEditor, error) {
	flows, _ := e.flowService.FindByBotId(botId)

	var editorFlows []FlowEditor
	for _, flow := range flows {
		editorFlow := FlowDomainToEditorFlow(flow)

		editorFlows = append(editorFlows, editorFlow)
	}

	var finalEditorFlows []FlowEditor
	for _, editorFlow := range editorFlows {

		domainSteps, _ := e.stepService.FindByFlowId(editorFlow.ID)

		for _, domainStep := range domainSteps {
			editorFlow.AddStep(StepDomainToEditorStep(domainStep))

			finalEditorFlows = append(finalEditorFlows, editorFlow)
		}
	}
	return BotEditor{Flows: finalEditorFlows}, nil
}
