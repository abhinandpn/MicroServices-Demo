package interfaces

import "github.com/gin-gonic/gin"

type UserHandler interface {
	UserHome(ctx *gin.Context)
	UserInfo(ctx *gin.Context)
	Logout(ctx *gin.Context)
	Login(ctx *gin.Context)
	SignUp(ctx *gin.Context)
}
