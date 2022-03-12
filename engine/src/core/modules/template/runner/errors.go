package runner

import "fmt"

type NoStepInBotError struct {
	BotId string
}

func (n NoStepInBotError) Error() string {
	return fmt.Sprintf("bot with id:%s doest not have a valid step", n.BotId)
}

type MissingState struct {
	UserId string
}

func (n MissingState) Error() string {
	return fmt.Sprintf("user with id:%s doest not have a valid state", n.UserId)
}

type InvalidStep struct {
	BotId  string
	FlowId string
	StepId string
}

func (n InvalidStep) Error() string {
	return fmt.Sprintf("Invalid step in BotId: %s, FlowId: %s, StepId: %s ", n.BotId, n.FlowId, n.StepId)
}
