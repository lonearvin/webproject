package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webproject/global"
	"webproject/models"
	"webproject/utils"
)

func Register(ctx *gin.Context) {
	// 将数据引入其中
	var user models.User
	//    检查是否可以将用户的请求体数据与`User`模型绑定。如果不能绑定（例如数据格式错误），返回HTTP 400 错误，并返回相应的错误信息。
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 哈希密码进行加密
	hashedPW, err := utils.HashedPassword(user.Password)
	if err != nil {
		// 如果哈希失败，那么返回的是bad request状态吗
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 拿到哈希的密码 这里已经转换成了string
	user.Password = hashedPW
	// 通过用户名生成JWT
	token, err := utils.GenerateJWT(user.Username)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 文档解释
	// AutoMigrate 用于自动迁移您的 schema，保持您的 schema 是最新的。
	//    调用全局数据库的`AutoMigrate`方法，用于生成或更新用户模型对应的表结构。如果操作失败，返回HTTP 400错误。
	if err := global.GlobalDB.AutoMigrate(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	if err := global.GlobalDB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func Login(ctx *gin.Context) {

	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	// 从数据库中取出
	if err := global.GlobalDB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 取出结果是否相同
	if !utils.CheckPassword(input.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "wrong password"})
		return
	}
	token, err := utils.GenerateJWT(input.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
