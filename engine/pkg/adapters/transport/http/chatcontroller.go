package adapters

import (
	"github.com/gin-gonic/gin"

	evm "loquigo/engine/pkg/core/modules/eventmanager"
)

func NewChatController(e evm.ChatService) ChatController {
	return ChatController{eventService: e}
}

type ChatController struct {
	eventService evm.ChatService
}

func (chat ChatController) PostMessage(c *gin.Context) {

}
