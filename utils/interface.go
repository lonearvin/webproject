package utils

type ContactPostData struct {
	Name    string `form:"name"`
	Email   string `form:"email"`
	Subject string `form:"subject"`
	Message string `form:"message"`
}

type SubscribeData struct {
	ID    uint   `gorm:"primaryKey"`
	Email string `gorm:"type:varchar(191);uniqueIndex;not null" form:"Email"`
}
