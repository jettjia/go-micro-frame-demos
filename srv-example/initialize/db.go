package initialize

import (
	"go.uber.org/zap"

	mySql "github.com/jettjia/go-micro-frame/service/gmysql"

	"srv-example/global"
)

func InitDB() {

	c := global.ServerConfig.MysqlInfo

	m := &mySql.Mysql{
		Host:         c.Host,
		Port:         c.Port,
		User:         c.User,
		Password:     c.Password,
		Db:           c.Name,
		MaxIdleConns: c.MaxIdleConns,
		MaxOpenConns: c.MaxOpenConns,
		MaxLifetime:  c.MaxLifetime,
	}

	var err error
	global.DB, err = m.GetDB()
	if err != nil {
		zap.S().Error("db connect err:", err.Error())
	}
}
