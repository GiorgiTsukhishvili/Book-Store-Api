package requests

import "mime/multipart"

type UserPutRequest struct {
	Name      string                `form:"name" binding:"required"`
	Image     *multipart.FileHeader `form:"image"`
	ImagePath string                `form:"image_path"`
}

type UserEmailPutRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type UserPasswordPutRequest struct {
	Password       string `json:"password" binding:"required"`
	RepeatPassword string `json:"repeat_password" binding:"required"`
}
