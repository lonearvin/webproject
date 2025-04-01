
# 凌图智控（上海）科技有限公司 网站
工具：go+gin+gorm+vue+nginx+redis+mysql
# 主分支是main，当前在github上


# Nginx做反向代理的配置

```text
        location / {
            # 反向代理到后端服务器
            proxy_pass http://127.0.0.1:8080;  # 后端服务器地址
            proxy_set_header Host $host;  # 保留原始的主机名
            proxy_set_header X-Real-IP $remote_addr;  # 保留客户端真实 IP 地址
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;  # 保留客户端的所有转发 IP
            proxy_set_header X-Forwarded-Proto $scheme;  # 保留原始协议（http/https）
        }
        
        # 使用正则表达式匹配所有 /auth/api/ 下的请求接口
		location ~ ^/auth/api/(.*) {
			proxy_pass http://127.0.0.1:8080;  # 后端的实际处理接口地址
			proxy_set_header Host $host;
			proxy_set_header X-Real-IP $remote_addr;
			proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
			proxy_set_header X-Forwarded-Proto $scheme;
		}
```
其他更多的配置暂时不书写

### 后续将开发的细节和架构的开发文档书写一下


