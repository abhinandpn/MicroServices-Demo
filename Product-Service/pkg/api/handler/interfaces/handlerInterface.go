package interfaces

import "github.com/gin-gonic/gin"

type ProductHandler interface {
	AddProduct(ctx *gin.Context)
	GetProductById(ctx *gin.Context)
	GetProducts(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
}
