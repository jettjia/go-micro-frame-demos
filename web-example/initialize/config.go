package initialize

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	"web-gin/global"
	"microframe.com/nacos"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	debug := GetEnvInfo("MXSHOP_DEBUG")
	configFileName := fmt.Sprintf("config-prod.yaml")
	if debug {
		configFileName = fmt.Sprintf("config-dev.yaml")
	}

	v := viper.New()
	//文件的路径如何设置
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	// 读取远程nacos配置信息
	{
		// 把本地设置的nacos配置连接信息解析到全局变量的 NacosConfig
		if err := v.Unmarshal(&global.NacosConfig); err != nil {
			panic(err)
		}

		// 读取远程nacos的配置内容
		content, err := nacos.GetClientContent(global.NacosConfig.Host, global.NacosConfig.Port, global.NacosConfig.Namespace, global.NacosConfig.User, global.NacosConfig.Password,
			global.NacosConfig.DataId, global.NacosConfig.Group)
		if err != nil {
			zap.S().Errorf("读取远程nacos配置信息失败: %s", err.Error())
		}
		err = json.Unmarshal([]byte(content), &global.ServerConfig)
		if err != nil {
			zap.S().Fatalf("解析远程nacos信息失败：%s", err.Error())
		}
	}

	// 监听 nacos 配置变化
	//err = configClient.ListenConfig(vo.ConfigParam{
	//	DataId: global.NacosConfig.DataId,
	//	Group:  global.NacosConfig.Group,
	//	OnChange: func(namespace, group, dataId, data string) {
	//		// todo, 这里能获取到 nacos 变化的值，但是没能赋值到 ServerConfig中
	//		fmt.Println("nacos中的配置", data)
	//		// 这里输出的格式是：  { "name": "user-srv", "host": "10.4.7.71" }
	//		err = json.Unmarshal([]byte(data), &global.ServerConfig)
	//		if err != nil {
	//			zap.S().Errorf("配置中心文件改变后，解析 Json失败")
	//		}
	//		zap.S().Infof("nacos 改变后配置：", global.ServerConfig)
	//	},
	//})
	//if err != nil {
	//	zap.S().Errorf("配置中心文件变化，解析失败!")
	//}

	zap.S().Infof("从nacos读取到的全部配置如下：", global.ServerConfig)
	////////////////////////////////

}
