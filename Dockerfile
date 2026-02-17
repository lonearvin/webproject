# 阶段1：构建阶段（包含编译环境）
FROM golang:1.22-alpine AS builder

# 合并安装所有必要依赖（避免重复构建层）
RUN apk add --no-cache git gcc musl-dev

# 设置工作目录
WORKDIR /app

# 配置国内Go代理，解决下载超时
ENV GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
ENV GOSUMDB=off

# 先复制依赖文件（利用Docker缓存，代码修改不重新下载依赖）
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制所有源代码
COPY . .

# 构建应用（添加CGO_ENABLED=0关闭CGO，生成静态编译二进制，无需系统依赖）
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 阶段2：运行阶段（仅保留运行所需文件，镜像体积从几百M降到几M）
FROM alpine:3.20

# 安装时区等基础工具（可选，根据你的应用需求）
RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai

# 设置工作目录
WORKDIR /app

# 从构建阶段复制编译好的二进制文件
COPY --from=builder /app/main .

# 暴露端口
EXPOSE 8080

# 运行应用
CMD ["./main"]