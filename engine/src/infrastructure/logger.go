package infrastructure

func NewLogger() Logger {
	return Logger{}
}

type Logger struct{}

func (l Logger) Info(message string, context interface{})  {}
func (l Logger) Error(message string, context interface{}) {}
func (l Logger) Debug(message string, context interface{}) {}
