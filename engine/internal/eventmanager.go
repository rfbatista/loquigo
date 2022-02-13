package internal

import (
	//Core
	evm "loquigo/engine/pkg/core/modules/eventmanager"

	//Infra
	infra "loquigo/engine/pkg/infrastructure"

	"github.com/google/wire"
)

var EventManagerSet = wire.NewSet(
	infra.NewHttpClient,
	evm.NewSendMessageService,
	evm.NewChatService)
