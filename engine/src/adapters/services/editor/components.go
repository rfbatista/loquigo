package editorservice

import (
	"loquigo/engine/src/core/domain"
)

type BotEditor struct {
	Id      string        `yaml:"id"`
	Name    string        `yaml:"name"`
	Begin   string        `yaml:"begin"`
	Version string        `yaml:"version"`
	Groups  []GroupEditor `yaml:"groups"`
}

func (b BotEditor) ToDomain() domain.Bot {
	return domain.Bot{ID: b.Id, CurrentVersion: b.Version, BeginId: b.Begin, Name: b.Name}
}

type GroupEditor struct {
	ID    string       `yaml:"id"`
	Name  string       `yaml:"name"`
	Begin string       `yaml:"begin"`
	Nodes []StepEditor `yaml:"nodes"`
}

func (f *GroupEditor) ToDomain(botReference string) domain.Group {
	return domain.Group{ID: f.ID, BotReference: botReference, Name: f.Name, BeginId: f.Begin}
}

func (f *GroupEditor) NodesToDomain(botReference string) []domain.Node {
	var nodes []domain.Node
	for _, node := range f.Nodes {
		nodes = append(nodes, node.ToDomain(botReference, f.ID))
	}
	return nodes
}

func (f *GroupEditor) AddNode(steps StepEditor) {
	f.Nodes = append(f.Nodes, steps)
}

func (f GroupEditor) ComponentsToDomain(botReference string) []domain.Component {
	var components []domain.Component
	for _, node := range f.Nodes {
		for _, component := range node.Components {
			components = append(components, component.ToDomain(botReference, f.ID, node.ID))
		}
	}
	return components
}

func FlowDomainToEditorFlow(group domain.Group) GroupEditor {
	return GroupEditor{ID: group.ID, Name: group.Name, Begin: group.BeginId}
}

type StepEditor struct {
	ID         string            `yaml:"id"`
	Name       string            `yaml:"name"`
	Components []ComponentEditor `yamls:"components"`
}

func (f StepEditor) ToDomain(botReference string, groupId string) domain.Node {
	return domain.Node{ID: f.ID, BotReference: botReference, Name: f.Name, NodeId: groupId}
}

func NodeDomainToEditorNode(step domain.Node) StepEditor {
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
	GroupId  string `yaml:"group,omitempty"`
	NodeId   string `yaml:"node,omitempty"`
	Sequence int    `yaml:"sequence"`
}

func (f ComponentEditor) ToDomain(botReference string, groupId string, nodeId string) domain.Component {
	return domain.Component{
		ID:           f.ID,
		BotReference: botReference,
		GroupId:      groupId,
		NodeId:       nodeId,
		Sequence:     f.Sequence,
		Type:         f.Type,
		Data: domain.ComponentData{
			Text:    f.Text,
			GroupId: f.GroupId,
			NodeId:  f.NodeId,
		},
	}
}

func DomainComponentToEditorComponent(component domain.Component) ComponentEditor {
	return ComponentEditor{
		ID:       component.ID,
		Type:     component.Type,
		Sequence: component.Sequence,
		GroupId:  component.Data.GroupId,
		NodeId:   component.Data.NodeId,
		Text:     component.Data.Text,
	}
}
