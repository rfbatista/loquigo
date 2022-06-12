package logger

import "go.uber.org/zap"

var logger *zap.Logger

func Start(log *zap.Logger) {
	logger = log
}

func Info(message string) {
	logger.Info(message)
}
func Error(message string) {
	logger.Error(message)
}
func Debug(message string) {
	logger.Debug(message)
}
