package initialize

func InitWeb() {
	// 初始化logger
	InitLogger()

	// 初始化配置文件
	InitConfig()

	// 初始化srv的连接
	InitSrvConn()

	// 初始 mylogger
	InitMyLogger()
}
