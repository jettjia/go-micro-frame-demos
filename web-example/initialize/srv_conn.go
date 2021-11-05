package initialize

import (
	"github.com/jettjia/go-micro-frame/core/register/nacos"
	mylogger "github.com/jettjia/go-micro-frame/service/logger"

	"web-example/global"
	"web-example/proto"
)

func InitSrvConn() {
	c := nacos.NewRegistryClient(global.NacosConfig.Host, global.NacosConfig.Port, global.NacosConfig.Namespace, global.NacosConfig.User, global.NacosConfig.Password)
	userConn, err := c.Discovery(global.ServerConfig.UserSrvInfo.Name, global.ServerConfig.Env)
	if err != nil {
		mylogger.Fatal("[InitSrvConn] 连接 【用户服务失败】")
	}

	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient
}
