package utils

type ContactPostData struct {
	Name    string `form:"name"`
	Email   string `form:"email"`
	Subject string `form:"subject"`
	Message string `form:"message"`
}

type SubscribeData struct {
	Email string `form:"Email"`
}
