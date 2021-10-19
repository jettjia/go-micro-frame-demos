package global

import (
	"github.com/go-redsync/redsync/v4"
	"github.com/jinzhu/gorm"
	"github.com/olivere/elastic/v7"

	"srv-example/config"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
	NacosConfig  config.NacosConfig
	EsClient     *elastic.Client
	RedsyncLock  *redsync.Redsync
)
