// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cmd

import (
	"go.uber.org/zap"
	bot2 "loquigo/engine/pkg/adapters/bot"
	"loquigo/engine/pkg/adapters/chat"
	"loquigo/engine/pkg/adapters/editor"
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

func InitializeEvent(db mongo.MongoDB, logger *zap.Logger) (infrastructure.Server, error) {
	groupRepository := repositories.NewGroupRepository(db)
	groupService := nodes.NewGroupService(groupRepository)
	nodeRepository := repositories.NewNodeRepository(db)
	componentRepository := repositories.NewComponentRepo(db)
	componentService := components.NewComponentService(componentRepository)
	nodeService := nodes.NewNodeService(nodeRepository, componentService)
	botRepository := repositories.NewBotRepository(db)
	botVersionRepository := repositories.NewBotVersionRepository(db)
	botService := bot.NewBotService(botRepository, botVersionRepository)
	infrastructureLogger := infrastructure.NewLogger(logger)
	editorService := editor.NewEditor(groupService, nodeService, componentService, botService, infrastructureLogger)
	editorController := editor.NewEditorController(editorService)
	userStatesRepo := repositories.NewUserStatestRepo(db)
	runnerComponentService := components.NewRunnerComponentService(componentRepository)
	runnerNodeService := nodes.NewRunnerNodeService(botService, groupRepository, nodeRepository, runnerComponentService)
	runnerRunner := runner.NewRunner(runnerNodeService)
	runnerService := runner.NewRunnerService(userStatesRepo, runnerNodeService, runnerRunner)
	userContextRepository := repositories.NewUserContextRepo(db)
	findContextService := dialogmanager.NewFindContextService(userContextRepository)
	dialogManagerService := dialogmanager.NewDialogManagerService(runnerService, findContextService)
	chatService := eventmanager.NewChatService(dialogManagerService)
	chatController := chat.NewChatController(chatService)
	botController := bot2.NewBotController(botService)
	server := infrastructure.NewServer(editorController, chatController, botController)
	return server, nil
}
