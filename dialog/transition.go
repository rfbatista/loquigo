package dialog

type Transition struct {
	from       string
	to         string
	conditions []Condition
	actions    []Action
}

func (t Transition) IsValid(event Event) bool {
	for _, condition := range t.conditions {
		if condition.IsValid(event) {
			return true
		}
	}
	return false
}

type TransitionDAO struct {
	From string
	To   string
}
