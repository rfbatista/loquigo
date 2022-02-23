package domain

type UserContext struct {
	ID     string            `bson:"_id,omitempty"`
	userId string            `bson:"user_id"`
	botId  string            `bson:"bot_id"`
	memory map[string]string `bson:"memory"`
}

func NewUserContext() UserContext {
	return UserContext{}
}

func (u UserContext) GetMemoryVariable(variable string) string {
	return u.memory[variable]
}
