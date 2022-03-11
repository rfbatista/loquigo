package adapters

import (
	adapterservices "loquigo/engine/src/adapters/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewEditorController(e adapterservices.EditorService) EditorController {
	return EditorController{EditorService: e}
}

func (r HttpRouter) AddEditorRoutes(rg *gin.RouterGroup, controller EditorController) {
	route := rg.Group("/editor")

	route.PUT("/", controller.UpdateBot)
	route.GET("/:botId", controller.FindBot)
}

type EditorController struct {
	EditorService adapterservices.EditorService
}

type Input struct {
	Data string `json:"data"`
}

func (e EditorController) UpdateBot(c *gin.Context) {
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, _ := e.EditorService.UpdateBot(input.Data)
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (e EditorController) FindBot(c *gin.Context) {
	botID := c.Param("botId")
	response, _ := e.EditorService.FindBot(botID)
	c.YAML(http.StatusOK, response)
}
