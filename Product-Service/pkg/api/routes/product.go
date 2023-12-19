package routes

import (
	"github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/api/handler/interfaces"
	"github.com/gin-gonic/gin"
)

func ProductRoute(api *gin.RouterGroup,
	ProductHandler interfaces.ProductHandler) {

	// Product CURD
	api.POST("/add", ProductHandler.AddProduct)
	api.GET(":id", ProductHandler.GetProductById)
	api.GET("", ProductHandler.GetProducts)
	api.PATCH("/:id", ProductHandler.UpdateProduct)
	api.DELETE(":id", ProductHandler.DeleteProduct)
}
