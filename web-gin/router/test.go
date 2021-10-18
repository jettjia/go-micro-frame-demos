package router

import (
	"github.com/gin-gonic/gin"
	"web-gin/api/ceshi"
	"web-gin/middlewares"
)

func InitTestRouter(Router *gin.RouterGroup){
	UserRouter := Router.Group("test").Use(middlewares.Trace())
	{
		UserRouter.GET("send-mq", ceshi.SendMq)
	}
}