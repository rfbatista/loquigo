// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cmd

import (
	"loquigo/engine/src/adapters/services"
	"loquigo/engine/src/adapters/transport/http"
	"loquigo/engine/src/core/modules/template/pool"
	"loquigo/engine/src/infrastructure"
	"loquigo/engine/src/infrastructure/database/mongo"
	"loquigo/engine/src/infrastructure/database/mongo/repositories"
)

// Injectors from wire.go:

func InitializeEvent(db mongo.MongoDB) (infrastructure.Server, error) {
	flowRepository := repositories.NewFlowRepository(db)
	flowService := pool.NewFlowService(flowRepository)
	flowController := adapters.NewFlowController(flowService)
	stepRepository := repositories.NewStepRepository(db)
	componentRepository := repositories.NewComponentRepo(db)
	componentService := pool.NewComponentService(componentRepository)
	stepService := pool.NewStepService(stepRepository, componentService)
	stepController := adapters.NewStepController(stepService)
	componentController := adapters.NewComponentController(componentService)
	flowMapService := adapterservices.NewFlowMapService(flowService, stepService, componentService)
	flowMapController := adapters.NewFlowMapController(flowMapService)
	server := infrastructure.NewServer(flowController, stepController, componentController, flowMapController)
	return server, nil
}
