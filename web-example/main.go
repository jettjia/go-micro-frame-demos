package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/jettjia/go-micro-frame/core/register/register"
	mylogger "github.com/jettjia/go-micro-frame/service/logger"
	microExtertalIp "github.com/jettjia/go-micro-frame/tool/external_ip"
	microFreePort "github.com/jettjia/go-micro-frame/tool/free_port"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"web-example/global"
	"web-example/initialize"
)

func main() {
	IP := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("port", 0, "端口")
	flag.Parse()

	// 初始化web
	initialize.InitWeb()

	// 初始化routers
	Router := initialize.Routers()

	/////////////////////////////////////////////
	//注册服务健康检查
	grpcServer := grpc.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())

	//// 如果是生产环境，读取本机的IP
	// 如果是生产环境，读取本机的IP
	if global.ServerConfig.Env == "prod" {
		eIp, err := microExtertalIp.ExternalIP()
		if err != nil {
			mylogger.Error(err, mylogger.Any("prod环境：", "获取IP错误"))
		}
		global.ServerConfig.RegisterInfo.ServiceHost = eIp.String()

		ePort, err := microFreePort.GetFreePort()
		if err != nil {
			mylogger.Error(err, mylogger.Any("prod环境：", "获取Port错误"))
		}
		global.ServerConfig.RegisterInfo.ServicePort = ePort
	} else if global.ServerConfig.Env == "debug" {
		global.ServerConfig.RegisterInfo.ServiceHost = *IP
		global.ServerConfig.RegisterInfo.ServicePort = *Port
	}

	//服务注册到nacos
	reg := register.Reg{}
	copier.Copy(&reg, &global.ServerConfig.RegisterInfo)
	client := register.NewRegClient(reg)

	err := client.Register()
	if err != nil {
		zap.S().Panic("服务注册失败:", err.Error())
	}

	/////////////////////////////////////////////

	// 启动 web服务
	mylogger.Info("", mylogger.Any("main函数从nacos读取到的配置:", global.ServerConfig))
	mylogger.Info("", mylogger.Any("启动web服务Host:", global.ServerConfig.RegisterInfo.ServiceHost))
	mylogger.Info("", mylogger.Any("启动web服务Port:", global.ServerConfig.RegisterInfo.ServicePort))
	go func() {
		if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.RegisterInfo.ServicePort)); err != nil {
			zap.S().Panic("启动失败:", err.Error())
		}
	}()

	//接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = client.DelRegister(); err != nil {
		mylogger.Info("", mylogger.Any("注销失败:", err.Error()))
	} else {
		mylogger.Info("", mylogger.Any("注销成功:", "success"))
	}
}
