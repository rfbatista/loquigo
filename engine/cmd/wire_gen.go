// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cmd

import (
	"loquigo/engine/pkg/adapters/services/editor"
	"loquigo/engine/pkg/adapters/transport/http"
	"loquigo/engine/pkg/core/services/bot"
	"loquigo/engine/pkg/core/services/components"
	"loquigo/engine/pkg/core/services/dialogmanager"
	"loquigo/engine/pkg/core/services/eventmanager"
	"loquigo/engine/pkg/core/services/nodes"
	"loquigo/engine/pkg/core/services/runner"
	"loquigo/engine/pkg/infrastructure"
	"loquigo/engine/pkg/infrastructure/database/mongo"
	"loquigo/engine/pkg/infrastructure/database/mongo/repositories"
)

// Injectors from wire.go:

func InitializeEvent(db mongo.MongoDB) (infrastructure.Server, error) {
	groupRepository := repositories.NewGroupRepository(db)
	groupService := nodes.NewGroupService(groupRepository)
	nodeRepository := repositories.NewNodeRepository(db)
	componentRepository := repositories.NewComponentRepo(db)
	componentService := components.NewComponentService(componentRepository)
	nodeService := nodes.NewNodeService(nodeRepository, componentService)
	botRepository := repositories.NewBotRepository(db)
	botService := bot.NewBotService(botRepository)
	editorService := editorservice.NewEditor(groupService, nodeService, componentService, botService)
	editorController := adapters.NewEditorController(editorService)
	userStatesRepo := repositories.NewUserStatestRepo(db)
	runnerComponentService := components.NewRunnerComponentService(componentRepository)
	runnerNodeService := nodes.NewRunnerNodeService(botService, groupRepository, nodeRepository, runnerComponentService)
	runnerRunner := runner.NewRunner(runnerNodeService)
	runnerService := runner.NewRunnerService(userStatesRepo, runnerNodeService, runnerRunner)
	userContextRepository := repositories.NewUserContextRepo(db)
	findContextService := dialogmanager.NewFindContextService(userContextRepository)
	dialogManagerService := dialogmanager.NewDialogManagerService(runnerService, findContextService)
	chatService := eventmanager.NewChatService(dialogManagerService)
	chatController := adapters.NewChatController(chatService)
	botController := adapters.NewBotController(botService)
	server := infrastructure.NewServer(editorController, chatController, botController)
	return server, nil
}
