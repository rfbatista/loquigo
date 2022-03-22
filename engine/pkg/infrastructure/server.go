package infrastructure

import (
	adapters "loquigo/engine/pkg/adapters/transport/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewServer(e adapters.EditorController, c adapters.ChatController, b adapters.BotController) Server {
	return Server{EditorController: e, ChatController: c, BotController: b}
}

type Server struct {
	EditorController adapters.EditorController
	ChatController   adapters.ChatController
	BotController    adapters.BotController
}

func (s Server) Start() {
	r := adapters.NewRouter(gin.Default())
	r.Router.Use(cors.Default())
	v1 := r.Router.Group("/v1")
	r.AddEditorRoutes(v1, s.EditorController)
	r.AddChatRoutes(v1, s.ChatController)
	r.AddBotRoutes(v1, s.BotController)
	r.Router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
