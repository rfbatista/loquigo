package cmd

import (
	chat "loquigo/engine/src/adapters/transport/http"
	"loquigo/engine/src/core/modules/contextmanager"
	"loquigo/engine/src/core/modules/dialogmanager"
	"loquigo/engine/src/core/modules/eventmanager"
	"loquigo/engine/src/core/modules/templatepool"
	infra "loquigo/engine/src/infrastructure"
	"loquigo/engine/src/infrastructure/database/mongo/repositories"

	"github.com/google/wire"
)

var UserContextSet = wire.NewSet(
	repositories.NewUserContextRepo,
	wire.Bind(new(contextmanager.UserContextRepository), new(repositories.UserContextRepository)),
)

var ContextSet = wire.NewSet(
	UserContextSet,
	contextmanager.NewFindContextService,
)
var StateSet = wire.NewSet(
	repositories.NewUserStatestRepo,
	templatepool.NewTemplatePoolService,
)

var ChatSet = wire.NewSet(
	infra.NewHttpClient,
	ContextSet,
	StateSet,
	dialogmanager.NewRunDialogService,
	eventmanager.NewSendMessageService,
	eventmanager.NewChatService,
	chat.NewChatController)
