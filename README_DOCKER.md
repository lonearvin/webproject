# Docker部署说明

## 项目简介

这是凌图智控（上海）科技有限公司的网站项目，使用Go + Gin + GORM开发的后端服务。

## 部署步骤

### 1. 确保安装了Docker、Docker Compose和harbor-registry

- [Docker安装指南](https://docs.docker.com/get-docker/)
- [Docker Compose安装指南](https://docs.docker.com/compose/install/)
- [harbor-registry安装指南](https://goharbor.io/docs/2.4.0/install-config/configure-harbor/)

### 2. 构建和运行容器

需要配置dockerfile中的镜像地址为harbor-registry的地址
需要配置容器运行时的环境变量docker-compose.yml

```bash
sudo chmod +X deploy.sh

./deploy.sh
```

## 注意事项

1. 确保项目的配置文件（config/config.yaml）中的端口设置与Docker暴露的端口一致
2. 在生产环境中，建议移除volumes挂载，使用构建时的代码
3. 可以根据需要修改docker-compose.yml中的配置，如端口映射、环境变量等