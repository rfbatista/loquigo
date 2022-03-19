package nodes

func NewNodeRunner(node Node) RunnerNode {
	return NodeRunner{}
}

type NodeRunner struct {
	ID         string
	BotId      string `json:"bot_id"`
	GroupId    string `json:"group_id"`
	Name       string `json:"name"`
	Components []RunnerComponent
}

func (s NodeRunner) Run(message Message, context UserContext, messages []Message) ([]Message, *Stop, *GoTo) {
	var goTo *GoTo
	var stop *Stop
	var newMessages []Message = messages
	for _, component := range s.Components {
		newMessages, stop, goTo = component.Run(message, context, newMessages)
		if stop != nil {
			return newMessages, stop, nil
		}
		if goTo != nil {
			return newMessages, nil, goTo
		}
	}
	return []Message{}, nil, nil
}

func (s NodeRunner) AddComponents(components []RunnerComponent) (RunnerNode, error) {
	return NodeRunner{Components: components}, nil
}

func (s NodeRunner) IsValid() bool {
	if len(s.Components) > 0 {
		return true
	} else {
		return false
	}
}
