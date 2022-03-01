package adapters

import "github.com/gin-gonic/gin"

func NewRouter(g *gin.Engine) HttpRouter {
	return HttpRouter{Router: g}
}

type HttpRouter struct {
	Router *gin.Engine
}
