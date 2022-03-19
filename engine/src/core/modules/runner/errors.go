package runner

import "fmt"

type NoStepInBotError struct {
	BotReference string
}

func (n NoStepInBotError) Error() string {
	return fmt.Sprintf("bot with reference:%s doest not have a valid step", n.BotReference)
}

type MissingUserState struct {
	UserId string
}

func (n MissingUserState) Error() string {
	return fmt.Sprintf("user with id:%s doest not have a valid state", n.UserId)
}

type MissingValidBotState struct {
	BotReference string
}

func (n MissingValidBotState) Error() string {
	return fmt.Sprintf("bot with reference:%s doest not have a valid state", n.BotReference)
}
