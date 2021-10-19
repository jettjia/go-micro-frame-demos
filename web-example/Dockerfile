# 表示依赖 alpine 最新版
FROM alpine:latest
MAINTAINER jettjia <jettjia@qq.com>
ENV VERSION 1.0

# 在容器根目录 创建一个 apps 目录
WORKDIR /apps

# 拷贝当前执行文件到项目中
COPY web-example/web-example /apps/

# 拷贝配置文件到容器中
COPY web-example/config-dev.yaml /apps/config-dev.yaml
COPY web-example/config-prod.yaml /apps/config-prod.yaml

# 设置时区为上海
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

# 设置编码
ENV LANG C.UTF-8

# 暴露端口
EXPOSE 8021

# 运行golang程序的命令
ENTRYPOINT ["/apps/web-example"]
