package cmd

import (
	editorservice "loquigo/engine/src/adapters/services/editor"
	adapters "loquigo/engine/src/adapters/transport/http"
	"loquigo/engine/src/core/modules/bot"
	"loquigo/engine/src/core/modules/components"
	"loquigo/engine/src/core/modules/dialogmanager"
	"loquigo/engine/src/core/modules/eventmanager"
	"loquigo/engine/src/core/modules/nodes"
	"loquigo/engine/src/core/modules/runner"
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
	wire.Bind(new(runner.UserStateRepo), new(repositories.UserStatesRepo)),
)

var UserContextSet = wire.NewSet(
	repositories.NewUserContextRepo,
	wire.Bind(new(dialogmanager.UserContextRepository), new(repositories.UserContextRepository)),
)

var GroupRepoSet = wire.NewSet(
	repositories.NewGroupRepository,
	wire.Bind(new(nodes.GroupRepository), new(repositories.GroupRepository)),
)

var NodeRepoSet = wire.NewSet(
	repositories.NewNodeRepository,
	wire.Bind(new(nodes.NodeRepository), new(repositories.NodeRepository)),
)

var ComponentRepoSet = wire.NewSet(
	repositories.NewComponentRepo,
	wire.Bind(new(components.ComponentRepository), new(repositories.ComponentRepository)),
)

var BotRepoSet = wire.NewSet(
	repositories.NewBotRepository,
	wire.Bind(new(bot.BotRepository), new(repositories.BotRepository)),
)

//*****************
// Services
//*****************

var ContextServiceSet = wire.NewSet(
	UserContextSet,
	dialogmanager.NewFindContextService,
)

var BotServiceSet = wire.NewSet(
	BotRepoSet,
	bot.NewBotService,
)

var ComponentServiceSet = wire.NewSet(
	ComponentRepoSet,
	components.NewComponentService,
	components.NewRunnerComponentService,
)

var NodeServiceSet = wire.NewSet(
	ComponentServiceSet,
	BotServiceSet,
	NodeRepoSet,
	GroupRepoSet,
	nodes.NewGroupService,
	nodes.NewNodeService,
	nodes.NewRunnerNodeService,
)

var RunnerServiceSet = wire.NewSet(
	UserStateRepoSet,
	NodeServiceSet,
	runner.NewRunner,
	runner.NewRunnerService,
)

var DialogManagerServiceSet = wire.NewSet(
	RunnerServiceSet,
	ContextServiceSet,
	dialogmanager.NewDialogManagerService,
)

var ChatServiceSet = wire.NewSet(
	DialogManagerServiceSet,
	eventmanager.NewChatService,
)

//*****************
// Controller
//*****************

var ControllersSet = wire.NewSet(
	ChatServiceSet,
	editorservice.NewEditor,
	adapters.NewEditorController,
	adapters.NewChatController,
	adapters.NewBotController,
)
