package infrastructure

import "go.uber.org/zap"

func NewLogger(log *zap.Logger) Logger {
	logger := Logger{}
	return logger.New(log)
}

type Logger struct {
	log *zap.Logger
}

func (l Logger) New(log *zap.Logger) Logger {
	return Logger{log: log}
}

func (l Logger) Info(message string, context interface{}) {
	l.log.Info(message)
}
func (l Logger) Error(message string, context interface{}) {
	l.log.Error(message)
}
func (l Logger) Debug(message string, context interface{}) {}
