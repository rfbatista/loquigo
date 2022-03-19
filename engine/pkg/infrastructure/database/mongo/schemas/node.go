package schemas

import "loquigo/engine/pkg/core/domain"

func NewNodeSchema(node domain.Node) (NodeSchema, error) {
	return NodeSchema{
		ID:           node.ID,
		BotReference: node.BotReference,
		NodeId:       node.NodeId,
		Name:         node.Name,
	}, nil
}

type NodeSchema struct {
	ID           string `bson:"id"`
	BotReference string `bson:"bot_reference"`
	NodeId       string `bson:"node_id"`
	Name         string `bson:"name"`
}

func (s NodeSchema) ToDomain() domain.Node {
	return domain.Node{ID: s.ID, BotReference: s.BotReference, NodeId: s.NodeId, Name: s.Name}
}
