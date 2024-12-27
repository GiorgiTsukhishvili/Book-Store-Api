package requests

import "mime/multipart"

type UserPutRequest struct {
	Name      string                `from:"name" binding:"required"`
	Image     *multipart.FileHeader `from:"image" binding:"required"`
	ImagePath string                `from:"image_path" binding:"required"`
}

type UserEmailPutRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type UserPasswordPutRequest struct {
	Password       string `json:"password" binding:"required"`
	RepeatPassword string `json:"repeat_password" binding:"required"`
}
