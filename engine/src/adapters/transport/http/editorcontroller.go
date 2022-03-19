package adapters

import (
	editorservice "loquigo/engine/src/adapters/services/editor"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewEditorController(e editorservice.EditorService) EditorController {
	return EditorController{EditorService: e}
}

func (r HttpRouter) AddEditorRoutes(rg *gin.RouterGroup, controller EditorController) {
	route := rg.Group("/editor")

	route.PUT("/", controller.UpdateBot)
	route.GET("/:botId", controller.FindBot)
	route.GET("/:botId/version", controller.FindVersions)
	route.GET("/:botId/version/:versionId", controller.FindVersionByIdAndBotID)
}

type EditorController struct {
	EditorService editorservice.EditorService
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
	response, err := e.EditorService.UpdateBot(input.Data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": response})
	}
}

func (e EditorController) FindBot(c *gin.Context) {
	botID := c.Param("botId")
	response, _ := e.EditorService.FindBot(botID)
	c.YAML(http.StatusOK, response)
}

func (e EditorController) FindVersions(c *gin.Context) {
	botID := c.Param("botId")
	response, _ := e.EditorService.FindBotVersions(botID)
	c.YAML(http.StatusOK, response)
}

func (e EditorController) FindVersionByIdAndBotID(c *gin.Context) {
	botID := c.Param("botId")
	versionID := c.Param("versionId")
	response, _ := e.EditorService.FindVersionByIdAndBotId(versionID, botID)
	c.YAML(http.StatusOK, response)
}
