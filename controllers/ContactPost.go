package controllers

import (
	"github.com/gin-gonic/gin"
	"webproject/global"
	"webproject/utils"
)

func ContactPost(c *gin.Context) {
	var data utils.ContactPostData
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	var existing utils.ContactPostData
	err := global.Db.Where("name = ? AND email = ? AND subject = ? AND message = ?",
		data.Name, data.Email, data.Subject, data.Message).First(&existing).Error

	if err == nil {
		c.JSON(409, gin.H{"error": "信息重复提交"})
		return
	}

	if err := global.Db.Create(&data).Error; err != nil {
		c.JSON(500, gin.H{"error": "保存失败"})
		return
	}

	c.JSON(200, gin.H{"message": "提交成功"})
}
