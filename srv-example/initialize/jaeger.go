package initialize

import (
	"github.com/jettjia/go-micro-frame/core/jaeger"

	"srv-example/global"
)

func InitJaeger(){
	jaeger.InitJaeger(global.ServerConfig.JaegerInfo.Host, global.ServerConfig.JaegerInfo.Port, global.ServerConfig.JaegerInfo.Name,
		global.ServerConfig.JaegerInfo.User, global.ServerConfig.JaegerInfo.Password)
}