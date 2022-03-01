package adapters

import (
	"loquigo/engine/src/core/modules/templatepool"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewFlowController(s templatepool.FlowService) FlowController {
	return FlowController{flowService: s}
}

type FlowController struct {
	flowService templatepool.FlowService
}

func (r HttpRouter) AddFlowRoutes(rg *gin.RouterGroup, controller FlowController) {
	route := rg.Group("/flow")

	route.POST("/", controller.CreateFlow)
	route.PUT("/", controller.UpdateFlow)
	route.DELETE("/", controller.DeleteFlow)
	route.GET("/:bot_id", controller.FindByBotId)
}

func (f FlowController) CreateFlow(c *gin.Context) {
	var input templatepool.Flow
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	flow, _ := f.flowService.NewFlow(input)
	c.JSON(http.StatusOK, gin.H{"data": flow})
}

func (f FlowController) UpdateFlow(c *gin.Context) {
	var input templatepool.Flow
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	flow, _ := f.flowService.Update(input)
	c.JSON(http.StatusOK, gin.H{"data": flow})
}

func (f FlowController) DeleteFlow(c *gin.Context) {
	var input templatepool.Flow
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	flow, _ := f.flowService.Delete(input)
	c.JSON(http.StatusOK, gin.H{"data": flow})
}

func (f FlowController) FindByBotId(c *gin.Context) {
	botId := c.Param("bot_id")
	flow, _ := f.flowService.FindByBotId(botId)
	c.JSON(http.StatusOK, gin.H{"data": flow})
}
