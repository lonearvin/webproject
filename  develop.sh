
# 打包镜像
sudo docker build -t webproject:latest .

# 打标签
sudo docker tag webproject:latest 192.168.1.12:19001/webproject/webproject:latest

# 推送镜像
sudo docker push 192.168.1.12:19001/webproject/webproject:latest

# 拉取镜像，部署
sudo docker pull 192.168.1.12:19001/webproject/webproject:latest

# 部署容器
sudo docker run -d --name webproject \
  -p 8080:8080 \
  -v $(pwd)/config:/app/config \
  -v $(pwd)/templates:/app/templates \
  -v $(pwd)/static:/app/static \
  -v $(pwd)/log:/app/log \
  -e GIN_MODE=release \
  --restart unless-stopped \
  192.168.1.12:19001/webproject/webproject:latest