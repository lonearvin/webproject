# Docker部署说明

## 项目简介

这是凌图智控（上海）科技有限公司的网站项目，使用Go + Gin + GORM开发的后端服务。

## 部署步骤

### 1. 确保安装了Docker和Docker Compose

- [Docker安装指南](https://docs.docker.com/get-docker/)
- [Docker Compose安装指南](https://docs.docker.com/compose/install/)

### 2. 构建和运行容器

在项目根目录下执行以下命令：

```bash
# 构建镜像
docker-compose build

# 运行容器
docker-compose up -d
```

### 3. 访问网站

容器运行后，可以通过以下地址访问网站：

```
http://localhost:8080
```

### 4. 查看容器状态

```bash
# 查看容器状态
docker-compose ps

# 查看容器日志
docker-compose logs -f
```

### 5. 停止容器

```bash
docker-compose down
```

## 配置说明

### Dockerfile

- 使用最新的Go Alpine镜像作为基础
- 安装必要的依赖（git）
- 构建应用并运行

### docker-compose.yml

- 构建本地镜像
- 映射端口8080到宿主机
- 设置环境变量GIN_MODE=release
- 挂载本地目录到容器（便于开发调试）
- 配置容器自动重启

## 注意事项

1. 确保项目的配置文件（config/config.yaml）中的端口设置与Docker暴露的端口一致
2. 在生产环境中，建议移除volumes挂载，使用构建时的代码
3. 可以根据需要修改docker-compose.yml中的配置，如端口映射、环境变量等