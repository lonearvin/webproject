# 多阶段构建：阶段1 - 前端构建阶段
FROM docker.m.daocloud.io/library/node:20-alpine AS frontend-builder

WORKDIR /app

# 复制前端依赖文件
COPY package.json package-lock.json ./

# 设置npm代理
ENV npm_config_registry=https://registry.npmmirror.com

# 安装依赖（包括开发依赖，因为tailwindcss是开发依赖）
RUN npm ci

# 复制前端资源文件
COPY tailwind.config.js ./
COPY static/css/input.css static/css/custom.css ./static/css/

# 构建CSS
RUN npm run build:css

# 阶段2 - Go构建阶段
FROM docker.m.daocloud.io/library/golang:1.23-alpine AS go-builder

RUN apk add --no-cache git gcc musl-dev

WORKDIR /app

ENV GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
ENV GOSUMDB=off

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# 复制前端构建产物
COPY --from=frontend-builder /app/static/css/output.css ./static/css/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 阶段3 - 运行阶段
FROM docker.m.daocloud.io/library/alpine:3.20

RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai

WORKDIR /app

COPY --from=go-builder /app/main .
COPY --from=go-builder /app/static ./static
COPY --from=go-builder /app/templates ./templates
COPY --from=go-builder /app/config ./config

EXPOSE 8080

CMD ["./main"]
