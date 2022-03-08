package adapters

import (
	"loquigo/engine/src/core/domain"
	"loquigo/engine/src/core/modules/eventmanager"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewEditorController(e eventmanager.ChatService) ChatController {
	return ChatController{eventService: e}
}

func (r HttpRouter) AddEditorRoutes(rg *gin.RouterGroup, controller ChatController) {
	route := rg.Group("/chat")

	route.POST("/", controller.PostMessage)
}

type EditorController struct {
	eventService eventmanager.ChatService
}

func (chat ChatController) UpdateBot(c *gin.Context) {
	var input domain.Event
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	messages, _ := chat.eventService.Run(input)
	c.JSON(http.StatusOK, gin.H{"data": messages})
}
