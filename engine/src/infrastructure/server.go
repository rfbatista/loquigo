package infrastructure

import (
	adapters "loquigo/engine/src/adapters/transport/http"
	controller "loquigo/engine/src/adapters/transport/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewServer(
	// c controller.ChatController,
	f controller.FlowController,
	s controller.StepController,
	co controller.ComponentController,
	fo controller.FlowMapController) Server {
	return Server{
		// ChatController:      c,
		FlowController:      f,
		StepController:      s,
		ComponentController: co,
		FlowMapController:   fo}
}

type Server struct {
	// ChatController      controller.ChatController
	FlowController      controller.FlowController
	StepController      controller.StepController
	ComponentController controller.ComponentController
	FlowMapController   controller.FlowMapController
}

func (s Server) Start() {
	r := adapters.NewRouter(gin.Default())
	r.Router.Use(gin.Recovery())
	r.Router.Use(cors.Default())
	v1 := r.Router.Group("/v1")
	r.AddFlowRoutes(v1, s.FlowController)
	r.AddStepRoutes(v1, s.StepController)
	r.AddComponentRoutes(v1, s.ComponentController)
	// r.AddChatRoutes(v1, s.ChatController)
	r.AddFlowMapRoutes(v1, s.FlowMapController)
	r.Router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
