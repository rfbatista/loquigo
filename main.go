package main

import (
	"go.uber.org/zap"
)

func main() {
	// cfg := infrastructure.NewConfig()
	zapLogger, _ := zap.NewDevelopment()
	defer zapLogger.Sync()
	// logger := infrastructure.NewLogger(zapLogger)
	// db := mongo.NewMongoDb(cfg.Database, &logger)
	// db.Connect()
	// defer db.Disconnect()
	// e, _ := cmd.InitializeEvent(db, zapLogger, cfg)
	// e.Start()
}
