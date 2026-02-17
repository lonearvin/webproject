# 使用最新的Go Alpine镜像
FROM golang:alpine

# 安装必要的依赖
RUN apk add --no-cache git

# 设置工作目录
WORKDIR /app

# 复制go.mod和go.sum文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制所有源代码
COPY . .

# 构建应用
RUN go build -o main .

# 暴露端口（与应用配置一致）
EXPOSE 8080

# 运行应用
CMD ["./main"]