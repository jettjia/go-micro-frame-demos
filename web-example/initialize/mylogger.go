package initialize

import (
	"github.com/jettjia/go-micro-frame/service/logger"

	"web-example/global"
)

// mylogger 自定义logger也启动
func InitMyLogger() {
	// 初始化logger配置
	logger.NewLogger(global.ServerConfig.Name, global.ServerConfig.LoggerInfo.LogFilePath, global.ServerConfig.LoggerInfo.LogLevel, global.ServerConfig.LoggerInfo.MaxSize,
		global.ServerConfig.LoggerInfo.MaxBackups, global.ServerConfig.LoggerInfo.MaxAge)

	logger.Debug("初始化日志")
}
