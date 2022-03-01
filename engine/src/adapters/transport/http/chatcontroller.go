package adapters

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"loquigo/engine/src/core/domain"
	"loquigo/engine/src/core/modules/eventmanager"
)

func NewChatController(e eventmanager.ChatService) ChatController {
	return ChatController{eventService: e}
}

func (r HttpRouter) AddChatRoutes(rg *gin.RouterGroup, controller ChatController) {
	route := rg.Group("/chat")

	route.POST("/", controller.PostMessage)
}

type ChatController struct {
	eventService eventmanager.ChatService
}

func (chat ChatController) PostMessage(c *gin.Context) {
	var input domain.Event
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	messages, _ := chat.eventService.Run(input)
	c.JSON(http.StatusOK, gin.H{"data": messages})
}
