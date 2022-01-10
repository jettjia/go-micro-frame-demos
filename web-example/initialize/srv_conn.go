package initialize

import (
	"github.com/jettjia/go-micro-frame/core/register/register"
	mylogger "github.com/jettjia/go-micro-frame/service/logger"
	"github.com/jinzhu/copier"

	"web-example/global"
	"web-example/proto"
)

func InitSrvConn() {
	reg := register.Reg{}
	copier.Copy(&reg, &global.ServerConfig.UserSrvInfo)

	c := register.NewRegClient(reg)
	userConn, err := c.Discovery()
	if err != nil {
		mylogger.Fatal("[InitSrvConn] 连接 【用户服务失败】")
	}

	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient
}
