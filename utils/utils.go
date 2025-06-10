package utils

import (
	"webproject/global"
)

// HandleContactPost 实现的功能是，接收数据，然后使用异步去发送邮件
func HandleContactPost(data ContactPostData) error {

	if err := global.GlobalDB.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
