package cmd

import (
	adapterservices "loquigo/engine/src/adapters/services"
	adapters "loquigo/engine/src/adapters/transport/http"
	"loquigo/engine/src/core/modules/dialogmanager"
	"loquigo/engine/src/core/modules/eventmanager"
	"loquigo/engine/src/core/modules/templatepool"
	infra "loquigo/engine/src/infrastructure"
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
	wire.Bind(new(templatepool.UserStateRepo), new(repositories.UserStatesRepo)),
)

var UserContextSet = wire.NewSet(
	repositories.NewUserContextRepo,
	wire.Bind(new(dialogmanager.UserContextRepository), new(repositories.UserContextRepository)),
)

var FlowRepoSet = wire.NewSet(
	repositories.NewFlowRepository,
	wire.Bind(new(templatepool.FlowRepository), new(repositories.FlowRepository)),
)

var StepRepoSet = wire.NewSet(
	repositories.NewStepRepository,
	wire.Bind(new(templatepool.StepRepository), new(repositories.StepRepository)),
)

var ComponentRepoSet = wire.NewSet(
	repositories.NewComponentRepo,
	wire.Bind(new(templatepool.ComponentRepository), new(repositories.ComponentRepository)),
)

//*****************
// Services
//*****************

var ContextSet = wire.NewSet(
	UserContextSet,
	dialogmanager.NewFindContextService,
)

var TemplatePoolSet = wire.NewSet(
	UserStateRepoSet,
	templatepool.NewTemplatePoolService,
)

var ChatServiceSet = wire.NewSet(
	infra.NewHttpClient,
	ContextSet,
	TemplatePoolSet,
	dialogmanager.NewRunDialogService,
	eventmanager.NewSendMessageService,
	UserRepoSet,
	eventmanager.NewChatService,
)

var FlowServiceSet = wire.NewSet(
	FlowRepoSet,
	templatepool.NewFlowService,
)

var StepServiceSet = wire.NewSet(
	StepRepoSet,
	templatepool.NewStepService,
)

var ComponentServiceSet = wire.NewSet(
	ComponentRepoSet,
	templatepool.NewComponentService,
)

var FlowMapServiceSet = wire.NewSet(
	adapterservices.NewFlowMapService,
)

//*****************
// Controllers
//*****************

var ChatSet = wire.NewSet(
	ChatServiceSet,
	adapters.NewChatController)

var FlowSet = wire.NewSet(
	FlowServiceSet,
	adapters.NewFlowController,
)

var StepSet = wire.NewSet(
	StepServiceSet,
	adapters.NewStepController,
)

var ComponentSet = wire.NewSet(
	ComponentServiceSet,
	adapters.NewComponentController,
)

var FlowMapSet = wire.NewSet(
	FlowMapServiceSet,
	adapters.NewFlowMapController,
)
