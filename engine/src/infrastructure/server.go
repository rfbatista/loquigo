package infrastructure

import (
	adapters "loquigo/engine/src/adapters/transport/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewServer(e adapters.EditorController, c adapters.ChatController) Server {
	return Server{EditorController: e, ChatController: c}
}

type Server struct {
	EditorController adapters.EditorController
	ChatController   adapters.ChatController
}

func (s Server) Start() {
	r := adapters.NewRouter(gin.Default())
	r.Router.Use(gin.Recovery())
	r.Router.Use(cors.Default())
	v1 := r.Router.Group("/v1")
	r.AddEditorRoutes(v1, s.EditorController)
	r.AddChatRoutes(v1, s.ChatController)
	r.Router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
