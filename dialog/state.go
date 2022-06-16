package dialog

func Start(event Event) State {
	return State{}
}

type State struct {
	Id          string
	Name        string
	Transitions []Transition
}

func (s State) Next(event Event) (State, *string, error) {
	return State{}, nil, nil
}

type StateDAO struct {
	Id   string
	Name string
}
