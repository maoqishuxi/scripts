# 第一阶段：构建 Go 应用程序
FROM golang:1.21-alpine AS builder

# 为构建设置工作目录
WORKDIR /app

# 将所有代码复制到工作目录中
COPY . .

# 编译 Go 应用程序
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# 第二阶段：创建精简版镜像
FROM alpine:latest

# 将时区数据包安装到镜像中（如果你的应用需要时区支持）
RUN apk --no-cache add ca-certificates tzdata

# 创建工作目录
WORKDIR /root/

# 从构建阶段复制生成的二进制文件
COPY --from=builder /app/main .

# 运行应用程序
CMD ["./main"]

# 暴露应用运行的端口
EXPOSE 9001
