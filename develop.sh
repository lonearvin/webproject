
export VERSION=v.1.0.0

# 打包镜像
sudo docker build -t webproject:${VERSION} .

# 打标签
sudo docker tag webproject:${VERSION} 192.168.1.12:19001/webproject/webproject:${VERSION}

# 推送镜像
sudo docker push 192.168.1.12:19001/webproject/webproject:${VERSION}

# 拉取镜像，部署
sudo docker pull 192.168.1.12:19001/webproject/webproject:${VERSION}

sudo docker stop webproject
sudo docker rm webproject

sudo chmod +x develop.sh

# 部署容器
sudo docker run -d --name webproject \
  -p 8080:8080 \
  -v $(pwd)/config:/app/config \
  -v $(pwd)/templates:/app/templates \
  -v $(pwd)/static:/app/static \
  -v $(pwd)/log:/app/log \
  -e GIN_MODE=release \
  --restart unless-stopped \
  192.168.1.12:19001/webproject/webproject:${VERSION}
