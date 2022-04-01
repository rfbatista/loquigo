package adapters

type Logger interface {
	Info(message string, context interface{})
	Error(message string, context interface{})
	Debug(message string, context interface{})
}
