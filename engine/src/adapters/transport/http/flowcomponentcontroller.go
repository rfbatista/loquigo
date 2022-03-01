package adapters

import (
	"loquigo/engine/src/core/modules/templatepool"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewComponentController(service templatepool.ComponentService) ComponentController {
	return ComponentController{service: service}
}

type ComponentController struct {
	service templatepool.ComponentService
}

func (r HttpRouter) AddComponentRoutes(rg *gin.RouterGroup, controller ComponentController) {
	route := rg.Group("/component")

	route.POST("/", controller.Create)
	route.PUT("/", controller.Update)
	route.DELETE("/", controller.Delete)
	route.GET("/flow/:flow_id/step/:step_id", controller.GetComponents)
}

func (co ComponentController) Create(c *gin.Context) {
	var input templatepool.Component
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	flow, _ := co.service.NewComponent(input)
	c.JSON(http.StatusOK, gin.H{"data": flow})
}

func (co ComponentController) Update(c *gin.Context) {
	var input templatepool.Component
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	flow, _ := co.service.UpdateComponent(input)
	c.JSON(http.StatusOK, gin.H{"data": flow})
}
func (co ComponentController) Delete(c *gin.Context) {
	var input templatepool.Component
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	flow, _ := co.service.DeleteComponent(input)
	c.JSON(http.StatusOK, gin.H{"data": flow})
}
func (co ComponentController) GetComponents(c *gin.Context) {
	flowID := c.Param("flow_id")
	stepID := c.Param("step_id")
	flow, _ := co.service.GetComponents(flowID, stepID)
	c.JSON(http.StatusOK, gin.H{"data": flow})
}
