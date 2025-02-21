package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webproject/utils"
)

// AuthMiddleWare 验证请求身份，验证JWT是否有效或者是是否有JWT
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 判断用户状态
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			c.Abort() // 当出现错误的时候停止请求任务，
			return
		}

		username, err := utils.ParseJWT(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			c.Abort()
			return
		}
		c.Set("username", username)
		c.Next()
	}
}
