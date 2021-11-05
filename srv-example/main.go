package main

import (
	"flag"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/jettjia/go-micro-frame/core/register/nacos"
	"github.com/jettjia/go-micro-frame/core/register/otgrpc"
	mylogger "github.com/jettjia/go-micro-frame/service/logger"
	"github.com/opentracing/opentracing-go"

	"srv-example/domain/repository"
	service2 "srv-example/domain/service"
	"srv-example/global"
	"srv-example/handler"
	"srv-example/initialize"
	"srv-example/proto"
	"srv-example/util/externalIP"
)

func main() {
	// 判断是否生成 随机的 微服务端口号
	IP := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("port", 0, "端口")
	flag.Parse()

	// 初始化
	initialize.InitSrv()

	*Port = int(global.ServerConfig.Port)
	//if global.ServerConfig.Env != "dev" {
	//	*Port, _ = publicUtil.GetFreePort()
	//}

	zap.S().Info("main函数：", global.ServerConfig)
	mylogger.Info("", mylogger.Any("main函数从nacos读取到的配置:", global.ServerConfig))

	// 启动grpc，并且使用 jaeger
	server := grpc.NewServer(grpc.UnaryInterceptor(otgrpc.OpenTracingServerInterceptor(opentracing.GlobalTracer())))

	// 创建实例
	userService := service2.NewUserService(repository.NewUserRepository())

	proto.RegisterUserServer(server, &handler.UserServer{UserService: userService})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
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
	client := nacos.NewRegistryClient(global.NacosConfig.Host, global.NacosConfig.Port, global.NacosConfig.Namespace, global.NacosConfig.User, global.NacosConfig.Password)
	{
		// 如果是生产环境，读取本机的IP
		if global.ServerConfig.Env == "prod" {
			eIP, err := externalIP.ExternalIP()
			if err != nil {
				fmt.Println(err)
			}
			global.ServerConfig.Host = eIP.String()
		}

		err = client.Register(global.ServerConfig.Host, uint64(*Port), global.ServerConfig.Name, global.ServerConfig.Env, 10)
		if err != nil {
			zap.S().Panic("服务注册失败:", err.Error())
		}
	}

	zap.S().Debugf("启动grpc服务IP： %s", global.ServerConfig.Host)
	zap.S().Debugf("启动grpc服务端口： %d", *Port)
	mylogger.Info("", mylogger.Any("启动grpc服务IP:", global.ServerConfig.Host))
	mylogger.Info("", mylogger.Any("启动grpc服务端口,port:", *Port))

	//接收终止信号
	{
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		if err = client.DelRegister(global.ServerConfig.Host, uint64(*Port), global.ServerConfig.Name, global.ServerConfig.Env); err != nil {
			zap.S().Info("注销失败:", err.Error())
		} else {
			zap.S().Info("注销成功")
		}
	}
}
