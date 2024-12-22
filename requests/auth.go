package requests

import "github.com/GiorgiTsukhishvili/BookShelf-Api/models"

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserRegisterRequest struct {
	Name           string          `json:"name" binding:"required"`
	Email          string          `json:"email" binding:"required,email"`
	PhoneNumber    string          `json:"phone_number" binding:"required"`
	Password       string          `json:"password" binding:"required"`
	RepeatPassword string          `json:"repeat_password" binding:"required"`
	Type           models.UserType `json:"type" binding:"required"`
}

type UserVerifyRequest struct {
	Code string `json:"code" binding:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}
