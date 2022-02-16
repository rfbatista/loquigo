package internal

import (
	"github.com/google/wire"

	chat "loquigo/engine/src/adapters/transport/http"
)

var ChatProviderSet = wire.NewSet(
	ChatSet,
	chat.NewChatController)
