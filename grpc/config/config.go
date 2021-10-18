package config

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

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
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
	Name string   `mapstructure:"name" json:"name"`
	Host string   `mapstructure:"host" json:"host"`
	Port uint64   `mapstructure:"port" json:"port"`
	Env  string   `mapstructure:"env" json:"env"`
	Tags []string `mapstructure:"tags" json:"tags"`

	MysqlInfo   MysqlConfig  `mapstructure:"mysql" json:"mysql"`
	ConsulInfo  ConsulConfig `mapstructure:"consul" json:"consul"`
	RedisConfig ConsulConfig `mapstructure:"redis" json:"redis"`
	JaegerInfo  JaegerConfig `mapstructure:"consul" json:"jaeger"`
	EsInfo      EsConfig     `mapstructure:"es" json:"es"`
	MqInfo      MqConfig     `mapstructure:"mq" json:"mq"`
	LoggerInfo  LoggerConfig `mapstructure:"logger" json:"logger"`
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
