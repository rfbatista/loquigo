package main

import (
	"loquigo/engine/cmd"
	"loquigo/engine/pkg/infrastructure"
	"loquigo/engine/pkg/infrastructure/database/mongo"

	"go.uber.org/zap"
)

func main() {
	cfg := infrastructure.NewConfig()
	zapLogger, _ := zap.NewDevelopment()
	defer zapLogger.Sync()
	db := mongo.NewMongoDb(cfg)
	db.Connect()
	defer db.Disconnect()
	e, _ := cmd.InitializeEvent(db, zapLogger)
	e.Start()
}
