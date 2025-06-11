package utils

import (
	"github.com/gin-gonic/gin"
	"webproject/global"
)

// HandleContactPost 实现的功能是，接收数据，然后使用异步去发送邮件
func HandleContactPost(c *gin.Context) {
	var data SubscribeData
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	var existing SubscribeData
	if err := global.GlobalDB.Where("email = ?", data.Email).First(&existing).Error; err == nil {
		c.JSON(409, gin.H{"error": "This email is already subscribed"})
		return
	}

	if err := global.GlobalDB.Create(&data).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to save"})
		return
	}

	c.JSON(200, gin.H{"message": "Subscribed successfully"})
}
