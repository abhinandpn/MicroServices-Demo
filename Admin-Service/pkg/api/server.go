package http

import (
	handlerinterface "github.com/abhinandpn/MicroServices-Demo/Admin-Service/pkg/api/handler/interfaces"
	"github.com/abhinandpn/MicroServices-Demo/Admin-Service/pkg/api/routes"
	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(adminHandler handlerinterface.Adminhandler) *ServerHTTP {
	Engine := gin.New()

	routes.AdminRoute(Engine.Group("/"), adminHandler)

	return &ServerHTTP{engine: Engine}
}

func (sh *ServerHTTP) Run() {
	sh.engine.Run(":5000")
}
