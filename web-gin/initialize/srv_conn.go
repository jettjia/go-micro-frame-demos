package initialize

import (
	"go.uber.org/zap"

	"microframe.com/nacos"
	"web-gin/global"
	"web-gin/proto"
)

func InitSrvConn() {
	nc := nacos.NacosClient{
		Host:      global.NacosConfig.Host,
		Port:      global.NacosConfig.Port,
		Namespace: global.NacosConfig.Namespace,
		User:      global.NacosConfig.User,
		Password:  global.NacosConfig.Password,
	}
	userConn, err := nc.Discovery(nc, global.ServerConfig.UserSrvInfo.Name, global.ServerConfig.Env)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 【用户服务失败】")
	}

	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient
}
