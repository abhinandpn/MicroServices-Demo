package http

import (
	handlerinterface "github.com/abhinandpn/MicroServices-Demo/User-Service/pkg/api/handler/interfaces"
	"github.com/abhinandpn/MicroServices-Demo/User-Service/pkg/api/routes"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)


type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(
userhandler handlerinterface.UserHandler,
)*ServerHTTP{
	Engine := gin.New()
	Engine.Use(gin.Logger())

	// For Swagger
	Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//route
	routes.UserRoute(Engine.Group("/user"),userhandler)

	return &ServerHTTP{engine: Engine}
	
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":5010")
}