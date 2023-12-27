package handler

import (
	"errors"
	"net/http"

	handlerinterface "github.com/abhinandpn/MicroServices-Demo/User-Service/pkg/api/handler/interfaces"
	"github.com/abhinandpn/MicroServices-Demo/User-Service/pkg/helper"
	"github.com/abhinandpn/MicroServices-Demo/User-Service/pkg/model"
	services "github.com/abhinandpn/MicroServices-Demo/User-Service/pkg/usecase/interfaces"
	req "github.com/abhinandpn/MicroServices-Demo/User-Service/pkg/util/Req"
	"github.com/abhinandpn/MicroServices-Demo/User-Service/pkg/util/res"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type UserHandler struct {
	UserUsecase services.UserUsecase
}

func NewUserUseHandler(usecase services.UserUsecase) handlerinterface.UserHandler {
	return &UserHandler{UserUsecase: usecase}
}

func (u *UserHandler) UserHome(ctx *gin.Context) {

	UserId := helper.GetUserId(ctx)
	_, err := u.UserUsecase.GetUserInfo(ctx, UserId)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"StatusCode": 200,
		"message":    "Welcome to User Home ",
	})
}

func (u *UserHandler) Logout(ctx *gin.Context) {
	ctx.SetCookie("user-auth", "", -1, "", "", false, true)
	response := res.SuccessResponse(200, "successfully logged out", nil)
	ctx.JSON(http.StatusOK, response)
}

func (u *UserHandler) UserInfo(ctx *gin.Context) {
	UserId := helper.GetUserId(ctx)
	UserInfo, err := u.UserUsecase.GetUserInfo(ctx, UserId)
	if err != nil {
		response := res.ErrorResponse(500, "faild to show user details", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	response := res.SuccessResponse(200, "Successfully user account details found", UserInfo)
	ctx.JSON(http.StatusOK, response)

}

func (u *UserHandler) Login(ctx *gin.Context) {

	var body req.UserLoginReq

	err := ctx.ShouldBindJSON(&body)

	if err != nil {

		response := res.ErrorResponse(400, "invalid input", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if body.Username == "" {
		err := errors.New("enter username")
		response := res.ErrorResponse(400, "invalid input", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var user model.User
	copier.Copy(&user, &body)
	user, err = u.UserUsecase.Login(ctx, user)
	if err != nil {
		response := res.ErrorResponse(400, "faild to login", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// JWT jenerate
}

func (u *UserHandler) SignUp(ctx *gin.Context) {

	var body req.UserSignUpReq
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		response := res.ErrorResponse(400, "invalid input", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var user model.User
	copier.Copy(&user, body)
	err = u.UserUsecase.SignUp(ctx, user)

	if err != nil {
		response := res.ErrorResponse(400, "faild to signup", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := res.SuccessResponse(200, "Successfully Created Account", body)
	ctx.JSON(200, response)
}
