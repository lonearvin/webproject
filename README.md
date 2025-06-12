以下是基于你提供的代码库信息生成的项目技术文档示例，你可以根据实际情况进行调整和补充。

# 凌图智控（上海）科技有限公司网站项目技术文档

## 一、项目概述
本项目是凌图智控（上海）科技有限公司的官方网站，旨在展示公司的业务、产品、服务以及团队信息等。项目采用了多种技术栈，包括 Go 语言、Gin 框架、GORM 数据库操作库、Vue 前端框架、Nginx 服务器、Redis 缓存和 MySQL 数据库。

## 二、技术栈
### 后端
- **Go 语言**：作为主要的编程语言，提供高效的并发处理能力和性能。
- **Gin 框架**：轻量级的 Web 框架，用于构建 RESTful API 和处理 HTTP 请求。
- **GORM**：强大的 Go 语言 ORM 库，简化数据库操作。
- **Google App Engine**：提供云服务支持。

### 前端
- **Vue**：渐进式 JavaScript 框架，用于构建用户界面。

### 数据库和缓存
- **MySQL**：关系型数据库，用于存储网站的各种数据，如用户信息、订阅信息等。
- **Redis**：内存数据库，可用于缓存数据，提高网站的响应速度。

### 服务器
- **Nginx**：高性能的 Web 服务器，用于反向代理和负载均衡。

## 三、项目结构
```
webproject
├── .gitignore
├── README.md
├── go.mod
├── go.sum
├── main.go
├── utils
│   ├── interface.go
│   └── utils.go
├── controllers
│   ├── ContactPost.go
│   ├── GetProductVideo.go
│   ├── GetService.go
│   ├── Subscribe.go
│   └── home_controllers.go
├── router
│   └── router.go
├── config
│   ├── config.go
│   └── config.yaml
├── templates
│   ├── 404.html
│   ├── ServicePages
│   │   └── ServicePages3C.html
│   ├── css
│   └── homepage.html
├── static
│   └── picture
└── global
    └── global.go
```

### 主要目录和文件说明
- **main.go**：项目的入口文件，负责初始化配置和启动服务器。
- **utils**：存放工具函数和接口定义。
- **controllers**：处理各种业务逻辑的控制器。
- **router**：路由配置文件，定义了不同 URL 的处理函数。
- **config**：配置文件目录，包含配置文件和读取配置的代码。
- **templates**：存放 HTML 模板文件。
- **static**：存放静态资源，如图片等。
- **global**：全局变量定义文件。

## 四、配置文件
项目使用 `config.yaml` 文件进行配置，配置内容包括 Gin 模式、端口号等。在 `main.go` 中通过 `config.InitConfig()` 函数读取配置文件。

```yaml
# config.yaml 示例
GinMode:
  Model: debug
App:
  Port: 8080
  TemplatePath: ./templates
```

## 五、数据库操作
### 全局数据库变量
在 `global/global.go` 中定义了全局数据库变量 `GlobalDB`，用于在项目中共享数据库连接。

```go
package global

import (
    "gorm.io/gorm"
)

var (
    GlobalDB *gorm.DB
)
```

### 数据库操作示例
以 `controllers/Subscribe.go` 为例，展示了如何进行订阅信息的数据库操作：
```go
package controllers

import (
    "github.com/gin-gonic/gin"
    "google.golang.org/appengine/log"
    "net/http"
    "webproject/global"
    "webproject/utils"
)

func Subscribe(ctx *gin.Context) {
    var data utils.SubscribeData
    err := ctx.ShouldBind(&data)

    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "code":    http.StatusBadRequest,
            "message": err.Error(),
        })
        log.Errorf(ctx, "subscribe faild: %v", err.Error())
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "code":    http.StatusInternalServerError,
            "message": err.Error(),
        })
        return
    }
    var existing utils.SubscribeData
    if err := global.GlobalDB.Where("email=?", data.Email).First(&existing).Error; err == nil {
        ctx.JSON(409, gin.H{
            "code":    409,
            "message": "信息重复提交",
        })
        return
    }

    // 进行数据库保存
    if err := global.GlobalDB.Create(&data).Error; err != nil {
        log.Errorf(ctx, "subscribe faild: %v", err.Error())
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "code":    http.StatusInternalServerError,
            "message": err.Error(),
        })
        log.Errorf(ctx, "subscribe faild: %v", err.Error())
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "code":    http.StatusOK,
        "message": "success",
    })
}
```

## 六、路由配置
在 `router/router.go` 中定义了项目的路由规则，包括静态资源路径、首页、视频页、服务页面、联系我们和订阅等接口。

```go
package router

import (
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
    "os"
    "webproject/config"
    "webproject/controllers"
)

func SetRouter() *gin.Engine {
    // 初始化 gin 引擎
    r := gin.Default()
    if gin.Mode() == gin.DebugMode {
        r.Use(func(c *gin.Context) {
            c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
            c.Writer.Header().Set("Pragma", "no-cache")
            c.Writer.Header().Set("Expires", "0")
            c.Next()
        })
    }
    // 注册静态资源路径
    staticPaths := map[string]string{
        "/static":  "./static",
        "/picture": "./picture",
        "/video":   "./video",
    }
    for route, path := range staticPaths {
        r.Static(route, path)
    }
    r.StaticFile("/favicon.ico", "./picture/favicon.ico")

    // 首页
    r.GET("/", controllers.Home)

    // 视频页
    r.GET("/api/video", controllers.GetProductVideo)

    // 服务页面（根据 id 显示对应 HTML）
    r.GET("/service", controllers.GetService)

    // 可选支持 REST 风格 URL: /service/3C
    r.GET("/service/:id", func(ctx *gin.Context) {
        ctx.Request.URL.RawQuery = "id=" + ctx.Param("id")
        controllers.GetService(ctx)
    })
    r.POST("/contact", controllers.ContactPost)
    r.POST("/subscribe", controllers.Subscribe)

    // 默认 404 页面处理
    r.NoRoute(func(ctx *gin.Context) {
        htmlFilePath := config.AppConfig.App.TemplatePath + "/404.html"
        if _, err := os.Stat(htmlFilePath); err == nil {
            ctx.File(htmlFilePath)
        } else {
            log.Printf("找不到404页面文件: %v", err)
            ctx.String(http.StatusNotFound, "404 Not Found")
        }
    })

    return r
}
```

## 七、前端页面
项目的前端页面使用 HTML 和 CSS 构建，存放在 `templates` 目录下。例如 `homepage.html` 展示了公司的简介、使命、愿景、团队介绍等信息。

```html
<!-- templates/homepage.html 部分代码 -->
<p class="text-muted mb-6">提供化自动化化工的生产线的解决方案。</p>
<!--                            TODO-->
<a href="#" class="text-primary font-medium flex items-center">
    了解更多 <i
        class="fa fa-arrow-right ml-2 transition-transform duration-300 group-hover:translate-x-1"></i>
</a>
```

## 八、部署和运行
### 1. 环境准备
确保已经安装了 Go 语言环境、MySQL 数据库和 Redis 缓存。

### 2. 配置数据库连接
在 `config.yaml` 中配置数据库连接信息。

### 3. 安装依赖
在项目根目录下执行以下命令安装项目依赖：
```sh
go mod tidy
```

### 4. 启动项目
在项目根目录下执行以下命令启动项目：
```sh
go run main.go
```

### 5. 访问项目
打开浏览器，访问 `http://localhost:8080` 即可查看网站。

## 九、后续开发计划
后续将进一步完善网站的功能，如添加更多的服务页面、优化用户体验、加强安全防护等。同时，将详细书写开发细节和架构文档，方便团队成员协作和项目的维护。
