# 表示依赖 alpine 最新版
FROM golang:1.16 AS build
MAINTAINER jettjia <jettjia@qq.com>

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"

WORKDIR /go/src/goods-srv/

# 复制源码并执行build，此处当文件有变化会产生新的一层镜像层
COPY . /go/src/goods-srv/
## 拷贝配置文件到容器中
COPY config-dev.yaml /go/src/goods-srv/config-dev.yaml
COPY config-prod.yaml /go/src/goods-srv/config-prod.yaml

# 安装依赖包
RUN go mod tidy

RUN go build -o /bin/goods-srv

# 缩小到一层镜像
FROM busybox
COPY --from=build /bin/goods-srv /bin/goods-srv

# 暴露端口
EXPOSE 50051

# 设置时区为上海
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

# 运行golang程序的命令
ENTRYPOINT ["/bin/goods-srv"]