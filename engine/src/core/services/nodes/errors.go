package nodes

import "fmt"

type InvalidNode struct {
	BotReference string
	GroupId      string
	NodeId       string
}

func (n InvalidNode) Error() string {
	return fmt.Sprintf("Invalid node in BotReference: %s, GroupId: %s, NodeId: %s ", n.BotReference, n.GroupId, n.NodeId)
}
