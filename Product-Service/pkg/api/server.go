package http

import (
	"github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/api/handler/interfaces"
	"github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/api/routes"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServeHTTP struct {
	engine *gin.Engine
}

func NewServeHTTP(ProductHandler interfaces.ProductHandler) *ServeHTTP {

	Engine := gin.New()
	Engine.Use(gin.Logger())

	// For Swagger
	Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// for route
	routes.ProductRoute(Engine.Group("/product"), ProductHandler)
	return &ServeHTTP{engine: Engine}
}
func (sh *ServeHTTP) Start() {
	sh.engine.Run(":5005")
}
