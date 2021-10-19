package global

import (
	ut "github.com/go-playground/universal-translator"

	"web-example/config"
	"web-example/proto"
)

var (
	Trans ut.Translator
	ServerConfig *config.ServerConfig = &config.ServerConfig{}

	UserSrvClient proto.UserClient

	NacosConfig *config.NacosConfig = &config.NacosConfig{}
)
