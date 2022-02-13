package infrastructure

import (
	controller "loquigo/engine/pkg/adapters/transport/http"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	ChatController controller.ChatController
}

func NewServer(c controller.ChatController) Server {
	return Server{ChatController: c}
}

func (s Server) Start() {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	v1 := r.Group("/v1")
	{
		v1.POST("/chat", s.ChatController.PostMessage)
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
