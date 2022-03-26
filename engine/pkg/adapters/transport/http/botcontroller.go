package adapters

import (
	"loquigo/engine/pkg/core/domain"
	"loquigo/engine/pkg/core/services/bot"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewBotController(b bot.BotService) BotController {
	return BotController{service: b}
}

func (r HttpRouter) AddBotRoutes(rg *gin.RouterGroup, controller BotController) {
	route := rg.Group("/bot")

	route.GET("/", controller.GetBots)
	route.POST("/", controller.CreateBot)
}

type BotController struct {
	service bot.BotService
}

func (b BotController) GetBots(c *gin.Context) {
	bots, err := b.service.GetBots()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": bots})
	}
}

func (b BotController) CreateBot(c *gin.Context) {
	var input domain.Bot
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bot, err := b.service.CreateBot(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": bot})
	}
	return
}
