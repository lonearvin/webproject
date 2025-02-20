# webproject
Go+Gin+Gorm+Vue3+Redis+MySQL

## gin框架的使用

头信息传递
```go
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, header)
	})
```
也可以使用gin自带的
```go
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		}
	})
```