package requests

type UserPutRequest struct {
	Name  string `json:"name" binding:"required"`
	Image string `json:"image" binding:"required"`
}

type UserEmailPutRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type UserPasswordPutRequest struct {
	Password       string `json:"password" binding:"required"`
	RepeatPassword string `json:"repeat_password" binding:"required"`
}
