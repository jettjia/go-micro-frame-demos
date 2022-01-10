package config

// 服务注册和发现
type RegisterConfig struct {
	Typ         string   `mapstructure:"typ" json:"typ"`                 // 类型：nacos, consul
	Host        string   `mapstructure:"host" json:"host"`               // nacos或consul地址
	Port        int      `mapstructure:"port" json:"port"`               // nacos或consul port
	User        string   `mapstructure:"user" json:"user"`               // nacos的user
	Password    string   `mapstructure:"password" json:"password"`       // nacos的password
	ServiceHost string   `mapstructure:"serviceHost" json:"serviceHost"` // 服务的host
	ServicePort int      `mapstructure:"servicePort" json:"servicePort"` // 服务的port
	ServiceName string   `mapstructure:"serviceName" json:"serviceName"` // 服务的名称
	GroupName   string   `mapstructure:"groupName" json:"groupName"`     // nacos的group
	Weight      float64  `mapstructure:"weight" json:"weight"`           // nacos的weight
	Tags        []string `mapstructure:"tags" json:"tags"`               // consul的tags
}

// User服务
type UserSrvConfig struct {
	Typ         string   `mapstructure:"typ" json:"typ"`                 // 类型：nacos, consul
	Host        string   `mapstructure:"host" json:"host"`               // nacos或consul地址
	Port        int      `mapstructure:"port" json:"port"`               // nacos或consul port
	User        string   `mapstructure:"user" json:"user"`               // nacos的user
	Password    string   `mapstructure:"password" json:"password"`       // nacos的password
	ServiceHost string   `mapstructure:"serviceHost" json:"serviceHost"` // 服务的host
	ServicePort int      `mapstructure:"servicePort" json:"servicePort"` // 服务的port
	ServiceName string   `mapstructure:"serviceName" json:"serviceName"` // 服务的名称
	GroupName   string   `mapstructure:"groupName" json:"groupName"`     // nacos的group
	Weight      float64  `mapstructure:"weight" json:"weight"`           // nacos的weight
	Tags        []string `mapstructure:"tags" json:"tags"`               // consul的tags
}

type RedisConfig struct {
	Host   string `mapstructure:"host" json:"host"`
	Port   int    `mapstructure:"port" json:"port"`
	Expire int    `mapstructure:"expire" json:"expire"`
}

type JaegerConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type LoggerConfig struct {
	LogFilePath string `mapstructure:"logFilePath" json:"logFilePath"`
	LogLevel    string `mapstructure:"logLevel" json:"logLevel"`
	MaxSize     int    `mapstructure:"maxSize" json:"maxSize"`
	MaxBackups  int    `mapstructure:"maxBackups" json:"maxBackups"`
	MaxAge      int    `mapstructure:"maxAge" json:"maxAge"`
}

type ServerConfig struct {
	Env          string         `mapstructure:"env" json:"env"`
	RegisterInfo RegisterConfig `mapstructure:"register" json:"register"`
	UserSrvInfo  UserSrvConfig  `mapstructure:"user_srv" json:"user_srv"`
	RedisInfo    RedisConfig    `mapstructure:"redis" json:"redis"`
	JaegerInfo   JaegerConfig   `mapstructure:"consul" json:"jaeger"`
	LoggerInfo   LoggerConfig   `mapstructure:"logger" json:"logger"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host" `
	Port      uint64 `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}
