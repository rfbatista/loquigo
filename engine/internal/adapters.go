package internal

import (
	"github.com/google/wire"

	chat "loquigo/engine/pkg/adapters/transport/http"
)

var ChatProviderSet = wire.NewSet(
	EventManagerSet,
	chat.NewChatController)
