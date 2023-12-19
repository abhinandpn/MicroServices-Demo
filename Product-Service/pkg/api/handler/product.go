package handler

import (
	"net/http"

	"github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/api/handler/interfaces"
	"github.com/abhinandpn/MicroServices-Demo/Product-Service/pkg/helper"
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

	var body req.ReqProduct
	if err := ctx.ShouldBindJSON(&body); err != nil {
		respones := res.ErrorResponse(400, "invalid input", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, respones)
		return
	}

	ParamId := ctx.Param("id")
	id, err := helper.StringToUInt(ParamId)
	if err != nil {
		respones := res.ErrorResponse(400, "failed to convert param id", err.Error(), id)
		ctx.JSON(http.StatusBadRequest, respones)
		return
	}

	product, err := p.productusecase.GetProductById(id)
	if err != nil {
		respones := res.ErrorResponse(400, "failed to get product", err.Error(), id)
		ctx.JSON(http.StatusBadRequest, respones)
		return
	}

	response := res.SuccessResponse(200, "successfully product added", product)
	ctx.JSON(http.StatusOK, response)

}

func (p *ProductHandler) GetProducts(ctx *gin.Context) {

	products, err := p.productusecase.GetProducts()
	if err != nil {
		respones := res.ErrorResponse(400, "failed to get product", err.Error(), products)
		ctx.JSON(http.StatusBadRequest, respones)
		return
	}

	response := res.SuccessResponse(200, "successfully get producrs", products)
	ctx.JSON(http.StatusOK, response)

}

func (p *ProductHandler) DeleteProduct(ctx *gin.Context) {

	ParamId := ctx.Param("id")
	id, err := helper.StringToUInt(ParamId)
	if err != nil {
		respones := res.ErrorResponse(400, "failed to convert param id", err.Error(), id)
		ctx.JSON(http.StatusBadRequest, respones)
		return
	}

	err = p.productusecase.DeleteProduct(id)
	if err != nil {
		respones := res.ErrorResponse(400, "failed to get product", err.Error(), id)
		ctx.JSON(http.StatusBadRequest, respones)
		return
	}

	response := res.SuccessResponse(200, "successfully get producrs", id)
	ctx.JSON(http.StatusOK, response)

}

func (p *ProductHandler) UpdateProduct(ctx *gin.Context) {

	var body req.ReqProduct
	if err := ctx.ShouldBindJSON(&body); err != nil {
		respones := res.ErrorResponse(400, "invalid input", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, respones)
		return
	}

	ParamId := ctx.Param("id")
	id, err := helper.StringToUInt(ParamId)
	if err != nil {
		respones := res.ErrorResponse(400, "failed to convert param id", err.Error(), id)
		ctx.JSON(http.StatusBadRequest, respones)
		return
	}

	err = p.productusecase.UpdateProduct(body, id)
	if err != nil {
		respones := res.ErrorResponse(400, "failed to update product", err.Error(), id)
		ctx.JSON(http.StatusBadRequest, respones)
		return
	}

	response := res.SuccessResponse(200, "successfully update producrs", id)
	ctx.JSON(http.StatusOK, response)
}
