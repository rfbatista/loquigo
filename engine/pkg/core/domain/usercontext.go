package domain

type UserContext struct {
	ID     string
	memory map[string]string
}

func (u UserContext) GetMemoryVariable(variable string) string {
	return u.memory[variable]
}
