package adapters

import (
	adapterservices "loquigo/engine/src/adapters/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewFlowMapController(service adapterservices.FlowMapService) FlowMapController {
	return FlowMapController{service: service}
}

type FlowMapController struct {
	service adapterservices.FlowMapService
}

func (r HttpRouter) AddFlowMapRoutes(rg *gin.RouterGroup, controller FlowMapController) {
	route := rg.Group("/flow/map")

	route.GET("/:flow_id", controller.GetFlowMap)
}

func (f FlowMapController) GetFlowMap(c *gin.Context) {
	flowId := c.Param("flow_id")
	flow := f.service.GetMapFromFlow(flowId)
	c.JSON(http.StatusOK, gin.H{"data": flow})
}
