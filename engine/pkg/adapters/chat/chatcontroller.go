package chat

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"loquigo/engine/pkg/core/domain"
	"loquigo/engine/pkg/core/services/eventmanager"
)

func NewChatController(e eventmanager.ChatService) ChatController {
	return ChatController{eventService: e}
}

type ChatController struct {
	eventService eventmanager.ChatService
}

func (r ChatController) AddChatRoutes(rg *gin.RouterGroup) {
	route := rg.Group("/chat")

	route.POST("/", r.PostMessage)
}

func (chat ChatController) PostMessage(c *gin.Context) {
	var input domain.Event
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	messages, err := chat.eventService.Run(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": messages})
	}
}
