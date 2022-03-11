// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cmd

import (
	"loquigo/engine/src/adapters/services"
	"loquigo/engine/src/adapters/transport/http"
	"loquigo/engine/src/core/modules/dialogmanager"
	"loquigo/engine/src/core/modules/eventmanager"
	"loquigo/engine/src/core/modules/template/pool"
	"loquigo/engine/src/core/modules/template/runner"
	"loquigo/engine/src/infrastructure"
	"loquigo/engine/src/infrastructure/database/mongo"
	"loquigo/engine/src/infrastructure/database/mongo/repositories"
)

// Injectors from wire.go:

func InitializeEvent(db mongo.MongoDB) (infrastructure.Server, error) {
	flowRepository := repositories.NewFlowRepository(db)
	flowService := pool.NewFlowService(flowRepository)
	stepRepository := repositories.NewStepRepository(db)
	componentRepository := repositories.NewComponentRepo(db)
	componentService := pool.NewComponentService(componentRepository)
	stepService := pool.NewStepService(stepRepository, componentService)
	editorService := adapterservices.NewEditor(flowService, stepService, componentService)
	editorController := adapters.NewEditorController(editorService)
	userStatesRepo := repositories.NewUserStatestRepo(db)
	botRepository := repositories.NewBotRepository(db)
	runnerStepService := runner.NewRunnerStepService(botRepository, flowRepository, stepRepository, componentRepository)
	runnerRunner := runner.NewRunnerService(runnerStepService)
	runnerService := runner.NewChatRunnerService(userStatesRepo, runnerStepService, runnerRunner)
	userContextRepository := repositories.NewUserContextRepo(db)
	findContextService := dialogmanager.NewFindContextService(userContextRepository)
	runDialogService := dialogmanager.NewRunDialogService(runnerService, findContextService)
	userRepository := repositories.NewUserRepository(db)
	chatService := eventmanager.NewChatService(runDialogService, userRepository)
	chatController := adapters.NewChatController(chatService)
	server := infrastructure.NewServer(editorController, chatController)
	return server, nil
}
