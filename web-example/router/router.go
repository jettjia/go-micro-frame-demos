package router

import (
	"github.com/gin-gonic/gin"
	"web-gin/api/user"
	"web-gin/middlewares"
)

func InitUserRouter(Router *gin.RouterGroup){
	UserRouter := Router.Group("user").Use(middlewares.Trace())
	{
		UserRouter.GET("list", user.GetUserList)
	}
}