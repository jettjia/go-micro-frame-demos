package initialize

func InitSrv() {
	// 初始化 logger
	InitLogger()

	//初始化配置文件
	InitConfig()

	// 初始化db
	InitDB()

	// 初始化es
	//InitEs()

	// 初始化jaeger
	InitJaeger()

	// 初始 mylogger
	InitMyLogger()
}