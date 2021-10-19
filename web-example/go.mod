module web-gin

go 1.16

require (
	github.com/gin-gonic/gin v1.7.4
	github.com/go-playground/universal-translator v0.18.0
	github.com/go-playground/validator/v10 v10.9.0
	github.com/mbobakov/grpc-consul-resolver v1.4.4
	github.com/opentracing/opentracing-go v1.2.0
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/viper v1.8.1
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	github.com/ugorji/go v1.1.13 // indirect
	go.uber.org/zap v1.19.1
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
	microframe.com/logger v0.0.0-00010101000000-000000000000
	microframe.com/nacos v0.0.0-00010101000000-000000000000
	microframe.com/publicUtil v0.0.0-00010101000000-000000000000
)

replace (
	microframe.com/jaeger => ../microframe.com/jaeger
	microframe.com/logger => ../microframe.com/logger
	microframe.com/nacos => ../microframe.com/nacos
	microframe.com/otgrpc => ../microframe.com/otgrpc
	microframe.com/publicUtil => ../microframe.com/utils/publicUtil
)
