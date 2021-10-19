package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"web-example/util/externalIP"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/jettjia/go-micro-frame/core/register/nacos"
	mylogger "github.com/jettjia/go-micro-frame/service/logger"

	"web-example/global"
	"web-example/initialize"
)

func main() {
	// 初始化web
	initialize.InitWeb()

	// 初始化routers
	Router := initialize.Routers()

	/////////////////////////////////////////////
	// 随机生成 port, 如果是本地开发环境端口号固定，线上环境启动获取端口号
	//if global.ServerConfig.Env != "dev" {
	//	global.ServerConfig.Port, _ = publicUtil.GetFreePort()
	//}

	//注册服务健康检查
	grpcServer := grpc.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())

	//// 如果是生产环境，读取本机的IP
	if global.ServerConfig.Env == "prod" {
		eIP, err := externalIP.ExternalIP()
		if err != nil {
			fmt.Println(err)
		}
		global.ServerConfig.Host = eIP.String()
	}

	//服务注册到nacos
	client := nacos.NewRegistryClient(global.NacosConfig.Host, global.NacosConfig.Port, global.NacosConfig.Namespace, global.NacosConfig.User, global.NacosConfig.Password)

	err := client.Register(global.ServerConfig.Host, uint64(global.ServerConfig.Port), global.ServerConfig.Name, global.ServerConfig.Env,10)
	if err != nil {
		zap.S().Panic("服务注册失败:", err.Error())
	}

	/////////////////////////////////////////////

	// 启动 web服务
	zap.S().Debugf("启动grpc服务IP： %s", global.ServerConfig.Host)
	zap.S().Debugf("启动web服务的端口： %d", global.ServerConfig.Port)
	mylogger.Info("", mylogger.Any("启动web服务端口,Host:", global.ServerConfig.Host))
	mylogger.Info("", mylogger.Any("启动web服务端口,port:", global.ServerConfig.Port))
	go func() {
		if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
			zap.S().Panic("启动失败:", err.Error())
		}
	}()

	//接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = client.DelRegister(global.ServerConfig.Host, uint64(global.ServerConfig.Port), global.ServerConfig.Name, global.ServerConfig.Env); err != nil {
		zap.S().Info("注销失败:", err.Error())
	} else {
		zap.S().Info("注销成功")
	}
}
