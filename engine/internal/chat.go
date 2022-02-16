package internal

import (
	//Core
	"loquigo/engine/src/core/modules/contextmanager"
	"loquigo/engine/src/core/modules/dialogmanager"
	evm "loquigo/engine/src/core/modules/eventmanager"
	"loquigo/engine/src/core/modules/templatepool"

	//Infra
	infra "loquigo/engine/src/infrastructure"

	"github.com/google/wire"
)

var ChatSet = wire.NewSet(
	infra.NewHttpClient,
	dialogmanager.NewRunDialogService,
	evm.NewSendMessageService,
	contextmanager.NewFindContextService,
	templatepool.NewTemplatePoolService,
	evm.NewChatService)
