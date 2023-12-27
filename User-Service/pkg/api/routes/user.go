package routes

import (
	handlerinterface "github.com/abhinandpn/MicroServices-Demo/User-Service/pkg/api/handler/interfaces"

	"github.com/gin-gonic/gin"
)

func UserRoute(api *gin.RouterGroup,
	userhandler handlerinterface.UserHandler) {

	api.POST("/signup", userhandler.SignUp)
	api.POST("/login", userhandler.Login)
	api.POST("/logout", userhandler.Logout)
	api.GET("/", userhandler.UserHome)
	api.GET("/info", userhandler.UserInfo)
}
