package initialize

import (
	"microframe.com/jaeger"

	"go-micro-frame/global"
)

func InitJaeger(){
	jaeger.InitJaeger(global.ServerConfig.JaegerInfo.Host, global.ServerConfig.JaegerInfo.Port, global.ServerConfig.JaegerInfo.Name,
		global.ServerConfig.JaegerInfo.User, global.ServerConfig.JaegerInfo.Password)
}