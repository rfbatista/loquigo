package adapters

import (
	adapterservices "loquigo/engine/src/adapters/services"
	"loquigo/engine/src/core/modules/template/pool"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewStepController(service pool.StepService) StepController {
	return StepController{service: service}
}

type StepController struct {
	service pool.StepService
}

func (r HttpRouter) AddStepRoutes(rg *gin.RouterGroup, controller StepController) {
	route := rg.Group("/step")

	route.POST("/", controller.Create)
	route.PUT("/:step_id", controller.Update)
	route.GET("/:step_id", controller.FindById)
	route.DELETE("/", controller.Delete)
	route.GET("/flow/:flow_id", controller.FindByFlowId)
}

func (s StepController) Create(c *gin.Context) {
	var input pool.Step
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	flow, _ := s.service.NewStep(input)
	c.JSON(http.StatusOK, gin.H{"data": flow})
}

func (s StepController) Update(c *gin.Context) {
	var input adapterservices.Node
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	step := input.ToDomain()
	flow, _ := s.service.UpdateStep(step)
	c.JSON(http.StatusOK, gin.H{"data": flow})
}

func (s StepController) Delete(c *gin.Context) {
	var input pool.Step
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	flow, _ := s.service.DeleteStep(input)
	c.JSON(http.StatusOK, gin.H{"data": flow})
}

func (s StepController) FindByFlowId(c *gin.Context) {
	botId := c.Param("flow_id")
	flow, _ := s.service.FindByFlowId(botId)
	c.JSON(http.StatusOK, gin.H{"data": flow})
}

func (s StepController) FindById(c *gin.Context) {
	stepId := c.Param("step_id")
	step, _ := s.service.FindById(stepId)
	c.JSON(http.StatusOK, gin.H{"data": step})
}
