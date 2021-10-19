package config

type UserSrvConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type RedisConfig struct {
	Host   string `mapstructure:"host" json:"host"`
	Port   int    `mapstructure:"port" json:"port"`
	Expire int    `mapstructure:"expire" json:"expire"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
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
	Name string   `mapstructure:"name" json:"name"`
	Host string   `mapstructure:"host" json:"host"`
	Tags []string `mapstructure:"tags" json:"tags"`
	Port int      `mapstructure:"port" json:"port"`
	Env  string   `mapstructure:"env" json:"env"`

	UserSrvInfo UserSrvConfig `mapstructure:"user_srv" json:"user_srv"`
	RedisInfo   RedisConfig   `mapstructure:"redis" json:"redis"`
	ConsulInfo  ConsulConfig  `mapstructure:"consul" json:"consul"`
	JaegerInfo  JaegerConfig  `mapstructure:"consul" json:"jaeger"`
	LoggerInfo  LoggerConfig  `mapstructure:"logger" json:"logger"`
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
