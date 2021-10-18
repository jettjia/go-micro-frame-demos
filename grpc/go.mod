module go-micro-frame

go 1.16

require (
	github.com/go-redis/redis/v8 v8.11.3
	github.com/go-redsync/redsync/v4 v4.3.0
	github.com/jinzhu/gorm v1.9.16
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/olivere/elastic/v7 v7.0.29
	github.com/opentracing/opentracing-go v1.2.0
	github.com/spf13/viper v1.8.1
	go.uber.org/zap v1.19.1
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gorm.io/driver/mysql v1.1.2
	gorm.io/gorm v1.21.13
	microframe.com/jaeger v0.0.0-00010101000000-000000000000
	microframe.com/logger v0.0.0-00010101000000-000000000000
	microframe.com/mysql v0.0.0-00010101000000-000000000000
	microframe.com/nacos v0.0.0-00010101000000-000000000000
	microframe.com/otgrpc v0.0.0-00010101000000-000000000000
	microframe.com/publicUtil v0.0.0-00010101000000-000000000000
)

replace (
	microframe.com/jaeger => ../microframe.com/jaeger
	microframe.com/logger => ../microframe.com/logger
	microframe.com/mysql => ../microframe.com/database/mysql
	microframe.com/nacos => ../microframe.com/nacos
	microframe.com/otgrpc => ../microframe.com/otgrpc
	microframe.com/publicUtil => ../microframe.com/utils/publicUtil
)
