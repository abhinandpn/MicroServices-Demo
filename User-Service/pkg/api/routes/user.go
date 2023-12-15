package routes

import (
	handlerinterface "github.com/abhinandpn/MicroServices-Demo/User-Service/pkg/api/handler/interfaces"

	"github.com/gin-gonic/gin"
)

func UserRoute(api *gin.RouterGroup,
userhandler handlerinterface.UserHandler ){
	
}