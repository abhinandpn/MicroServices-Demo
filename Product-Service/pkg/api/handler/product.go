package handler

import (
	"net/http"

	"github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/api/handler/interfaces"
	service "github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/usecase/interfaces"
	"github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/util/req"
	"github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/util/res"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productusecase service.ProductUsecase
}

func NewProductHandler(ProductUserCase service.ProductUsecase) interfaces.ProductHandler {
	return &ProductHandler{productusecase: ProductUserCase}
}

func (p *ProductHandler) AddProduct(ctx *gin.Context) {

	var body req.ReqProduct

	if err := ctx.ShouldBindJSON(&body); err != nil {
		respones := res.ErrorResponse(400, "invalid input", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, respones)
		return
	}

	err := p.productusecase.CreateProduct(body)
	if err != nil {
		response := res.ErrorResponse(400, "faild to add product", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := res.SuccessResponse(200, "successfully product added", body)
	ctx.JSON(http.StatusOK, response)

}

func (p *ProductHandler) GetProductById(ctx *gin.Context) {

}

func (p *ProductHandler) GetProducts(ctx *gin.Context) {

}

func (p *ProductHandler) DeleteProduct(ctx *gin.Context) {

}

func (p *ProductHandler) UpdateProduct(ctx *gin.Context) {

}
