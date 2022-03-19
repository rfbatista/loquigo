package adapterservices

import (
	"loquigo/engine/src/core/modules/components"
	"loquigo/engine/src/core/modules/nodes"
)

type NodePosition struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type NodeHandler struct {
	Id string `json:"id"`
}

type NodeData struct {
	ID         string                 `json:"id"`
	FlowID     string                 `json:"flowId"`
	Name       string                 `json:"name"`
	Components []components.Component `json:"components"`
	Handlers   []NodeHandler          `json:"handlers"`
}

type Node struct {
	Id       string       `json:"id"`
	Type     string       `json:"type"`
	Data     NodeData     `json:"data"`
	Position NodePosition `json:"position"`
}

func (n Node) ToDomain() nodes.Node {
	step := nodes.NewStep(n.Data.ID, n.Data.FlowID, n.Data.Name)
	step.Components = n.Data.Components
	return step
}

type Connection struct {
	Id           string
	Source       string
	Target       string
	Animated     bool
	TargetHandle string
	SourceHandle string
}

type FlowMap struct {
}

func NewFlowMapService(flowService nodes.GroupService, stepService nodes.NodeService, componentService components.ComponentService) FlowMapService {
	return FlowMapService{flowService: flowService, stepService: stepService, componentService: componentService}
}

type FlowMapService struct {
	flowService      nodes.GroupService
	stepService      nodes.NodeService
	componentService components.ComponentService
}

func (f FlowMapService) GetMapFromFlow(flowId string) []Node {
	steps, _ := f.stepService.FindByGroupId(flowId)
	var nodes []Node
	for idx, step := range steps {
		nodes = append(nodes, createNodeFromStep(step, NodePosition{X: 200 + 20*idx, Y: 25 + 50*idx}))
	}
	return nodes
}

func createNodeFromStep(step nodes.Node, position NodePosition) Node {
	node := Node{
		Id:   step.ID,
		Type: "step",
		Data: NodeData{
			ID:         step.ID,
			FlowID:     step.NodeId,
			Name:       step.Name,
			Components: step.Components,
		},
		Position: position,
	}
	return node
}