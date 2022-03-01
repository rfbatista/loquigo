package adapterservices

import "loquigo/engine/src/core/modules/templatepool"

type NodePosition struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type NodeHandler struct {
	Id string `json:"id"`
}

type NodeData struct {
	Name       string                   `json:"name"`
	Components []templatepool.Component `json:"components"`
	Handlers   []NodeHandler            `json:"handlers"`
}

type Node struct {
	Id       string       `json:"id"`
	Type     string       `json:"type"`
	Data     NodeData     `json:"data"`
	Position NodePosition `json:"position"`
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

func NewFlowMapService(flowService templatepool.FlowService, stepService templatepool.StepService, componentService templatepool.ComponentService) FlowMapService {
	return FlowMapService{flowService: flowService, stepService: stepService, componentService: componentService}
}

type FlowMapService struct {
	flowService      templatepool.FlowService
	stepService      templatepool.StepService
	componentService templatepool.ComponentService
}

func (f FlowMapService) GetMapFromFlow(flowId string) []Node {
	steps, _ := f.stepService.FindByFlowId(flowId)
	var nodes []Node
	for idx, step := range steps {
		nodes = append(nodes, createNodeFromStep(step, NodePosition{X: 200 + 20*idx, Y: 25 + 50*idx}))
	}
	return nodes
}

func createNodeFromStep(step templatepool.Step, position NodePosition) Node {
	node := Node{
		Id:   step.ID,
		Type: "step",
		Data: NodeData{
			Name: step.Name,
		},
		Position: position,
	}
	return node
}
