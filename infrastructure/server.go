package infrastructure

import (
	"loquigo/engine/pkg/adapters"
	"loquigo/engine/pkg/adapters/bot"
	"loquigo/engine/pkg/adapters/chat"
	"loquigo/engine/pkg/adapters/editor"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewServer(e editor.EditorController, c chat.ChatController, b bot.BotController, config ServerConfig) Server {
	return Server{EditorController: e, ChatController: c, BotController: b, trustedSources: config.TrustedSources}
}

type Server struct {
	EditorController editor.EditorController
	ChatController   chat.ChatController
	BotController    bot.BotController
	trustedSources   []string
}

func (s Server) Start() {
	r := adapters.NewRouter(gin.Default())
	r.Router.Use(cors.New(cors.Config{
		AllowOrigins:  s.trustedSources,
		AllowMethods:  []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:  []string{"Origin", "Content-Type"},
		ExposeHeaders: []string{"Content-Length", "Content-Type"},
	}))
	v1 := r.Router.Group("/v1")
	s.EditorController.AddEditorRoutes(v1)
	s.ChatController.AddChatRoutes(v1)
	s.BotController.AddBotRoutes(v1)
	r.Router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
