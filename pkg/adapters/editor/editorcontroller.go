package editor

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewEditorController(e EditorService) EditorController {
	return EditorController{service: e}
}

type EditorController struct {
	service EditorService
}

type Input struct {
	Data string `json:"data"`
}

func (r EditorController) AddEditorRoutes(rg *gin.RouterGroup) {
	route := rg.Group("/editor")

	route.POST("/:botId/version/:versionId", r.UpdateBot)
	route.GET("/:botId", r.FindBot)
	route.GET("/:botId/version", r.FindVersions)
	route.GET("/:botId/version/:versionId", r.FindVersionByIdAndBotID)
}

func (e EditorController) UpdateBot(c *gin.Context) {
	botID := c.Param("botId")
	versionID := c.Param("versionId")
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := e.service.UpdateBot(input.Data, versionID, botID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": response})
	}
}

func (e EditorController) FindBot(c *gin.Context) {
	botID := c.Param("botId")
	response, _ := e.service.FindBot(botID)
	c.YAML(http.StatusOK, response)
}

func (e EditorController) FindVersions(c *gin.Context) {
	botID := c.Param("botId")
	response, _ := e.service.FindBotVersions(botID)
	c.YAML(http.StatusOK, response)
}

func (e EditorController) FindVersionByIdAndBotID(c *gin.Context) {
	botID := c.Param("botId")
	versionID := c.Param("versionId")
	response, _ := e.service.FindVersionByIdAndBotId(versionID, botID)
	c.YAML(http.StatusOK, response)
}
