package router

import (
	"github.com/gin-gonic/gin"

	"web-example/api/user"
	"web-example/middlewares"
)

func InitUserRouter(Router *gin.RouterGroup){
	UserRouter := Router.Group("user").Use(middlewares.Trace())
	{
		UserRouter.GET("list", user.GetUserList)
	}
}