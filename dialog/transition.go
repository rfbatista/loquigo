package dialog

type Transition struct {
	from       string
	to         string
	conditions []Condition
	actions    []Action
}
