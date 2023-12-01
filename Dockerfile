# 分阶段构建
# 第一阶段：使用Go编译源码，生成二进制文件
# 第二阶段：使用alpine作为基础镜像，将二进制文件复制到镜像中，运行（这个环境只需要能运行二进制文件即可）

# 选择构建用基础镜像（选择原则：在包含所有用到的依赖前提下尽可能提及小）。如需更换，请到[dockerhub官方仓库](https://hub.docker.com/_/golang?tab=tags)自行选择后替换。
FROM golang:1.21.4-alpine3.18 as builder

# 设置环境变量
#ENV GO111MODULE=on \
#    GOPROXY=https://goproxy.cn,direct \
#    CGO_ENABLED=0 \
#    GOOS=linux \
#    GOARCH=amd64

# 设置工作目录
WORKDIR /app

# 复制文件到容器工作目录
COPY . .

# 编译
RUN go mod download && go build -o main cmd/main.go

# 选用运行时所用基础镜像（GO语言选择原则：尽量体积小、包含基础linux内容的基础镜像）
FROM alpine:3.13

# 容器默认时区为UTC，如需使用上海时间请启用以下时区设置命令
RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo Asia/Shanghai > /etc/timezone

# 使用 HTTPS 协议访问容器云调用证书安装
RUN apk add ca-certificates

# 指定运行时的工作目录
WORKDIR /app

# 从builder镜像中把可执行文件和配置文件拷贝到工作目录
COPY --from=builder /app/main .
COPY  ./configs ./configs

ENV ENV=LIVE

# 运行
EXPOSE 80
ENTRYPOINT ["./main"]