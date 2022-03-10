package adapterservices

import (
	"fmt"
	"log"
	"loquigo/engine/src/core/modules/template/pool"

	"gopkg.in/yaml.v2"
)

func NewEditor(f pool.FlowService, s pool.StepService, c pool.ComponentService) EditorService {
	return EditorService{flowService: f, stepService: s, componentService: c}
}

type BotEditor struct {
	Flows []FlowEditor `yaml:"flows"`
}

type FlowEditor struct {
	ID    string       `yaml:"id"`
	Name  string       `yaml:"name"`
	Steps []StepEditor `yaml:"steps"`
}

func (f FlowEditor) ToDomain(botId string) pool.Flow {
	return pool.Flow{ID: f.ID, BotId: botId, Name: f.Name}
}

func (f FlowEditor) StepsDomain(botId string) []pool.Step {
	var steps []pool.Step
	for _, step := range f.Steps {
		steps = append(steps, step.ToDomain(botId, f.ID))
	}
	return steps
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

type StepEditor struct {
	ID         string            `yaml:"id"`
	Name       string            `yaml:"name"`
	Components []ComponentEditor `yamls:"components"`
}

func (f StepEditor) ToDomain(botId string, flowId string) pool.Step {
	return pool.Step{ID: f.ID, BotId: botId, Name: f.Name, FlowId: flowId}
}

type ComponentEditor struct {
	ID       string `yaml:"id"`
	Type     string `yaml:"type"`
	Text     string `yaml:"text"`
	Flow     string `yaml:"flow"`
	Step     string `yaml:"step"`
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
			fmt.Println(step)
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
