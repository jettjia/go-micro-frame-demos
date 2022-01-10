package config

// nacos配置
type NacosConfig struct {
	Host      string `mapstructure:"host" `
	Port      uint64 `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}

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

type MysqlConfig struct {
	Host         string `mapstructure:"host" json:"host"`
	Port         int    `mapstructure:"port" json:"port"`
	Name         string `mapstructure:"db" json:"db"`
	User         string `mapstructure:"user" json:"user"`
	Password     string `mapstructure:"password" json:"password"`
	MaxIdleConns int    `maxIdleConns:"password" json:"maxIdleConns"`
	MaxOpenConns int    `maxIdleConns:"maxOpenConns" json:"maxOpenConns"`
	MaxLifetime  int    `maxIdleConns:"maxLifetime" json:"maxLifetime"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Password string `mapstructure:"password" json:"password"`
}

type JaegerConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"name" json:"name"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

type EsConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

type MqConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

type LoggerConfig struct {
	LogFilePath string `mapstructure:"logFilePath" json:"logFilePath"`
	LogLevel    string `mapstructure:"logLevel" json:"logLevel"`
	MaxSize     int    `mapstructure:"maxSize" json:"maxSize"`
	MaxBackups  int    `mapstructure:"maxBackups" json:"maxBackups"`
	MaxAge      int    `mapstructure:"maxAge" json:"maxAge"`
}

type ServerConfig struct {
	Env          string         `mapstructure:"env" json:"env"` // debug:本地赋值ip、port; prod:读取机器的ip、port; 其他环境，读取配置信息
	RegisterInfo RegisterConfig `mapstructure:"register" json:"register"`
	MysqlInfo    MysqlConfig    `mapstructure:"mysql" json:"mysql"`
	RedisConfig  RedisConfig    `mapstructure:"redis" json:"redis"`
	JaegerInfo   JaegerConfig   `mapstructure:"consul" json:"jaeger"`
	EsInfo       EsConfig       `mapstructure:"es" json:"es"`
	MqInfo       MqConfig       `mapstructure:"mq" json:"mq"`
	LoggerInfo   LoggerConfig   `mapstructure:"logger" json:"logger"`
}
