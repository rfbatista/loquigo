package cmd

import (
	adapterservices "loquigo/engine/src/adapters/services"
	adapters "loquigo/engine/src/adapters/transport/http"
	"loquigo/engine/src/core/modules/dialogmanager"
	"loquigo/engine/src/core/modules/eventmanager"
	"loquigo/engine/src/core/modules/template/pool"
	"loquigo/engine/src/core/modules/template/runner"
	"loquigo/engine/src/infrastructure/database/mongo/repositories"

	"github.com/google/wire"
)

//*****************
// Repositories
//*****************

var UserRepoSet = wire.NewSet(
	repositories.NewUserRepository,
	wire.Bind(new(eventmanager.UserRepository), new(repositories.UserRepository)),
)

var UserStateRepoSet = wire.NewSet(
	repositories.NewUserStatestRepo,
	wire.Bind(new(pool.UserStateRepo), new(repositories.UserStatesRepo)),
	wire.Bind(new(runner.UserStateRepo), new(repositories.UserStatesRepo)),
)

var UserContextSet = wire.NewSet(
	repositories.NewUserContextRepo,
	wire.Bind(new(dialogmanager.UserContextRepository), new(repositories.UserContextRepository)),
)

var FlowRepoSet = wire.NewSet(
	repositories.NewFlowRepository,
	wire.Bind(new(pool.FlowRepository), new(repositories.FlowRepository)),
	wire.Bind(new(runner.FlowRepository), new(repositories.FlowRepository)),
)

var StepRepoSet = wire.NewSet(
	repositories.NewStepRepository,
	wire.Bind(new(pool.StepRepository), new(repositories.StepRepository)),
	wire.Bind(new(runner.StepRepository), new(repositories.StepRepository)),
)

var ComponentRepoSet = wire.NewSet(
	repositories.NewComponentRepo,
	wire.Bind(new(pool.ComponentRepository), new(repositories.ComponentRepository)),
	wire.Bind(new(runner.ComponentRepository), new(repositories.ComponentRepository)),
)

var BotRepoSet = wire.NewSet(
	repositories.NewBotRepository,
	wire.Bind(new(runner.BotRepository), new(repositories.BotRepository)),
)

//*****************
// Services
//*****************

var ContextSet = wire.NewSet(
	UserContextSet,
	dialogmanager.NewFindContextService,
)

// var TemplatePoolSet = wire.NewSet(
// 	UserStateRepoSet,
// 	pool.NewTemplatePoolService,
// )

var ChatServiceSet = wire.NewSet(
	ComponentRepoSet,
	ContextSet,
	UserStateRepoSet,
	// TemplatePoolSet,
	FlowRepoSet,
	StepRepoSet,
	BotRepoSet,
	UserRepoSet,
	runner.NewRunnerStepService,
	runner.NewRunnerService,
	runner.NewChatRunnerService,
	pool.NewComponentService,
	pool.NewFlowService,
	pool.NewStepService,
	dialogmanager.NewRunDialogService,
	eventmanager.NewSendMessageService,
	eventmanager.NewChatService,
)

var ChatAndEditorServiceSet = wire.NewSet(
	ChatServiceSet,
	adapterservices.NewEditor,
	adapters.NewEditorController,
	adapters.NewChatController,
)
