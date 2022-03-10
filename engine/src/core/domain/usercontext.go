package domain

type UserContext struct {
	ID     string
	UserId string
	BotId  string
	User   User
	Memory map[string]string
}

func NewUserContext() UserContext {
	return UserContext{}
}

func (u UserContext) GetMemoryVariable(variable string) string {
	return u.Memory[variable]
}
