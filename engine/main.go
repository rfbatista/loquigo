package main

import (
	"loquigo/engine/cmd"
	"loquigo/engine/src/infrastructure"
	"loquigo/engine/src/infrastructure/database/mongo"
)

func main() {
	cfg := infrastructure.NewConfig()
	db := mongo.NewMongoDb(cfg)
	db.Connect()
	defer db.Disconnect()
	e, _ := cmd.InitializeEvent(db)
	e.Start()
}
