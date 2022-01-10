package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/jettjia/go-micro-frame/core/register/otgrpc"
	"github.com/jettjia/go-micro-frame/core/register/register"
	mylogger "github.com/jettjia/go-micro-frame/service/logger"
	microExtertalIp "github.com/jettjia/go-micro-frame/tool/external_ip"
	microFreePort "github.com/jettjia/go-micro-frame/tool/free_port"
	"github.com/jinzhu/copier"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"srv-example/domain/repository"
	service2 "srv-example/domain/service"
	"srv-example/global"
	"srv-example/handler"
	"srv-example/initialize"
	"srv-example/proto"
)

func main() {
	// 判断是否生成 随机的 微服务端口号
	IP := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("port", 0, "端口")
	flag.Parse()

	// 初始化
	initialize.InitSrv()

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

	// 启动grpc，并且使用 jaeger
	server := grpc.NewServer(grpc.UnaryInterceptor(otgrpc.OpenTracingServerInterceptor(opentracing.GlobalTracer())))

	// 创建实例
	userService := service2.NewUserService(repository.NewUserRepository())

	proto.RegisterUserServer(server, &handler.UserServer{UserService: userService})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", global.ServerConfig.RegisterInfo.ServiceHost, global.ServerConfig.RegisterInfo.ServicePort))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	//注册服务健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	// 启动grpc服务
	go func() {
		err = server.Serve(lis)
		if err != nil {
			panic("failed to start grpc:" + err.Error())
		}
	}()

	// 注册服务到 注册中心
	reg := register.Reg{}
	copier.Copy(&reg, &global.ServerConfig.RegisterInfo)
	client := register.NewRegClient(reg)
	{
		err = client.Register()
		if err != nil {
			zap.S().Panic("服务注册失败:", err.Error())
		}
	}

	mylogger.Info("", mylogger.Any("main函数从nacos读取到的配置:", global.ServerConfig))
	mylogger.Info("", mylogger.Any("启动grpc服务Host:", global.ServerConfig.RegisterInfo.ServiceHost))
	mylogger.Info("", mylogger.Any("启动grpc服务Port:", global.ServerConfig.RegisterInfo.ServicePort))

	//接收终止信号
	{
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		if err = client.DelRegister(); err != nil {
			mylogger.Info("", mylogger.Any("注销失败:", err.Error()))
		} else {
			mylogger.Info("", mylogger.Any("注销成功:", "success"))
		}
	}
}
